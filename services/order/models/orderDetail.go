package orderModel

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type OrderDetail struct {
	OrderDetailId primitive.ObjectID `json:"orderDetailId" bson:"_id,omitempty"`
	OrderId       primitive.ObjectID `json:"orderId" bson:"orderId"`
	ProductId     string             `json:"productId" bson:"productId"`
	Quantity      int16              `json:"quantity" bson:"quantity"`
	Price         int16              `json:"price" bson:"price"`
	CreatedAt     time.Time          `json:"createdAt" bson:"createdAt"`
}

func (OrderDetail) TableName() string {
	return "order_details"
}
