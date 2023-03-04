package service

import (
	"context"
	"fmt"
	"time"

	"github.com/go-redis/redis/v9"
	"github.com/google/uuid"
	"github.com/hinccvi/saga/internal/customer/repository"
	"github.com/hinccvi/saga/internal/entity"
	"github.com/hinccvi/saga/pkg/log"
	"github.com/shopspring/decimal"
)

type (
	Service interface {
		CreateCustomer(ctx context.Context, req CreateCustomerRequest) (uuid.UUID, error)
	}

	service struct {
		rds     redis.Client
		repo    repository.Repository
		logger  log.Logger
		timeout time.Duration
	}

	CreateCustomerRequest struct {
		Name   string
		Amount decimal.Decimal
	}
)

func New(rds redis.Client, repo repository.Repository, logger log.Logger, timeout time.Duration) Service {
	return service{rds, repo, logger, timeout}
}

func (s service) CreateCustomer(ctx context.Context, req CreateCustomerRequest) (uuid.UUID, error) {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	c := entity.Customer{
		Name:        req.Name,
		CreditLimit: req.Amount,
	}
	id, err := s.repo.CreateCustomer(ctx, c)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("[CreateCustomer] internal error: %w", err)
	}

	return id, nil
}
