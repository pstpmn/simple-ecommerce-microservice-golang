package orderCore

import (
	"context"
	"net"
	custPb "simple-ecomerce-microservice/services/customer/customerPb"
	orderModel "simple-ecomerce-microservice/services/order/models"
	prodPb "simple-ecomerce-microservice/services/product/productPb"

	echo "github.com/labstack/echo/v4"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"google.golang.org/grpc"
)

type IOrderUseCase interface {
	GetOrderDetail(pctx context.Context, orderId primitive.ObjectID) (*OrderProfile, error)
	CreateOrder(pctx context.Context, customerId string, order []OrderDetail) (*OrderProfile, error)
	CancelOrder(pctx context.Context, customerId string, orderId primitive.ObjectID) error
}

type IOrderRepo interface {
	// grpc service
	VerifyCustomerById(pctx context.Context, grpcUrl string, req *custPb.VerifyCustomerReq) (*custPb.VerifyCustomerRes, error)
	GetProductDetail(pctx context.Context, grpcUrl string, req *prodPb.ProductId) (*prodPb.ProductProfile, error)
	StockManager(pctx context.Context, grpcUrl string, req *prodPb.StockManagerReq) (*prodPb.ProductProfile, error)

	FindOrders(pctx context.Context) (*[]orderModel.Order, error)
	FindOneOrderByOrderId(pctx context.Context, orderId primitive.ObjectID) (*orderModel.Order, error)
	FindOrderDetailsByOrderId(pctx context.Context, orderId primitive.ObjectID) ([]*orderModel.OrderDetail, error)
	InserOneOrder(pctx context.Context, model orderModel.Order) (orderModel.Order, error)
	InserOrderDetail(pctx context.Context, model []orderModel.OrderDetail) ([]orderModel.OrderDetail, error)
	UpdateOneOrderStatus(pctx context.Context, orderId primitive.ObjectID, status string) (*orderModel.Order, error)
}

type IHelper interface {
	ConvertStrToPrimitiveObjectId(value string, target *primitive.ObjectID) error
	GenUuid() string
}

type IGrpc interface {
	Server(key, host string) (*grpc.Server, net.Listener)
	Client(host string) (*grpc.ClientConn, error)
}

type IResponse interface {
	Error(statusCode int, message string, ctx echo.Context) error
	Success(statusCode int, message string, result any, ctx echo.Context) error
}
