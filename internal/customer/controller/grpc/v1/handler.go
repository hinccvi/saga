package v1

import (
	"context"

	"github.com/hinccvi/saga/pkg/log"
	"github.com/shopspring/decimal"

	"github.com/hinccvi/saga/internal/customer/service"
	"github.com/hinccvi/saga/internal/errors"
	"github.com/hinccvi/saga/proto/pb"
)

type resource struct {
	pb.UnimplementedCustomerServiceServer
	logger  log.Logger
	service service.Service
}

func RegisterHandlers(service service.Service, logger log.Logger) resource {
	return resource{logger: logger, service: service}
}

func (r resource) CreateCustomer(ctx context.Context, req *pb.CreateCustomerRequest) (rep *pb.CreateCustomerReply, err error) {
	if req.Name == "" {
		err = errors.EmptyField.E()
		return
	}

	data := service.CreateCustomerRequest{
		Name:   req.Name,
		Amount: decimal.NewFromInt32(req.Amount),
	}

	id, err := r.service.CreateCustomer(ctx, data)
	if err != nil {
		return
	}

	rep = &pb.CreateCustomerReply{
		CustomerId: id.String(),
	}

	return
}
