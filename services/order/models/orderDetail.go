package orderModel

import "time"

type OrderDetail struct {
	OrderDetailId string    `json:"orderDetailId" gorm:"primaryKey"`
	OrderId       string    `json:"orderId" gorm:"index;not null"`
	ProductId     string    `json:"productId" gorm:"index;not null"`
	Quantity      int16     `json:"quantity" gorm:"not null"`
	Price         int16     `json:"price" gorm:"not null"`
	CreatedAt     time.Time `json:"createdAt" gorm:"not null"`
}

func (OrderDetail) TableName() string {
	return "orderDetails"
}
