package orderCore

import (
	orderModel "simple-ecomerce-microservice/services/order/models"
)

type IOrderUseCase interface {
	GetOrderDetail(customerId string)
	CreateOrder(order Order) (*OrderProfile, error)
	CreateOrderDetails(order []OrderDetail) (*OrderProfile, error)
	CancelOrder(customerId, orderId string) error
}

type IOrderRepo interface {
	FindOrders() (orderModel.Order, error)
	UpdateOneOrder(model orderModel.Order) (orderModel.Order, error)
	UpdateOneOrderDetail(model orderModel.OrderDetail) (orderModel.OrderDetail, error)
}

type IHelper interface {
	GenUuid() string
}
