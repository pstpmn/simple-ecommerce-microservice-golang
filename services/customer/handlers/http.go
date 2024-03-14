package customerHanlder

import (
	customerCore "simple-ecomerce-microservice/services/customer/core"

	nHttp "net/http"

	echo "github.com/labstack/echo/v4"
)

type (
	ICustomerHttp interface {
		GetProfile(c echo.Context) error
		CreateCustomer(c echo.Context) error
		CreateAddress(c echo.Context) error
		GetOrderHistory(c echo.Context) error
	}

	http struct {
		repo    customerCore.ICustomerRepo
		usecase customerCore.ICustomerUseCase
		helper  customerCore.IHelper
		reponse customerCore.IResponse
	}
)

// CreateAddress implements ICustomerHttp.
func (ob *http) CreateAddress(c echo.Context) error {
	var req CreateAddressDto
	c.Bind(&req)
	if err := req.Validate(); err != nil {
		return ob.reponse.Error(nHttp.StatusBadRequest, err.Error(), c)
	}
	result, err := ob.usecase.CreateAddress(req.CustomerId, req.StreetAddress, req.City, req.State, req.PostalCodes)
	if err != nil {
		return ob.reponse.Error(nHttp.StatusBadRequest, err.Error(), c)
	}
	return ob.reponse.Success(nHttp.StatusCreated, "successed", result, c)
}

// CreateCustomer implements ICustomerHttp.
func (ob *http) CreateCustomer(c echo.Context) error {
	var req CreateCustomerDto
	c.Bind(&req)
	if err := req.Validate(); err != nil {
		return ob.reponse.Error(nHttp.StatusBadRequest, err.Error(), c)
	}
	result, err := ob.usecase.CreateCustomer(req.FirstName, req.LastName, req.PhoneNo, req.Dob)
	if err != nil {
		return ob.reponse.Error(nHttp.StatusBadRequest, err.Error(), c)
	}
	return ob.reponse.Success(nHttp.StatusCreated, "successed", result, c)
}

// GetOrderHistory implements ICustomerHttp.
func (ob *http) GetOrderHistory(c echo.Context) error {
	result, err := ob.usecase.GetProfile(c.Param("customerId"))
	if err != nil {
		return ob.reponse.Error(nHttp.StatusBadRequest, err.Error(), c)
	}
	return ob.reponse.Success(nHttp.StatusOK, "successed", result, c)
}

// GetProfile implements ICustomerHttp.
func (ob *http) GetProfile(c echo.Context) error {
	result, err := ob.usecase.GetProfile(c.Param("customerId"))
	if err != nil {
		return ob.reponse.Error(nHttp.StatusBadRequest, err.Error(), c)
	}
	return ob.reponse.Success(nHttp.StatusOK, "successed", result, c)
}

func NewHttpHandler(repo customerCore.ICustomerRepo, usecase customerCore.ICustomerUseCase, helper customerCore.IHelper, response customerCore.IResponse) ICustomerHttp {
	return &http{
		repo:    repo,
		usecase: usecase,
		helper:  helper,
		reponse: response,
	}
}
