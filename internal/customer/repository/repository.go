package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/hinccvi/saga/internal/entity"
	"github.com/hinccvi/saga/pkg/log"
	"github.com/jmoiron/sqlx"
	"github.com/shopspring/decimal"
)

type (
	Repository interface {
		CreateCustomer(ctx context.Context, c entity.Customer) (uuid.UUID, error)
		GetCreditLimit(ctx context.Context, id uuid.UUID) (decimal.Decimal, error)
	}
	repository struct {
		db     *sqlx.DB
		logger log.Logger
	}
)

const (
	createCustomerQuery string = `INSERT INTO customer (name, credit_limit) VALUES (:name, :credit_limit) RETURNING id`
	getCreditLimitQuery string = `SELECT credit_limit FROM customer WHERE id = $1`
)

func New(db *sqlx.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) CreateCustomer(ctx context.Context, c entity.Customer) (uuid.UUID, error) {
	createCustomerStmt, err := r.db.PrepareNamedContext(ctx, createCustomerQuery)
	if err != nil {
		return uuid.UUID{}, err
	}
	defer createCustomerStmt.Close()

	if err = createCustomerStmt.GetContext(ctx, &c, c); err != nil {
		return uuid.UUID{}, err
	}

	return c.ID, nil
}

func (r repository) GetCreditLimit(ctx context.Context, id uuid.UUID) (decimal.Decimal, error) {
	getCreditLimitStmt, err := r.db.PreparexContext(ctx, getCreditLimitQuery)
	if err != nil {
		return decimal.Decimal{}, err
	}
	defer getCreditLimitStmt.Close()

	var c entity.Customer
	if err = getCreditLimitStmt.GetContext(ctx, &c, id); err != nil {
		return decimal.Decimal{}, err
	}

	return c.CreditLimit, nil
}
