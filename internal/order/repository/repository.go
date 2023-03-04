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
		CreateOrder(ctx context.Context, o entity.Order) (uuid.UUID, error)
	}
	repository struct {
		db     *sqlx.DB
		logger log.Logger
	}
)

const (
	createOrder string = `INSERT INTO "order" (customer_id, order_total, state) 
                        VALUES (:customer_id, :order_total, :state) RETURNING id`
)

func New(db *sqlx.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) CreateOrder(ctx context.Context, o entity.Order) (uuid.UUID, error) {
	createOrderStmt, err := r.db.PrepareNamedContext(ctx, createOrder)
	if err != nil {
		return uuid.UUID{}, err
	}
	defer createOrderStmt.Close()

	if err = createOrderStmt.GetContext(ctx, &o, o); err != nil {
		return uuid.UUID{}, err
	}

	return o.ID, nil
}
