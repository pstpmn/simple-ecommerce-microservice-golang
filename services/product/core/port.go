package productCore

import (
	"context"
	productModel "simple-ecomerce-microservice/services/product/models"

	"github.com/labstack/echo/v4"
)

type IProductUseCase interface {
	GetProducts(pctx context.Context) (*[]ProductIntroduction, error)
	GetProductDetail(pctx context.Context, productId string) (*ProductProfile, error)

	StockManager(pctx context.Context, productId, tobe string, amount int64) (*ProductProfile, error)
	// IncreseStock(productId string, amountIncrese int64) (*ProductProfile, error)
	// DecreseStock(productId string, amountDecrese int64) (*ProductProfile, error)
}

type IProductRepo interface {
	FindProduct(pctx context.Context, productId string) (productModel.ProductModel, error)
	UpdateOne(pctx context.Context, model productModel.ProductModel) error
	FindProducts(pctx context.Context) ([]productModel.ProductModel, error)
}

type IHelper interface {
	GenUuid() string
}
type IResponse interface {
	Error(statusCode int, message string, ctx echo.Context) error
	Success(statusCode int, message string, result any, ctx echo.Context) error
}
