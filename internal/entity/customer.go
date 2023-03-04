package entity

import (
	"time"

	"github.com/google/uuid"
	"github.com/shopspring/decimal"
)

type Customer struct {
	ID          uuid.UUID       `db:"id" json:"id"`
	Name        string          `db:"name" json:"name"`
	CreditLimit decimal.Decimal `db:"credit_limit" json:"credit_limit"`
	CreatedAt   time.Time       `db:"created_at" json:"created_at"`
}
