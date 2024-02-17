package orderModel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Order struct {
	OrderId         primitive.ObjectID `json:"orderId" bson:"_id,omitempty"`
	CustomerId      string             `json:"customerId" bson:"customerId"`
	Status          string             `json:"status" bson:"status"`
	ShippingAddress string             `json:"shippingAddress" bson:"shippingAddress"`
	OrderAt         time.Time          `json:"orderAt" bson:"orderAt"`
	CreatedAt       time.Time          `json:"createdAt" bson:"createdAt"`
}

func (Order) TableName() string {
	return "orders"
}
