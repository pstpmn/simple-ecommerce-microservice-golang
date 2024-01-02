package main

import (
	"fmt"
	"simple-ecomerce-microservice/services/customer"
)

func main() {
	customer.Handler()
	fmt.Println("Hello World This Is File main.go")
}
