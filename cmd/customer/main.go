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
	v1CustomerPB "github.com/hinccvi/saga/internal/customer/controller/grpc/v1"
	k "github.com/hinccvi/saga/internal/customer/controller/kafka"
	customerRepo "github.com/hinccvi/saga/internal/customer/repository"
	customerService "github.com/hinccvi/saga/internal/customer/service"
	healthcheckPB "github.com/hinccvi/saga/internal/healthcheck/controller/grpc"
	"github.com/hinccvi/saga/pkg/db"
	"github.com/hinccvi/saga/pkg/log"
	"github.com/hinccvi/saga/proto/pb"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials"
	"google.golang.org/grpc/examples/data"
)

var (
	//nolint:gochecknoglobals // value of ldflags must be a package level variable
	Version = "1.0.0"
	//nolint:gochecknoglobals // environment flag that only used in main
	flagEnv = flag.String("env", "local", "environment")
	//nolint:gochecknoglobals // grpc port
	port = 50051
)

func main() {
	flag.Parse()

	// create root context
	ctx := context.Background()

	// create root logger tagged with server version
	logger := log.NewWithZap(log.New(*flagEnv, log.ErrorLog)).With(ctx, "version", Version)

	// load application configurations
	cfg, err := config.Load("customer", *flagEnv)
	if err != nil {
		logger.Fatalf("fail to load app config: %v", err)
	}

	// connect to database
	db, err := db.Connect(ctx, &cfg)
	if err != nil {
		logger.Fatalf("fail to connect to db: %v", err)
	}

	// timeout duration for each request
	t := time.Duration(cfg.Context.Timeout) * time.Second

	cs := customerService.New(cfg, customerRepo.New(db, logger), logger, t)

	// setup kafka consumer / writer
	handler := k.RegisterOrderHandlers(
		cfg.Kafka.Host,
		cfg.Kafka.OrderGroupID,
		cfg.Kafka.OrderWALTopic,
		cs,
		logger,
	)
	go func() {
		for {
			select {
			case err := <-handler.StartConsumer(ctx):
				logger.Fatalf("fail to consume kafka: %v", err)
			}
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

	pb.RegisterCustomerServiceServer(
		grpcServer,
		v1CustomerPB.RegisterHandlers(cs, logger),
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
