package main

import (
	"os"
	"simple-ecomerce-microservice/pkg/database/gorm/migration"
)

func main() {
	service := os.Args[1]
	customer := migration.NewCustomerMigrate()
	product := migration.NewProductMigrate()
	switch service {
	case "customer":
		customer.AutoMigrate("postgres", "root", "localhost", "5432", "customer")
	case "product":
		product.AutoMigrate("postgres", "root", "localhost", "5432", "product")
	default:
		panic("not found service")
	}
}
