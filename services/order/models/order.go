package orderModel

import "time"

type Order struct {
	OrderId         string    `json:"orderId" gorm:"primaryKey;not null"`
	CustomerId      string    `json:"customerId" gorm:"index;not null"`
	TotalAmount     int16     `json:"totalAmount" gorm:"not null"`
	Status          string    `json:"status" gorm:"not null"`
	ShippingAddress string    `json:"shippingAddress" gorm:"not null"`
	OrderAt         time.Time `json:"orderAt" gorm:"not null"`
	CreatedAt       time.Time `json:"createdAt" gorm:"not null"`
}

func (Order) TableName() string {
	return "orders"
}
