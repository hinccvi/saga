package kafka

import (
	"context"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"math/big"
	"strconv"

	"github.com/google/uuid"
	"github.com/hinccvi/saga/pkg/log"
	"github.com/segmentio/kafka-go"
	"github.com/shopspring/decimal"

	"github.com/hinccvi/saga/internal/customer/service"
)

type (
	resource struct {
		host     string
		groupID  string
		topic    string
		minBytes int
		maxBytes int
		logger   log.Logger
		service  service.Service
	}
	orderWALMessage struct {
		Schema struct {
			Fields []struct {
				Fields []struct {
					Name       string `json:"name"`
					Parameters struct {
						Scale string `json:"scale"`
					} `json:"parameters"`
				} `json:"fields"`
			} `json:"fields"`
		} `json:"schema"`
		Payload struct {
			After struct {
				OrderID    uuid.UUID `json:"id"`
				CustomerID uuid.UUID `json:"customer_id"`
				OrderTotal string    `json:"order_total"`
			} `json:"after"`
		} `json:"payload"`
	}
)

const (
	minMessageBytes int = 10e3
	maxMessageBytes int = 10e6
)

func RegisterCustomerHandlers(host, groupID, topic string, service service.Service, logger log.Logger) resource {
	return resource{
		host:     host,
		groupID:  groupID,
		topic:    topic,
		minBytes: minMessageBytes,
		maxBytes: maxMessageBytes,
		logger:   logger,
		service:  service,
	}
}

func (r resource) StartConsumer(ctx context.Context) <-chan error {
	errCh := make(chan error)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{r.host},
		GroupID:  r.groupID,
		Topic:    r.topic,
		MinBytes: r.minBytes,
		MaxBytes: r.maxBytes,
	})

	go func() {
		for {
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				r.logger.Errorf("[startCustomerConsumer] internal error: %w", err)
				break
			}

			var om orderWALMessage
			if err = json.Unmarshal(m.Value, &om); err != nil {
				r.logger.Errorf("[startCustomerConsumer] internal error: %w", err)
			}

			orderTotalByte, err := base64.StdEncoding.DecodeString(om.Payload.After.OrderTotal)
			if err != nil {
				r.logger.Errorf("[startCustomerConsumer] internal error: %w", err)
			}

			scale, err := strconv.Atoi(om.Schema.Fields[1].Fields[2].Parameters.Scale)
			if err != nil {
				r.logger.Errorf("[startCustomerConsumer] internal error: %w", err)
			}

			orderTotal := decimal.NewFromBigInt(new(big.Int).SetBytes(orderTotalByte), -int32(scale))

			req := service.ReserveCreditRequest{
				OrderID:    om.Payload.After.OrderID,
				CustomerID: om.Payload.After.CustomerID,
				Amount:     orderTotal,
			}
			if err = r.service.ReserveCredit(ctx, req); err != nil {
				r.logger.Errorf("[startCustomerConsumer] internal error: %w", err)
			}
		}

		if err := reader.Close(); err != nil {
			errCh <- fmt.Errorf("[startCustomerConsumer] internal error: %w", err)
		}
	}()

	return errCh
}
