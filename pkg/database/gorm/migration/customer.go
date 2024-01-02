package migration

import (
	"fmt"
	"simple-ecomerce-microservice/pkg"
	customerModel "simple-ecomerce-microservice/services/customer/models"
)

type m struct{}

// AutoMigrate implements IMigration.
func (*m) AutoMigrate(user, pass, host, port, dbName string) {
	ob := pkg.NewGormORM()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		host, user, pass, dbName, port)
	db, err := ob.ConnectDB(dsn, "postgres")

	if err != nil {
		panic(err)
	}

	// init migrate auto create schema
	db.AutoMigrate(&customerModel.CustomerModel{}, &customerModel.AddressModel{})
}

func NewCustomerMigrate() IMigration {
	return &m{}
}
