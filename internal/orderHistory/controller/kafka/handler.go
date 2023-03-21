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

	"github.com/hinccvi/saga/internal/orderHistory/service"
)

type (
	resource struct {
		host    string
		groupID string
		logger  log.Logger
		service service.Service
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
				State      string    `json:"state"`
			} `json:"after"`
		} `json:"payload"`
	}
	customerWALMessage struct {
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
				ID          uuid.UUID `json:"id"`
				Name        string    `json:"name"`
				CreditLimit string    `json:"credit_limit"`
			} `json:"after"`
		} `json:"payload"`
	}
)

const (
	minMessageBytes  int    = 10e3
	maxMessageBytes  int    = 10e6
	groupID          string = "order_history"
	customerWALTopic string = "saga.public.customer"
	orderWALTopic    string = "saga.public.order"
)

func RegisterOrderHistoryHandlers(host string, service service.Service, logger log.Logger) resource {
	return resource{
		host:    host,
		groupID: groupID,
		logger:  logger,
		service: service,
	}
}

func (r resource) StartCustomerWALConsumer(ctx context.Context) <-chan error {
	errCh := make(chan error)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{r.host},
		GroupID:  groupID,
		Topic:    customerWALTopic,
		MinBytes: minMessageBytes,
		MaxBytes: maxMessageBytes,
	})

	go func() {
		for {
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				r.logger.Errorf("[StartCustomerWALConsumer] internal error: %w", err)
				break
			}

			var cm customerWALMessage
			if err = json.Unmarshal(m.Value, &cm); err != nil {
				r.logger.Errorf("[StartCustomerWALConsumer] internal error: %w", err)
				continue
			}

			creditLimitByte, err := base64.StdEncoding.DecodeString(cm.Payload.After.CreditLimit)
			if err != nil {
				r.logger.Errorf("[StartCustomerWALConsumer] internal error: %w", err)
				continue
			}

			scale, err := strconv.ParseInt(cm.Schema.Fields[1].Fields[2].Parameters.Scale, 10, 32)
			if err != nil {
				r.logger.Errorf("[StartCustomerWALConsumer] internal error: %w", err)
				continue
			}

			creditLimit := decimal.NewFromBigInt(new(big.Int).SetBytes(creditLimitByte), -int32(scale))

			req := service.SaveCustomerRequest{
				CustomerID:  cm.Payload.After.ID,
				Name:        cm.Payload.After.Name,
				CreditLimit: creditLimit,
			}
			if err = r.service.SaveCustomer(ctx, req); err != nil {
				r.logger.Errorf("[StartCustomerWALConsumer] internal error: %w", err)
				continue
			}
		}

		if err := reader.Close(); err != nil {
			errCh <- fmt.Errorf("[StartCustomerWALConsumer] internal error: %w", err)
		}
	}()

	return errCh
}

func (r resource) StartOrderWALConsumer(ctx context.Context) <-chan error {
	errCh := make(chan error)

	reader := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{r.host},
		GroupID:  groupID,
		Topic:    orderWALTopic,
		MinBytes: minMessageBytes,
		MaxBytes: maxMessageBytes,
	})

	go func() {
		for {
			m, err := reader.ReadMessage(ctx)
			if err != nil {
				r.logger.Errorf("[StartOrderWALConsumer] internal error: %w", err)
				break
			}

			var om orderWALMessage
			if err = json.Unmarshal(m.Value, &om); err != nil {
				r.logger.Errorf("[StartOrderWALConsumer] internal error: %w", err)
				continue
			}

			if om.Payload.After.State != "APPROVED" {
				continue
			}

			orderTotalByte, err := base64.StdEncoding.DecodeString(om.Payload.After.OrderTotal)
			if err != nil {
				r.logger.Errorf("[StartOrderWALConsumer] internal error: %w", err)
				continue
			}

			scale, err := strconv.ParseInt(om.Schema.Fields[1].Fields[2].Parameters.Scale, 10, 32)
			if err != nil {
				r.logger.Errorf("[StartOrderWALConsumer] internal error: %w", err)
				continue
			}

			orderTotal := decimal.NewFromBigInt(new(big.Int).SetBytes(orderTotalByte), -int32(scale))

			req := service.SaveApprovedOrderRequest{
				CustomerID: om.Payload.After.CustomerID,
				ID:         om.Payload.After.OrderID,
				State:      om.Payload.After.State,
				OrderTotal: orderTotal,
			}
			if err = r.service.SaveApprovedOrder(ctx, req); err != nil {
				r.logger.Errorf("[StartOrderWALConsumer] internal error: %w", err)
				continue
			}
		}

		if err := reader.Close(); err != nil {
			errCh <- fmt.Errorf("[StartOrderWALConsumer] internal error: %w", err)
		}
	}()

	return errCh
}
