package orderCore

import (
	"context"
	"errors"
	"fmt"
	"simple-ecomerce-microservice/config"
	custPb "simple-ecomerce-microservice/services/customer/customerPb"
	orderModel "simple-ecomerce-microservice/services/order/models"
	prodPb "simple-ecomerce-microservice/services/product/productPb"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type (
	u struct {
		repo   IOrderRepo
		helper IHelper
		config config.Config
	}
)

func (ob *u) convertModelToProfileDomin(order orderModel.Order, details []orderModel.OrderDetail) OrderProfile {
	return OrderProfile{
		Order: &Order{
			OrderId:         order.OrderId.Hex(),
			CustomerId:      order.CustomerId,
			TotalAmount:     0,
			Status:          order.Status,
			ShippingAddress: order.ShippingAddress,
			OrderAt:         order.OrderAt,
			CreatedAt:       order.CreatedAt,
		},
		Details: func() []OrderDetail {
			var result []OrderDetail
			for _, v := range details {
				result = append(result, OrderDetail{
					OrderDetailId: v.OrderDetailId.Hex(),
					OrderId:       v.OrderDetailId.Hex(),
					ProductId:     v.ProductId,
					Quantity:      v.Quantity,
					Price:         v.Price,
					CreatedAt:     v.CreatedAt,
				})
			}
			return result
		}(),
	}
}

func (ob u) validateProductStock(reqAmt int16, remainingReq int64) bool {
	if remainingReq < int64(reqAmt) {
		return false
	}
	return true
}

// CreateOrder implements IOrderUseCase.
func (ob *u) CreateOrder(pctx context.Context, customerId string, products []OrderDetail) (*OrderProfile, error) {
	// -> flow algorithm <-
	// verify customerId by grpc
	// verify and validate quantity for each products by grpc
	// save order
	// save product
	// -> end flow <-

	// verify customer
	if isValid, err := ob.repo.VerifyCustomerById(pctx, ob.config.Grpc.Customer, &custPb.VerifyCustomerReq{
		CustomerId: customerId,
	}); err != nil {
		return nil, err
	} else if isValid.IsValid == false {
		return nil, errors.New("not found customerId")
	}

	// validate product
	for _, v := range products {
		prod, err := ob.repo.GetProductDetail(pctx, ob.config.Grpc.Product, &prodPb.ProductId{
			ProductId: v.ProductId,
		})
		if err != nil {
			return nil, err
		}
		if ob.validateProductStock(v.Quantity, prod.Product.Stock) == false {
			return nil, errors.New(fmt.Sprintf("product name : %s quantity is not enough", prod.Product.Name))
		}

		// decrease product in stockc
		_, err = ob.repo.StockManager(pctx, ob.config.Grpc.Product, &prodPb.StockManagerReq{
			ProductId: v.ProductId,
			Topic:     "sub",
			Amount:    int64(v.Quantity),
		})
		if err != nil {
			return nil, err
		}
	}

	// create main order
	order, err := ob.repo.InserOneOrder(pctx, orderModel.Order{
		CustomerId:      customerId,
		Status:          "wait",
		ShippingAddress: "",
		OrderAt:         time.Now(),
		CreatedAt:       time.Now(),
	})
	if err != nil {
		return nil, err
	}

	// convert to model list
	var model []orderModel.OrderDetail
	for _, v := range products {
		model = append(model, orderModel.OrderDetail{
			OrderId:   order.OrderId,
			ProductId: v.ProductId,
			Quantity:  v.Quantity,
			Price:     v.Price,
			CreatedAt: v.CreatedAt,
		})
	}

	// create order details
	result, err := ob.repo.InserOrderDetail(pctx, model)
	domain := ob.convertModelToProfileDomin(order, result)
	return &domain, err
}

func (ob *u) CancelOrder(pctx context.Context, customerId string, orderId primitive.ObjectID) error {
	if _, err := ob.repo.VerifyCustomerById(pctx, ob.config.Grpc.Customer, &custPb.VerifyCustomerReq{CustomerId: customerId}); err != nil {
		return err
	}

	order, err := ob.repo.FindOneOrderByOrderId(pctx, orderId)
	if err != nil {
		return err
	} else if order.Status == "cancel" {
		return errors.New("the order canceled already")
	}

	if order.CustomerId != customerId {
		return errors.New("the customer is not the owner of the order")
	}

	detailsInOrder, err := ob.repo.FindOrderDetailsByOrderId(pctx, orderId)
	if err != nil {
		return err
	}

	if _, err = ob.repo.UpdateOneOrderStatus(pctx, orderId, "cancel"); err != nil {
		return err
	}

	for _, detail := range detailsInOrder {
		_, err := ob.repo.StockManager(pctx, ob.config.Grpc.Product, &prodPb.StockManagerReq{
			ProductId: detail.ProductId,
			Topic:     "add",
			Amount:    int64(detail.Quantity),
		})
		if err != nil {
			fmt.Println("error returning quantity for order:", err) // FIXME: log the error
			return err
		}
	}
	return nil
}

// GetOrderDetail implements IOrderUseCase.
func (ob *u) GetOrderDetail(pctx context.Context, orderId primitive.ObjectID) (profile *OrderProfile, err error) {
	order, err := ob.repo.FindOneOrderByOrderId(pctx, orderId)
	if err != nil {
		return
	}
	details, err := ob.repo.FindOrderDetailsByOrderId(pctx, orderId)
	if err != nil {
		return
	}
	var dereferencedDetails []orderModel.OrderDetail
	for _, detail := range details {
		dereferencedDetails = append(dereferencedDetails, *detail)
	}
	profile = func() *OrderProfile {
		r := ob.convertModelToProfileDomin(*order, dereferencedDetails)
		return &r
	}()
	return profile, nil
}

func NewUseCase(repo IOrderRepo, helper IHelper, config config.Config) IOrderUseCase {
	return &u{
		repo:   repo,
		helper: helper,
		config: config,
	}
}
