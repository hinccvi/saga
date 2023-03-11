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
		UpdateOrderState(ctx context.Context, o entity.Order) error
	}
	repository struct {
		db     *sqlx.DB
		logger log.Logger
	}
)

const (
	createOrderQuery string = `INSERT INTO "order" (customer_id, order_total, state) 
                        VALUES (:customer_id, :order_total, :state) RETURNING id`
	updateOrderStateQuery string = `UPDATE "order" SET state = :state WHERE id = :id`
)

func New(db *sqlx.DB, logger log.Logger) Repository {
	return repository{db, logger}
}

func (r repository) CreateOrder(ctx context.Context, o entity.Order) (uuid.UUID, error) {
	createOrderStmt, err := r.db.PrepareNamedContext(ctx, createOrderQuery)
	if err != nil {
		return uuid.UUID{}, err
	}
	defer createOrderStmt.Close()

	if err = createOrderStmt.GetContext(ctx, &o, o); err != nil {
		return uuid.UUID{}, err
	}

	return o.ID, nil
}

func (r repository) UpdateOrderState(ctx context.Context, o entity.Order) error {
	updateOrderStateStmt, err := r.db.PrepareNamedContext(ctx, updateOrderStateQuery)
	if err != nil {
		return err
	}
	defer updateOrderStateStmt.Close()

	if _, err := updateOrderStateStmt.ExecContext(ctx, o); err != nil {
		return err
	}

	return nil
}
