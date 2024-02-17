package main

import (
	"fmt"
	"log"

	"simple-ecomerce-microservice/pkg"
	"simple-ecomerce-microservice/services/customer"
	customerCore "simple-ecomerce-microservice/services/customer/core"
	custPb "simple-ecomerce-microservice/services/customer/customerPb"
	customerHanlder "simple-ecomerce-microservice/services/customer/handlers"
	"simple-ecomerce-microservice/services/product"
	productCore "simple-ecomerce-microservice/services/product/core"
	productHandler "simple-ecomerce-microservice/services/product/handlers"
	prodPb "simple-ecomerce-microservice/services/product/productPb"
)

type grpcAuth struct {
	secretKey string
}

func main() {
	postgres := pkg.NewGormORM()
	dsn := "host=localhost user=postgres password=root dbname=customer port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	helper := pkg.NewHelper()
	conn, err := postgres.ConnectDB(dsn, "postgres")
	if err != nil {
		panic(err)
	}
	repo := customer.NewRepository(conn)
	usecase := customerCore.NewUseCase(repo, helper)
	server, lis := pkg.NewGrpc().Server("111", "0.0.0.0:11111")
	handler := customerHanlder.NewCustomCustomerGrpcHandler(usecase)
	custPb.RegisterCustomerServiceServer(server, handler)

	// pb.RegisterCustomerSecustomerrviceServer(server, handler)
	prodRepo := product.NewRepository(conn)
	prodUseCase := productCore.NewUseCase(prodRepo, helper)
	prodHandler := productHandler.NewProductGrpcHandler(prodUseCase)
	prodPb.RegisterProductServiceServer(server, prodHandler)

	fmt.Println("gRPC Server Running ...")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("Error: Failed run server grpc : %v", err)
	}
}
