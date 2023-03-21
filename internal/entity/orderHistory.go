package entity

import (
	"github.com/google/uuid"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderHistory struct {
	ID         primitive.ObjectID `bson:"_id" json:"id"`
	CustomerID uuid.UUID          `bson:"customer_id" json:"customer_id"`
	Orders     []struct {
		ID         uuid.UUID `bson:"id" json:"id"`
		State      string    `bson:"state" json:"state"`
		OrderTotal struct {
			Amount primitive.Decimal128 `bson:"amount" json:"amount"`
		} `bson:"order_total" json:"order_total"`
	}
	Name        string `bson:"name" json:"name"`
	CreditLimit struct {
		Amount primitive.Decimal128 `bson:"amount" json:"amount"`
	} `bson:"credit_limit" json:"credit_limit"`
}
