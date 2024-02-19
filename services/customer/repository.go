package customer

import (
	"errors"
	"fmt"
	customerCore "simple-ecomerce-microservice/services/customer/core"
	customerModel "simple-ecomerce-microservice/services/customer/models"

	"gorm.io/gorm"
)

type (
	repo struct {
		db *gorm.DB
	}
)

// FindCustomer implements customerCore.ICustomerRepo.
func (ob *repo) FindCustomer(customerId string) (customerModel.CustomerModel, error) {
	var model customerModel.CustomerModel
	err := ob.
		db.
		Where(&customerModel.CustomerModel{Id: customerId}).
		Preload("Addresses").
		First(&model).
		Error
	if errors.Is(err, gorm.ErrRecordNotFound) {
		err = nil
	}
	return model, err
}

// FindCustomers implements customerCore.ICustomerRepo.
func (*repo) FindCustomers() ([]customerModel.CustomerModel, error) {
	panic("unimplemented")
}

// // FindCustomer implements customerCore.ICustomerRepo.
// func (ob *repo) FindCustomer(customerId string) (customerModel.CustomerModel, error) {
// }

// InsertOneAddress implements customerCore.ICustomerRepo.
func (ob *repo) InsertOneAddress(model *customerModel.AddressModel) error {
	if result := ob.db.Create(model).Error; result != nil {
		fmt.Println("Error function InsertOneAddress : %s", result.Error())
		return errors.New("somethin went wrong")
	}
	return nil
}

// InsertOnecustomerModel implements customerCore.ICustomerRepo.
func (ob *repo) InsertOneCustomer(model *customerModel.CustomerModel) error {
	if result := ob.db.Create(model).Error; result != nil {
		fmt.Println("Error function InsertOneCustomerModel : %s", result.Error())
		return errors.New("somethin went wrong")
	}
	return nil
}

func NewRepository(conn *gorm.DB) customerCore.ICustomerRepo {
	return &repo{
		db: conn,
	}
}
