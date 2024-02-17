package main

import (
	"fmt"
	"log"

	"simple-ecomerce-microservice/pkg"
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
	dsn := "host=localhost user=postgres password=root dbname=product port=5432 sslmode=disable TimeZone=Asia/Bangkok"
	helper := pkg.NewHelper()
	conn, err := postgres.ConnectDB(dsn, "postgres")
	if err != nil {
		panic(err)
	}
	server, lis := pkg.NewGrpc().Server("111", "0.0.0.0:11112")

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
