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
		host                     string
		groupID                  string
		creditReservedTopic      string
		creditLimitExceededTopic string
		minBytes                 int
		maxBytes                 int
		logger                   log.Logger
		service                  service.Service
	}
	orderMessage struct {
		OrderID uuid.UUID `json:"order_id"`
	}
)

const (
	minMessageBytes int = 10e3
	maxMessageBytes int = 10e6
)

func RegisterOrderHandlers(host, groupID, creditReservedTopic, creditLimitExceededTopic string, service service.Service, logger log.Logger) resource {
	return resource{
		host:                     host,
		groupID:                  groupID,
		creditReservedTopic:      creditReservedTopic,
		creditLimitExceededTopic: creditLimitExceededTopic,
		minBytes:                 minMessageBytes,
		maxBytes:                 maxMessageBytes,
		logger:                   logger,
		service:                  service,
	}
}

func (r resource) StartCreditReservedConsumer(ctx context.Context) <-chan error {
	errCh := make(chan error)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{r.host},
		GroupID:  r.groupID,
		Topic:    r.creditReservedTopic,
		MinBytes: r.minBytes,
		MaxBytes: r.maxBytes,
	})

	go func() {
		for {
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				r.logger.Errorf("[startCreditReservedConsumer] internal error: %w", err)
				break
			}

			var om orderMessage
			if err := json.Unmarshal(m.Value, &om); err != nil {
				r.logger.Errorf("[startCreditReservedConsumer] internal error: %w", err)
			}

			if err := r.service.ApproveOrder(ctx, om.OrderID); err != nil {
				r.logger.Errorf("[startCreditReservedConsumer] internal error: %w", err)
			}
		}

		if err := reader.Close(); err != nil {
			errCh <- fmt.Errorf("[startCreditReservedConsumer] internal error: %w", err)
		}
	}()

	return errCh
}

func (r resource) StartCreditLimitExceededConsumer(ctx context.Context) <-chan error {
	errCh := make(chan error)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{r.host},
		GroupID:  r.groupID,
		Topic:    r.creditLimitExceededTopic,
		MinBytes: r.minBytes,
		MaxBytes: r.maxBytes,
	})

	go func() {
		for {
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				r.logger.Errorf("[startCreditLimitExceededConsumer] internal error: %w", err)
				break
			}

			var om orderMessage
			if err := json.Unmarshal(m.Value, &om); err != nil {
				r.logger.Errorf("[startCreditLimitExceededConsumer] internal error: %w", err)
			}

			if err := r.service.RejectOrder(ctx, om.OrderID); err != nil {
				r.logger.Errorf("[startCreditLimitExceededConsumer] internal error: %w", err)
			}
		}

		if err := reader.Close(); err != nil {
			errCh <- fmt.Errorf("[startCreditLimitExceededConsumer] internal error: %w", err)
		}
	}()

	return errCh
}
