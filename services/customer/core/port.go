package customerCore

import (
	customerModel "simple-ecomerce-microservice/services/customer/models"
	"time"
)

type ICustomerUseCase interface {
	GetProfile(customerId string) (*CustomerProfile, error)
	CreateCustomer(firstName, lastName, phoneNo string, dob time.Time) (*Customer, error)
	CreateAddress(customerId, streetAddress, city, state, postalCodes string) (*Address, error)
	GetOrderHistory(customerId string) error
}

type ICustomerRepo interface {
	InsertOneCustomer(model *customerModel.CustomerModel) error
	InsertOneAddress(model *customerModel.AddressModel) error
	FindCustomer(customerId string) (customerModel.CustomerModel, error)
	FindCustomers() ([]customerModel.CustomerModel, error)
}

// type ICustomerHttpHandler interface {
// }

// type ICustomerGrpcHandler interface {
// 	mustEmbedUnimplementedGreeterServer()
// }

type IHelper interface {
	GenUuid() string
}
