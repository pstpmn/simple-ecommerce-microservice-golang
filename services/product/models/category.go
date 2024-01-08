package productModel

import "time"

type CategoryModel struct {
	Id        string    `json:"id" gorm:"primaryKey"`
	Name      string    `json:"name" gorm:"not null"`
	CreatedAt time.Time `json:"createdAt" gorm:"not null"`
	// Products     []ProductModel
}

func (CategoryModel) TableName() string {
	return "categories"
}
