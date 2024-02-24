package orderHandler

import (
	orderCore "simple-ecomerce-microservice/services/order/core"

	nHttp "net/http"

	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	IOrderHttp interface {
		GetOrderDetail(c echo.Context) error
		CreateOrder(c echo.Context) error
		CancelOrder(c echo.Context) error
	}

	http struct {
		repo    orderCore.IOrderRepo
		usecase orderCore.IOrderUseCase
		helper  orderCore.IHelper
		reponse orderCore.IResponse
	}
)

// CancelOrder implements IOrderHttp.
func (ob *http) CancelOrder(c echo.Context) error {
	var orderObjId primitive.ObjectID
	ob.helper.ConvertStrToPrimitiveObjectId(c.Param("orderId"), &orderObjId)
	err := ob.usecase.CancelOrder(c.Request().Context(), c.Param("customerId"), orderObjId)
	if err != nil {
		return ob.reponse.Error(nHttp.StatusBadRequest, err.Error(), c)
	}
	return ob.reponse.Success(nHttp.StatusOK, "successed", nil, c)
}

// CreateOrder implements IOrderHttp.
func (ob *http) CreateOrder(c echo.Context) error {
	var orderObjId primitive.ObjectID
	ob.helper.ConvertStrToPrimitiveObjectId(c.Param("orderId"), &orderObjId)
	result, err := ob.usecase.CreateOrder(c.Request().Context(), c.Param("customerId"), []orderCore.OrderDetail{})
	if err != nil {
		return ob.reponse.Error(nHttp.StatusBadRequest, err.Error(), c)
	}
	return ob.reponse.Success(nHttp.StatusOK, "successed", result, c)
}

// GetOrderDetail implements IOrderHttp.
func (ob *http) GetOrderDetail(c echo.Context) error {
	var orderObjId primitive.ObjectID
	ob.helper.ConvertStrToPrimitiveObjectId(c.Param("orderId"), &orderObjId)
	result, err := ob.usecase.GetOrderDetail(c.Request().Context(), orderObjId)
	if err != nil {
		return ob.reponse.Error(nHttp.StatusBadRequest, err.Error(), c)
	}
	return ob.reponse.Success(nHttp.StatusOK, "successed", result, c)
}

func NewHttpHandler(repo orderCore.IOrderRepo, usecase orderCore.IOrderUseCase, helper orderCore.IHelper, response orderCore.IResponse) IOrderHttp {
	return &http{
		repo:    repo,
		usecase: usecase,
		helper:  helper,
		reponse: response,
	}
}
