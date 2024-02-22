package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-ecomerce-microservice/config"
	"simple-ecomerce-microservice/pkg"
	"simple-ecomerce-microservice/services/product"
	productCore "simple-ecomerce-microservice/services/product/core"
	productHandler "simple-ecomerce-microservice/services/product/handlers"
	"simple-ecomerce-microservice/services/product/productPb"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	var cfg config.Config
	viper.SetConfigFile("env/product.yaml")
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
	repo := product.NewRepository(conn)
	usecase := productCore.NewUseCase(repo, helper)
	httpHandler := productHandler.NewHttpHandler(repo, usecase, helper, response)
	grpcHandler := productHandler.NewProductGrpcHandler(usecase)

	// create grpc app
	go func() {
		server, lis := pkg.NewGrpc().Server("", cfg.App.GrpcProt)
		productPb.RegisterProductServiceServer(server, grpcHandler)

		fmt.Println("gRPC Server Running ...")
		if err := server.Serve(lis); err != nil {
			log.Fatalf("Error: Failed run server grpc : %v", err)
		}
	}()

	// create http app
	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	r := e.Group("v1/product")
	{
		r.GET("", httpHandler.GetProducts)
		r.GET(":productId", httpHandler.GetProductDetail)
	}

	fmt.Println("Http Server Running ...")
	if err := e.Start(cfg.App.HttpProt); err != http.ErrServerClosed {
		panic(fmt.Errorf("error http server %s", err))
	}
}
