package kafka

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/google/uuid"
	"github.com/hinccvi/saga/pkg/log"
	"github.com/segmentio/kafka-go"

	"github.com/hinccvi/saga/internal/order/service"
)

type (
	resource struct {
		host    string
		groupID string
		logger  log.Logger
		service service.Service
	}
	orderMessage struct {
		OrderID uuid.UUID `json:"order_id"`
	}
)

//nolint:gosec // false positive.
const (
	minMessageBytes          int    = 10e3
	maxMessageBytes          int    = 10e6
	groupID                  string = "customer_group"
	creditReservedTopic      string = "saga.customer.credit_reserved"
	creditReservationFailed  string = "saga.customer.credit_reservation_failed"
	customerValidationFailed string = "saga.customer.validation_failed"
)

func RegisterOrderHandlers(host string, service service.Service, logger log.Logger) resource {
	return resource{
		host:    host,
		groupID: groupID,
		logger:  logger,
		service: service,
	}
}

func (r resource) StartCreditReservedConsumer(ctx context.Context) <-chan error {
	errCh := make(chan error)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{r.host},
		GroupID:  groupID,
		Topic:    creditReservedTopic,
		MinBytes: minMessageBytes,
		MaxBytes: maxMessageBytes,
	})

	go func() {
		for {
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				r.logger.Errorf("[startCreditReservedConsumer] internal error: %w", err)
				break
			}

			var om orderMessage
			if err = json.Unmarshal(m.Value, &om); err != nil {
				r.logger.Errorf("[startCreditReservedConsumer] internal error: %w", err)
				continue
			}

			if err = r.service.ApproveOrder(ctx, om.OrderID); err != nil {
				r.logger.Errorf("[startCreditReservedConsumer] internal error: %w", err)
				continue
			}
		}

		if err := reader.Close(); err != nil {
			errCh <- fmt.Errorf("[startCreditReservedConsumer] internal error: %w", err)
		}
	}()

	return errCh
}

func (r resource) StartCreditReservationFailedConsumer(ctx context.Context) <-chan error {
	errCh := make(chan error)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{r.host},
		GroupID:  groupID,
		Topic:    creditReservationFailed,
		MinBytes: minMessageBytes,
		MaxBytes: maxMessageBytes,
	})

	go func() {
		for {
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				r.logger.Errorf("[startCreditReservationFailedConsumer] internal error: %w", err)
				break
			}

			var om orderMessage
			if err = json.Unmarshal(m.Value, &om); err != nil {
				r.logger.Errorf("[startCreditReservationFailedConsumer] internal error: %w", err)
				continue
			}

			if err = r.service.RejectOrder(ctx, om.OrderID); err != nil {
				r.logger.Errorf("[startCreditReservationFailedConsumer] internal error: %w", err)
				continue
			}
		}

		if err := reader.Close(); err != nil {
			errCh <- fmt.Errorf("[startCreditReservationFailedConsumer] internal error: %w", err)
		}
	}()

	return errCh
}

func (r resource) StartCustomerValidationFailedConsumer(ctx context.Context) <-chan error {
	errCh := make(chan error)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{r.host},
		GroupID:  groupID,
		Topic:    customerValidationFailed,
		MinBytes: minMessageBytes,
		MaxBytes: maxMessageBytes,
	})

	go func() {
		for {
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				r.logger.Errorf("[startCreditReservationFailedConsumer] internal error: %w", err)
				break
			}

			var om orderMessage
			if err = json.Unmarshal(m.Value, &om); err != nil {
				r.logger.Errorf("[startCreditReservationFailedConsumer] internal error: %w", err)
				continue
			}

			if err = r.service.RejectOrder(ctx, om.OrderID); err != nil {
				r.logger.Errorf("[startCreditReservationFailedConsumer] internal error: %w", err)
				continue
			}
		}

		if err := reader.Close(); err != nil {
			errCh <- fmt.Errorf("[startCreditReservationFailedConsumer] internal error: %w", err)
		}
	}()

	return errCh
}
