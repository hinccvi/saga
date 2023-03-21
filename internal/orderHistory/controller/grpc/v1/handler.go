package v1

import (
	"context"

	"github.com/google/uuid"
	"github.com/hinccvi/saga/pkg/log"

	"github.com/hinccvi/saga/internal/orderHistory/service"
	"github.com/hinccvi/saga/proto/pb"
)

type resource struct {
	pb.OrderHistoryServiceServer
	logger  log.Logger
	service service.Service
}

func RegisterHandlers(service service.Service, logger log.Logger) resource {
	return resource{logger: logger, service: service}
}

func (r resource) GetOrderHistory(ctx context.Context, req *pb.GetOrderHistoryRequest) (*pb.GetOrderHistoryReply, error) {
	id, err := uuid.Parse(req.CustomerId)
	if err != nil {
		return &pb.GetOrderHistoryReply{}, err
	}

	ol, err := r.service.GetOrderHistory(ctx, id)
	if err != nil {
		return &pb.GetOrderHistoryReply{}, err
	}

	rep := &pb.GetOrderHistoryReply{
		CustomerId: ol.CustomerID.String(),
		Name:       ol.Name,
		CreditLimit: &pb.CreditLimit{
			Amount: ol.CreditLimit.Amount.String(),
		},
	}

	for _, order := range ol.Orders {
		rep.Orders = append(rep.Orders, &pb.Order{
			Id:    order.ID.String(),
			State: order.State,
			OrderHistoryDetail: &pb.OrderHistoryDetail{
				Amount: order.OrderTotal.Amount.String(),
			},
		})
	}

	return rep, nil
}
