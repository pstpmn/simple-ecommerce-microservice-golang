package main

import (
	"fmt"
	"log"
	"net/http"
	"simple-ecomerce-microservice/config"
	"simple-ecomerce-microservice/pkg"
	"simple-ecomerce-microservice/services/order"
	orderCore "simple-ecomerce-microservice/services/order/core"
	orderHandler "simple-ecomerce-microservice/services/order/handlers"

	"github.com/labstack/echo/v4"
	mw "github.com/labstack/echo/v4/middleware"
	"github.com/spf13/viper"
)

func main() {
	var cfg config.Config
	viper.SetConfigFile("env/order.yaml")
	viper.SetConfigType("yaml")
	viper.AutomaticEnv()

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file: %v", err)
	}

	if err := viper.Unmarshal(&cfg); err != nil {
		log.Fatalf("Unable to decode config into struct, %v", err)
	}

	// connect mongodb
	mongo := pkg.NewMongo()
	fmt.Println(cfg.MongoDb.Uri)
	conn, err := mongo.Connect(cfg.MongoDb.Uri)
	if err != nil {
		panic(err)
	}

	grpc := pkg.NewGrpc()

	// prepare object
	helper := pkg.NewHelper()
	response := pkg.NewResponse()
	repo := order.NewRepository(conn, grpc)
	usecase := orderCore.NewUseCase(repo, helper, cfg)
	httpHandler := orderHandler.NewHttpHandler(repo, usecase, helper, response)

	// create http app
	e := echo.New()
	e.Use(mw.Logger())
	e.Use(mw.Recover())
	r := e.Group("v1/order")
	{
		r.POST("", httpHandler.CreateOrder)
		r.GET(":orderId", httpHandler.GetOrderDetail)
		r.DELETE(":orderId", httpHandler.CancelOrder)
	}

	fmt.Println("Http Server Running ...")
	if err := e.Start(cfg.App.HttpProt); err != http.ErrServerClosed {
		panic(fmt.Errorf("error http server %s", err))
	}
}
