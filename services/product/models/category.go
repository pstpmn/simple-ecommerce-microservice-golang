package productModel

import "time"

type CategoryModel struct {
	Id        string    `gorm:"column:id;primaryKey" json:"id"`
	Name      string    `gorm:"column:name;not null" json:"name"`
	CreatedAt time.Time `gorm:"column:created_at;not null" json:"createdAt"`
}

func (CategoryModel) TableName() string {
	return "categories"
}
