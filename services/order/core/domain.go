package orderCore

import "time"

type Order struct {
	OrderId         string    `json:"orderId"`
	CustomerId      string    `json:"customerId"`
	TotalAmount     int16     `json:"totalAmount"`
	Status          string    `json:"status"`
	ShippingAddress string    `json:"shippingAddress"`
	OrderAt         time.Time `json:"orderAt"`
	CreatedAt       time.Time `json:"createdAt"`
}

type OrderDetail struct {
	OrderDetailId string    `json:"orderDetailId"`
	OrderId       string    `json:"orderId"`
	ProductId     string    `json:"productId"`
	Quantity      int16     `json:"quantity"`
	Price         int16     `json:"price"`
	CreatedAt     time.Time `json:"createdAt"`
}

type OrderProfile struct {
	*Order
	Details []OrderDetail
}
