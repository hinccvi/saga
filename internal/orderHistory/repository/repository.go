package repository

import (
	"context"

	"github.com/google/uuid"
	"github.com/hinccvi/saga/internal/entity"
	"github.com/hinccvi/saga/pkg/log"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type (
	Repository interface {
		SaveCustomer(ctx context.Context, oh entity.OrderHistory) error
		SaveApprovedOrder(ctx context.Context, oh entity.OrderHistory) error
		GetOrderHistory(ctx context.Context, id uuid.UUID) (entity.OrderHistory, error)
	}
	repository struct {
		db         *mongo.Client
		collection *mongo.Collection
		logger     log.Logger
	}
)

const (
	databaseName   string = "saga"
	collectionName string = "order_history"
)

func New(db *mongo.Client, logger log.Logger) Repository {
	return repository{
		db,
		db.Database(databaseName).Collection(collectionName),
		logger,
	}
}

func (r repository) SaveCustomer(ctx context.Context, oh entity.OrderHistory) error {
	if _, err := r.collection.InsertOne(ctx, oh); err != nil {
		return err
	}

	return nil
}

func (r repository) SaveApprovedOrder(ctx context.Context, oh entity.OrderHistory) error {
	filter := bson.M{"customer_id": oh.CustomerID}
	update := bson.M{"$push": bson.M{"orders": oh.Orders[0]}}

	if _, err := r.collection.UpdateOne(ctx, filter, update); err != nil {
		return err
	}

	return nil
}

func (r repository) GetOrderHistory(ctx context.Context, id uuid.UUID) (entity.OrderHistory, error) {
	filter := bson.M{"customer_id": id}

	oh := &entity.OrderHistory{}
	if err := r.collection.FindOne(ctx, filter).Decode(oh); err != nil {
		return entity.OrderHistory{}, err
	}

	return *oh, nil
}
