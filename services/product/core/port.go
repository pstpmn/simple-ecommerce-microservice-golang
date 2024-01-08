package productCore

import (
	productModel "simple-ecomerce-microservice/services/product/models"
)

type IProductUseCase interface {
	GetProducts() (*[]ProductIntroduction, error)
	GetProductDetail(productId string) (*ProductProfile, error)

	StockManager(productId, tobe string, amount int64) (*ProductProfile, error)
	// IncreseStock(productId string, amountIncrese int64) (*ProductProfile, error)
	// DecreseStock(productId string, amountDecrese int64) (*ProductProfile, error)
}

type IProductRepo interface {
	FindProduct(productId string) (productModel.ProductModel, error)
	UpdateOne(model productModel.ProductModel) error
	FindProducts() ([]productModel.ProductModel, error)
}

type IHelper interface {
	GenUuid() string
}
