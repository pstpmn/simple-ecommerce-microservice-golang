package customerHanlder

import (
	"time"

	validation "github.com/go-ozzo/ozzo-validation"
)

type (
	CreateCustomerDto struct {
		FirstName string    `json:"firstName"`
		LastName  string    `json:"lastName"`
		PhoneNo   string    `json:"phoneNo"`
		Dob       time.Time `json:"dob"`
	}
	CreateAddressDto struct {
		CustomerId    string `json:"customerId"`
		StreetAddress string `json:"streetAddress"`
		City          string `json:"city"`
		State         string `json:"state"`
		PostalCodes   string `json:"postalCodes"`
	}
)

func (c CreateCustomerDto) Validate() error {
	err := validation.ValidateStruct(&c,
		validation.Field(&c.FirstName, validation.Required.Error("First name is required"), validation.Length(4, 20).Error("First name must be between 4 and 20 characters")),
		validation.Field(&c.LastName, validation.Required.Error("Last name is required"), validation.Length(4, 20).Error("Last name must be between 4 and 20 characters")),
		validation.Field(&c.PhoneNo, validation.Required.Error("Phone number is required"), validation.Length(10, 10).Error("Phone number must be 10 digits")),
		validation.Field(&c.Dob, validation.Required.Error("Date of birth is required")),
	)
	if err != nil {
		for _, err := range err.(validation.Errors) {
			return err
		}
	}
	return nil
}

func (c CreateAddressDto) Validate() error {
	err := validation.ValidateStruct(&c,
		validation.Field(&c.CustomerId, validation.Required.Error("customerId is required")),
		validation.Field(&c.City, validation.Required.Error("city is required")),
		validation.Field(&c.StreetAddress, validation.Required.Error("streetAddress is required")),
		validation.Field(&c.State, validation.Required.Error("state is required")),
		validation.Field(&c.PostalCodes, validation.Required.Error("postalCodes is required")),
	)
	if err != nil {
		for _, err := range err.(validation.Errors) {
			return err
		}
	}
	return nil
}
