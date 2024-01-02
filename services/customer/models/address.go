package customerModel

import "time"

type AddressModel struct {
	Id            string    `json:"id"`
	CustomerId    string    `json:"customerId" gorm:"index;not null"`
	StreetAddress string    `json:"streetAddress"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	PostalCodes   string    `json:"postalCodes"`
	UpdatedAt     time.Time `json:"updatedAt"`
	CreatedAt     time.Time `json:"createdAt"`
}

func (AddressModel) TableName() string {
	// Replace "your_schema" with the desired schema name
	return "addresses"
}
