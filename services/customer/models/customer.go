package customerModel

import "time"

type CustomerModel struct {
	Id        string         `json:"id"`
	FirstName string         `json:"firstName" gorm:"not null"`
	LastName  string         `json:"lastName" gorm:"not null"`
	PhoneNo   string         `json:"phoneNo" gorm:"not null"`
	Status    bool           `json:"status" gorm:"not null"`
	Dob       time.Time      `json:"dob" gorm:"not null"`
	Addresses []AddressModel `json:"addresses" gorm:"foreignKey:CustomerId;references:Id"`
	UpdatedAt time.Time      `json:"updatedAt"`
	CreatedAt time.Time      `json:"createdAt" gorm:"not null"`
}

func (CustomerModel) TableName() string {
	return "customers"
}
