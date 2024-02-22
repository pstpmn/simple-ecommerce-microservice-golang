package customerCore

import (
	customerModel "simple-ecomerce-microservice/services/customer/models"
	"time"

	"github.com/labstack/echo/v4"
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

type IResponse interface {
	Error(statusCode int, message string, ctx echo.Context) error
	Success(statusCode int, message string, result any, ctx echo.Context) error
}
