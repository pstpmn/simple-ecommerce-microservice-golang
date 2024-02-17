package customerHanlder

import (
	"context"
	customerCore "simple-ecomerce-microservice/services/customer/core"
	pb "simple-ecomerce-microservice/services/customer/customerPb"
)

type (
	handler struct {
		usecase customerCore.ICustomerUseCase
		pb.UnimplementedCustomerServiceServer
	}
)

// VerifyCustomer implements pb.CustomerServiceServer.
func (ob *handler) VerifyCustomer(ctx context.Context, req *pb.VerifyCustomerReq) (*pb.VerifyCustomerRes, error) {
	_, err := ob.usecase.GetProfile(req.CustomerId)
	if err == nil {
		return &pb.VerifyCustomerRes{
			IsValid: true,
		}, nil
	}
	return &pb.VerifyCustomerRes{
		IsValid: false,
	}, err
}

// mustEmbedUnimplementedCustomerServiceServer implements pb.CustomerServiceServer.
// func (*handler) mustEmbedUnimplementedCustomerServiceServer() {
// 	panic("unimplemented")
// }

// func VerifyCustomer(customerId int) *pb.Ver
func NewCustomCustomerGrpcHandler(usecase customerCore.ICustomerUseCase) pb.CustomerServiceServer {
	return &handler{
		usecase: usecase,
	}
}
