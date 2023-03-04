package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/hinccvi/saga/internal/entity"
	"github.com/hinccvi/saga/internal/order/repository"
	"github.com/hinccvi/saga/pkg/log"
	"github.com/shopspring/decimal"
)

type (
	Service interface {
		CreateOrder(ctx context.Context, req CreateOrderRequest) (uuid.UUID, error)
	}

	service struct {
		rds     redis.Client
		repo    repository.Repository
		logger  log.Logger
		timeout time.Duration
	}

	CreateOrderRequest struct {
		CustomerID uuid.UUID
		Amount     decimal.Decimal
	}
)

func New(rds redis.Client, repo repository.Repository, logger log.Logger, timeout time.Duration) Service {
	return service{rds, repo, logger, timeout}
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
