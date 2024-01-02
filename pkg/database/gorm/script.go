package main

import (
	"os"
	"simple-ecomerce-microservice/pkg/database/gorm/migration"
)

func main() {
	service := os.Args[1]
	customer := migration.NewCustomerMigrate()
	switch service {
	case "customer":
		customer.AutoMigrate("postgres", "root", "localhost", "5432", "customer")
	default:
		panic("not found service")
	}
}
