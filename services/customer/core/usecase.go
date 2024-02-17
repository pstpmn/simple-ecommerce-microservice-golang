package customerCore

import (
	"errors"
	customerModel "simple-ecomerce-microservice/services/customer/models"
	"time"
)

type (
	u struct {
		repo   ICustomerRepo
		helper IHelper
	}
)

func (ob *u) customerModelToDomain(model customerModel.CustomerModel) Customer {
	return Customer{
		Id:        model.Id,
		FirstName: model.FirstName,
		LastName:  model.LastName,
		PhoneNo:   model.PhoneNo,
		Status:    model.Status,
		Dob:       model.Dob,
		CreatedAt: model.CreatedAt,
	}
}

func (ob *u) customerModelToCustomerProfileDomain(model customerModel.CustomerModel) CustomerProfile {
	var addressList []Address
	for _, v := range model.Addresses {
		address := Address{
			Id:            v.Id,
			CustomerId:    v.CustomerId,
			StreetAddress: v.StreetAddress,
			City:          v.City,
			State:         v.State,
			PostalCodes:   v.PostalCodes,
			CreatedAt:     v.CreatedAt,
		}
		addressList = append(addressList, address)
	}

	return CustomerProfile{
		Customer: &Customer{
			Id:        model.Id,
			FirstName: model.FirstName,
			LastName:  model.LastName,
			PhoneNo:   model.PhoneNo,
			Status:    model.Status,
			Dob:       model.Dob,
			CreatedAt: model.CreatedAt,
		},
		Address: addressList,
	}
}

func (ob *u) addressModelToDomain(model customerModel.AddressModel) Address {
	return Address{
		Id:            model.Id,
		CustomerId:    model.CustomerId,
		StreetAddress: model.StreetAddress,
		City:          model.City,
		State:         model.State,
		PostalCodes:   model.PostalCodes,
		CreatedAt:     model.CreatedAt,
	}
}

// CreateAddress implements ICustomerUseCase.
func (ob *u) CreateAddress(customerId string, streetAddress string, city string, state string, postalCodes string) (*Address, error) {
	customer, err := ob.repo.FindCustomer(customerId)
	if err != nil {
		return nil, err
	} else if customer.Id == "" {
		return nil, errors.New("not found customer")
	}

	model := customerModel.AddressModel{
		Id:            ob.helper.GenUuid(),
		CustomerId:    customerId,
		StreetAddress: streetAddress,
		City:          city,
		State:         state,
		PostalCodes:   postalCodes,
		CreatedAt:     time.Now(),
	}
	err = ob.repo.InsertOneAddress(&model)
	domain := ob.addressModelToDomain(model)
	return &domain, err
}

// CreateCustomer implements ICustomerUseCase.
func (ob *u) CreateCustomer(firstName string, lastName string, phoneNo string, dob time.Time) (*Customer, error) {
	customerModel := &customerModel.CustomerModel{
		Id:        ob.helper.GenUuid(),
		FirstName: firstName,
		LastName:  lastName,
		PhoneNo:   phoneNo,
		Status:    true,
		Dob:       dob,
		CreatedAt: time.Now(),
	}
	err := ob.repo.InsertOneCustomer(customerModel)
	domain := ob.customerModelToDomain(*customerModel)
	return &domain, err
}

// GetOrderHistory implements ICustomerUseCase.
func (ob *u) GetOrderHistory(customerId string) error {
	panic("unimplemented")
}

// GetProfile implements ICustomerUseCase.
func (ob *u) GetProfile(customerId string) (*CustomerProfile, error) {
	customer, err := ob.repo.FindCustomer(customerId)
	if err != nil {
		return nil, err
	} else if customer.Id == "" {
		return nil, errors.New("not found customer")
	}
	domain := ob.customerModelToCustomerProfileDomain(customer)
	return &domain, err
}

func NewUseCase(repo ICustomerRepo, helper IHelper) ICustomerUseCase {
	return &u{
		repo:   repo,
		helper: helper,
	}
}
