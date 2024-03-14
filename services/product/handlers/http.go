package productHandler

import (
	productCore "simple-ecomerce-microservice/services/product/core"

	nHttp "net/http"

	echo "github.com/labstack/echo/v4"
)

type (
	IProductHttp interface {
		GetProducts(c echo.Context) error
		GetProductDetail(c echo.Context) error
		// StockManager(c echo.Context) error
	}

	http struct {
		repo    productCore.IProductRepo
		usecase productCore.IProductUseCase
		helper  productCore.IHelper
		reponse productCore.IResponse
	}
)

// GetProductDetail implements ICustomerHttp.
func (ob *http) GetProductDetail(c echo.Context) error {
	result, err := ob.usecase.GetProductDetail(c.Request().Context(), c.Param("productId"))
	if err != nil {
		return ob.reponse.Error(nHttp.StatusBadRequest, err.Error(), c)
	}
	return ob.reponse.Success(nHttp.StatusOK, "successed", result, c)
}

// GetProducts implements ICustomerHttp.
func (ob *http) GetProducts(c echo.Context) error {
	result, err := ob.usecase.GetProducts(c.Request().Context())
	if err != nil {
		return ob.reponse.Error(nHttp.StatusBadRequest, err.Error(), c)
	}
	return ob.reponse.Success(nHttp.StatusOK, "successed", result, c)
}

func NewHttpHandler(repo productCore.IProductRepo, usecase productCore.IProductUseCase, helper productCore.IHelper, response productCore.IResponse) IProductHttp {
	return &http{
		repo:    repo,
		usecase: usecase,
		helper:  helper,
		reponse: response,
	}
}
