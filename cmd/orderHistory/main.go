//nolint:cyclop,funlen //driver program
package main

import (
	"context"
	"flag"
	"fmt"
	"net"
	"os"
	"os/signal"
	"syscall"
	"time"

	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	grpc_recovery "github.com/grpc-ecosystem/go-grpc-middleware/recovery"
	"github.com/hinccvi/saga/internal/config"
	healthcheckPB "github.com/hinccvi/saga/internal/healthcheck/controller/grpc"
	v1OrderHistoryPB "github.com/hinccvi/saga/internal/orderHistory/controller/grpc/v1"
	k "github.com/hinccvi/saga/internal/orderHistory/controller/kafka"
	orderHistoryRepo "github.com/hinccvi/saga/internal/orderHistory/repository"
	orderHistoryService "github.com/hinccvi/saga/internal/orderHistory/service"
	"github.com/hinccvi/saga/pkg/db"
	"github.com/hinccvi/saga/pkg/log"
	"github.com/hinccvi/saga/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
)

//nolint:gochecknoglobals //config
var (
	Version = "1.0.0"
	flagEnv = flag.String("env", "local", "environment")
	port    = 50053
)

func main() {
	flag.Parse()

	// create root context
	ctx := context.Background()

	// create root logger tagged with server version
	logger := log.NewWithZap(log.New(*flagEnv, log.ErrorLog)).With(ctx, "version", Version)

	// load application configurations
	cfg, err := config.Load("order", *flagEnv)
	if err != nil {
		logger.Fatalf("fail to load app config: %v", err)
	}

	// connect to database
	db, err := db.ConnectMongo(ctx, cfg.Mongo.Dsn)
	if err != nil {
		logger.Fatalf("fail to connect to db: %v", err)
	}

	// timeout duration for each request
	t := time.Duration(cfg.Context.Timeout) * time.Second

	ohs := orderHistoryService.New(orderHistoryRepo.New(db, logger), logger, t)

	// setup kafka consumer / writer
	handler := k.RegisterOrderHistoryHandlers(
		cfg.Kafka.Host,
		ohs,
		logger,
	)
	go func() {
		for err := range handler.StartCustomerWALConsumer(ctx) {
			logger.Fatalf("fail to consume customer wal topic: %v", err)
		}
	}()
	go func() {
		for err := range handler.StartOrderWALConsumer(ctx) {
			logger.Fatalf("fail to consume order wal topic: %v", err)
		}
	}()

	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", port))
	if err != nil {
		logger.Fatalf("failed to listen: %v", err)
	}

	var opts []grpc.ServerOption
	if cfg.App.Cert == "" {
		cfg.App.Cert = data.Path("x509/server_cert.pem")
	}
	if cfg.App.Key == "" {
		cfg.App.Key = data.Path("x509/server_key.pem")
	}
	creds, err := credentials.NewServerTLSFromFile(cfg.App.Cert, cfg.App.Key)
	if err != nil {
		logger.Fatalf("failed to generate credentials: %v", err)
	}

	opts = []grpc.ServerOption{
		grpc.Creds(creds),
		grpc.UnaryInterceptor(
			grpc_middleware.ChainUnaryServer(
				grpc_recovery.UnaryServerInterceptor(),
			)),
	}

	// register grpc server
	grpcServer := grpc.NewServer(opts...)

	pb.RegisterHealthcheckServiceServer(
		grpcServer,
		healthcheckPB.RegisterHandlers(Version),
	)

	pb.RegisterOrderHistoryServiceServer(
		grpcServer,
		v1OrderHistoryPB.RegisterHandlers(ohs, logger),
	)

	// seperate goroutine to listen on kill signal
	go func() {
		quit := make(chan os.Signal, 1)
		signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
		<-quit

		logger.Info("Server shutting down")

		grpcServer.GracefulStop()

		logger.Info("Server exiting")
	}()

	logger.Infof("grpc server listening on %v", lis.Addr())

	if err = grpcServer.Serve(lis); err != nil {
		logger.Fatalf("failed to serve grpc server: %v", err)
	}
}
