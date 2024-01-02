package customerCore

import "time"

type Customer struct {
	Id        string    `json:"id"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	PhoneNo   string    `json:"phoneNo"`
	Status    bool      `json:"status"`
	Dob       time.Time `json:"dob"`
	CreatedAt time.Time `json:"createdAt"`
}

type Address struct {
	Id            string    `json:"id"`
	CustomerId    string    `json:"customerId"`
	StreetAddress string    `json:"streetAddress"`
	City          string    `json:"city"`
	State         string    `json:"state"`
	PostalCodes   string    `json:"postalCodes"`
	CreatedAt     time.Time `json:"createdAt"`
}

type CustomerProfile struct {
	*Customer
	Address []Address
}
