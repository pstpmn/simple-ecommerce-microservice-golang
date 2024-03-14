package main

import (
	"fmt"
	"log"
	"simple-ecomerce-microservice/config"
	"simple-ecomerce-microservice/pkg"
	"simple-ecomerce-microservice/services/customer"
	customerCore "simple-ecomerce-microservice/services/customer/core"
	"simple-ecomerce-microservice/services/customer/customerPb"
	customerHanlder "simple-ecomerce-microservice/services/customer/handlers"

	mw "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"

	"net/http"

	"github.com/labstack/echo/v4"
)

func main() {
	var cfg config.Config
	viper.SetConfigFile("env/customer.yaml")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode config into struct, %v", err)
	}

	// connect mongodb
	postgres := pkg.NewGormORM()
	conn, err := postgres.ConnectDB(cfg.PostgresDb.Uri, "postgres")
	if err != nil {
		panic(err)
	}

	// prepare object
	helper := pkg.NewHelper()
	response := pkg.NewResponse()
	repo := customer.NewRepository(conn)
	usecase := customerCore.NewUseCase(repo, helper)
	httpHandler := customerHanlder.NewHttpHandler(repo, usecase, helper, response)
	grpcHandler := customerHanlder.NewCustomCustomerGrpcHandler(usecase)

	// create grpc app
	go func() {
		server, lis := pkg.NewGrpc().Server("", cfg.App.GrpcProt)
		customerPb.RegisterCustomerServiceServer(server, grpcHandler)

		fmt.Println("gRPC Server Running ...")
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Error: Failed run server grpc : %v", err)
		}
	}()

	// create http app
	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	r := e.Group("v1/customer")
	{
		r.GET("/:customerId", httpHandler.GetProfile)
		r.POST("", httpHandler.CreateCustomer)
		r.POST("/address", httpHandler.CreateAddress)
	}

	fmt.Println("Http Server Running ...")
	if err := e.Start(cfg.App.HttpProt); err != http.ErrServerClosed {
		panic(fmt.Errorf("error http server %s", err))
	}
}
