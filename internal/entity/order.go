package entity

import (
	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type OrderState string

type Order struct {
	ID         uuid.UUID       `db:"id" json:"id"`
	CustomerID uuid.UUID       `db:"customer_id" json:"customer_id"`
	OrderTotal decimal.Decimal `db:"order_total" json:"order_total"`
	State      OrderState      `db:"state" json:"state"`
}

const (
	OrderRejected OrderState = "REJECTED"
	OrderApproved OrderState = "APPROVED"
	OrderPending  OrderState = "PENDING"
)
