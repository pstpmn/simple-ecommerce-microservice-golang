package product

import (
	orderCore "simple-ecomerce-microservice/services/order/core"
	orderModel "simple-ecomerce-microservice/services/order/models"

	"go.mongodb.org/mongo-driver/mongo"
)

type (
	repo struct {
		db *mongo.Client
	}
)

// FindOrders implements orderCore.IOrderRepo.
func (*repo) FindOrders() (orderModel.Order, error) {
	panic("unimplemented")
}

// UpdateOneOrder implements orderCore.IOrderRepo.
func (*repo) UpdateOneOrder(model orderModel.Order) (orderModel.Order, error) {
	panic("unimplemented")
}

// UpdateOneOrderDetail implements orderCore.IOrderRepo.
func (*repo) UpdateOneOrderDetail(model orderModel.OrderDetail) (orderModel.OrderDetail, error) {
	panic("unimplemented")
}

func NewRepository(conn *mongo.Client) orderCore.IOrderRepo {
	return &repo{
		db: conn,
	}
}
