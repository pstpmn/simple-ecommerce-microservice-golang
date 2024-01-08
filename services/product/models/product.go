package productModel

import "time"

type ProductModel struct {
	ProductId   string        `json:"productId" gorm:"index;not null"`
	ProductName string        `json:"productName"`
	Description string        `json:"description"`
	Price       float64       `json:"price"`
	Stock       int64         `json:"stock"`
	CategoryId  string        `json:"categoryId"`
	Category    CategoryModel `json:"category" gorm:"foreignKey:CategoryId;references:Id"`
	UpdatedAt   time.Time     `json:"updatedAt"`
	CreatedAt   time.Time     `json:"createdAt"`
}

func (ProductModel) TableName() string {
	return "products"
}
