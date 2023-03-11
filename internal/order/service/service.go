package service

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hinccvi/saga/internal/entity"
	"github.com/hinccvi/saga/internal/order/repository"
	"github.com/hinccvi/saga/pkg/log"
	"github.com/shopspring/decimal"
)

type (
	Service interface {
		CreateOrder(ctx context.Context, req CreateOrderRequest) (uuid.UUID, error)
		ApproveOrder(ctx context.Context, id uuid.UUID) error
		RejectOrder(ctx context.Context, id uuid.UUID) error
	}

	service struct {
		repo    repository.Repository
		logger  log.Logger
		timeout time.Duration
	}

	CreateOrderRequest struct {
		CustomerID uuid.UUID
		Amount     decimal.Decimal
	}
)

func New(repo repository.Repository, logger log.Logger, timeout time.Duration) Service {
	return service{repo, logger, timeout}
}

func (s service) CreateOrder(ctx context.Context, req CreateOrderRequest) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	c := entity.Order{
		CustomerID: req.CustomerID,
		OrderTotal: req.Amount,
		State:      entity.OrderPending,
	}
	id, err := s.repo.CreateOrder(ctx, c)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("[CreateOrder] internal error: %w", err)
	}

	return id, nil
}

func (s service) ApproveOrder(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	o := entity.Order{
		ID:    id,
		State: entity.OrderApproved,
	}
	if err := s.repo.UpdateOrderState(ctx, o); err != nil {
		return err
	}

	return nil
}

func (s service) RejectOrder(ctx context.Context, id uuid.UUID) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	o := entity.Order{
		ID:    id,
		State: entity.OrderRejected,
	}
	if err := s.repo.UpdateOrderState(ctx, o); err != nil {
		return err
	}

	return nil
}
