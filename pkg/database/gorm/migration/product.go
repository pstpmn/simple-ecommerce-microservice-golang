package migration

import (
	"fmt"
	"simple-ecomerce-microservice/pkg"
	productModel "simple-ecomerce-microservice/services/product/models"
)

type p struct{}

// AutoMigrate implements IMigration.
func (*p) AutoMigrate(user, pass, host, port, dbName string) {
	ob := pkg.NewGormORM()
	dsn := fmt.Sprintf(
		"host=%s user=%s password=%s dbname=%s port=%s sslmode=disable TimeZone=Asia/Bangkok",
		host, user, pass, dbName, port)
	db, err := ob.ConnectDB(dsn, "postgres")

	if err != nil {
		panic(err)
	}

	// init migrate auto create schema
	db.AutoMigrate(&productModel.ProductModel{}, &productModel.CategoryModel{})
}

func NewProductMigrate() IMigration {
	return &p{}
}
