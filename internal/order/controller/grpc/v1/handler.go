package v1

import (
	"context"

	"github.com/google/uuid"
	"github.com/hinccvi/saga/pkg/log"
	"github.com/shopspring/decimal"

	"github.com/hinccvi/saga/internal/order/service"
	"github.com/hinccvi/saga/proto/pb"
)

type resource struct {
	pb.UnimplementedOrderServiceServer
	logger  log.Logger
	service service.Service
}

func RegisterHandlers(service service.Service, logger log.Logger) resource {
	return resource{logger: logger, service: service}
}

func (r resource) CreateOrder(ctx context.Context, req *pb.CreateOrderRequest) (rep *pb.CreateOrderReply, err error) {
	id, err := uuid.Parse(req.CustomerId)
	if err != nil {
		return &pb.CreateOrderReply{}, err
	}

	data := service.CreateOrderRequest{
		CustomerID: id,
		Amount:     decimal.NewFromInt32(req.OrderDetail.Amount),
	}
	id, err = r.service.CreateOrder(ctx, data)
	if err != nil {
		return
	}

	rep = &pb.CreateOrderReply{
		Id: id.String(),
	}

	return
}
