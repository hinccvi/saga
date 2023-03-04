package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/hinccvi/saga/internal/entity"
	"github.com/hinccvi/saga/pkg/log"
	"github.com/jmoiron/sqlx"
)

type (
	Repository interface {
		CreateCustomer(ctx context.Context, c entity.Customer) (uuid.UUID, error)
	}
	repository struct {
		db     *sqlx.DB
		logger log.Logger
	}
)

const (
	createCustomer string = `INSERT INTO customer (name, credit_limit) VALUES (:name, :credit_limit) RETURNING id`
)

func New(db *sqlx.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) CreateCustomer(ctx context.Context, c entity.Customer) (uuid.UUID, error) {
	createCustomerStmt, err := r.db.PrepareNamedContext(ctx, createCustomer)
	if err != nil {
		return uuid.UUID{}, err
	}
	defer createCustomerStmt.Close()

	if err = createCustomerStmt.GetContext(ctx, &c, c); err != nil {
		return uuid.UUID{}, err
	}

	return c.ID, nil
}
