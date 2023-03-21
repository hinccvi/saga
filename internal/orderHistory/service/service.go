package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hinccvi/saga/internal/entity"
	"github.com/hinccvi/saga/internal/orderHistory/repository"
	"github.com/hinccvi/saga/pkg/log"
	"github.com/shopspring/decimal"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	Service interface {
		SaveCustomer(ctx context.Context, req SaveCustomerRequest) error
		SaveApprovedOrder(ctx context.Context, req SaveApprovedOrderRequest) error
		GetOrderHistory(ctx context.Context, id uuid.UUID) (entity.OrderHistory, error)
	}
	service struct {
		repo    repository.Repository
		logger  log.Logger
		timeout time.Duration
	}
	SaveCustomerRequest struct {
		CustomerID  uuid.UUID
		Name        string
		CreditLimit decimal.Decimal
	}
	SaveApprovedOrderRequest struct {
		CustomerID uuid.UUID
		ID         uuid.UUID
		State      string
		OrderTotal decimal.Decimal
	}
	CreateOrderRequest struct {
		CustomerID uuid.UUID
		Amount     decimal.Decimal
	}
)

func New(repo repository.Repository, logger log.Logger, timeout time.Duration) Service {
	return service{repo, logger, timeout}
}

func (s service) SaveCustomer(ctx context.Context, req SaveCustomerRequest) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	amount, err := primitive.ParseDecimal128(req.CreditLimit.String())
	if err != nil {
		return fmt.Errorf("[SaveCustomer] internal error: %w", err)
	}

	oh := entity.OrderHistory{
		ID:         primitive.NewObjectID(),
		CustomerID: req.CustomerID,
		Name:       req.Name,
		Orders: make([]struct {
			ID         uuid.UUID "bson:\"id\" json:\"id\""
			State      string    "bson:\"state\" json:\"state\""
			OrderTotal struct {
				Amount primitive.Decimal128 "bson:\"amount\" json:\"amount\""
			} "bson:\"order_total\" json:\"order_total\""
		}, 0),
		CreditLimit: struct {
			Amount primitive.Decimal128 "bson:\"amount\" json:\"amount\""
		}{
			Amount: amount,
		},
	}
	if err = s.repo.SaveCustomer(ctx, oh); err != nil {
		return fmt.Errorf("[SaveCustomer] internal error: %w", err)
	}

	return nil
}

func (s service) SaveApprovedOrder(ctx context.Context, req SaveApprovedOrderRequest) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	amount, err := primitive.ParseDecimal128(req.OrderTotal.String())
	if err != nil {
		return fmt.Errorf("[SaveApprovedOrder] internal error: %w", err)
	}

	oh := entity.OrderHistory{
		CustomerID: req.CustomerID,
		Orders: []struct {
			ID         uuid.UUID "bson:\"id\" json:\"id\""
			State      string    "bson:\"state\" json:\"state\""
			OrderTotal struct {
				Amount primitive.Decimal128 "bson:\"amount\" json:\"amount\""
			} "bson:\"order_total\" json:\"order_total\""
		}{
			struct {
				ID         uuid.UUID "bson:\"id\" json:\"id\""
				State      string    "bson:\"state\" json:\"state\""
				OrderTotal struct {
					Amount primitive.Decimal128 "bson:\"amount\" json:\"amount\""
				} "bson:\"order_total\" json:\"order_total\""
			}{
				ID:    req.ID,
				State: req.State,
				OrderTotal: struct {
					Amount primitive.Decimal128 "bson:\"amount\" json:\"amount\""
				}{
					Amount: amount,
				},
			},
		},
	}
	if err = s.repo.SaveApprovedOrder(ctx, oh); err != nil {
		return fmt.Errorf("[SaveApproveOrder] internal error: %w", err)
	}

	return nil
}

func (s service) GetOrderHistory(ctx context.Context, id uuid.UUID) (entity.OrderHistory, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	oh, err := s.repo.GetOrderHistory(ctx, id)
	if err != nil {
		return entity.OrderHistory{}, fmt.Errorf("[GetOrderHistory] internal error: %w", err)
	}

	return oh, nil
}
