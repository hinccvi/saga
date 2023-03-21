package service

import (
	"context"
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/hinccvi/saga/internal/config"
	"github.com/hinccvi/saga/internal/customer/repository"
	"github.com/hinccvi/saga/internal/entity"
	"github.com/hinccvi/saga/pkg/log"
	"github.com/segmentio/kafka-go"
	"github.com/shopspring/decimal"
)

type (
	Service interface {
		CreateCustomer(ctx context.Context, req CreateCustomerRequest) (uuid.UUID, error)
		ReserveCredit(ctx context.Context, req ReserveCreditRequest) error
	}

	service struct {
		cfg     config.Config
		repo    repository.Repository
		logger  log.Logger
		timeout time.Duration
	}

	CreateCustomerRequest struct {
		Name   string
		Amount decimal.Decimal
	}

	ReserveCreditRequest struct {
		OrderID    uuid.UUID
		CustomerID uuid.UUID
		Amount     decimal.Decimal
	}
)

//nolint:gosec //false positive
const (
	creditReservedTopic      string = "saga.customer.credit_reserved"
	creditReservationFailed  string = "saga.customer.credit_reservation_failed"
	customerValidationFailed string = "saga.customer.validation_failed"
)

func New(cfg config.Config, repo repository.Repository, logger log.Logger, timeout time.Duration) Service {
	return service{cfg, repo, logger, timeout}
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

func (s service) ReserveCredit(ctx context.Context, req ReserveCreditRequest) error {
	ctx, cancel := context.WithTimeout(ctx, s.timeout)
	defer cancel()

	w := &kafka.Writer{
		Addr:     kafka.TCP(s.cfg.Kafka.Host),
		Balancer: &kafka.LeastBytes{},
		Topic:    creditReservationFailed,
	}

	creditLimit, err := s.repo.GetCreditLimit(ctx, req.CustomerID)
	if err != nil && errors.Is(err, sql.ErrNoRows) {
		w.Topic = customerValidationFailed
	}

	if err == nil && req.Amount.LessThan(creditLimit) {
		w.Topic = creditReservedTopic
	}

	bytes, err := json.Marshal(struct {
		OrderID uuid.UUID `json:"order_id"`
	}{
		req.OrderID,
	})
	if err != nil {
		return fmt.Errorf("[ReserveCredit] internal error: %w", err)
	}

	err = w.WriteMessages(ctx,
		kafka.Message{
			Value: bytes,
		},
	)
	if err != nil {
		return fmt.Errorf("[ReserveCredit] internal error: %w", err)
	}

	if err = w.Close(); err != nil {
		return fmt.Errorf("[ReserveCredit] internal error: %w", err)
	}

	return nil
}
