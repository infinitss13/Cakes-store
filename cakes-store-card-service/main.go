package main

import (
	"context"

	"github.com/infinitss13/Cakes-store-card-service/config"
	"github.com/infinitss13/Cakes-store-card-service/internal/handlers"
	"github.com/infinitss13/Cakes-store-card-service/internal/repository"
	"github.com/infinitss13/Cakes-store-card-service/internal/services"
	"github.com/sirupsen/logrus"
)

// @title Cakes store user-service
// @version 1.0
// @description API server for Cake Store user service

// @host localhost:8080
// @BasePath /

// @securityDefinitions.apiKey ApiKeyAuth
// @in header
// @name Authorization
func main() {
	ctx := context.Background()
	clientMongo, err := repository.NewClientMongo(ctx)
	if err != nil {
		logrus.Error("error in connection mongo : %w", err)
	}
	database := repository.NewDatabase(clientMongo)
	if err != nil {
		logrus.Error(err)
		return
	}
	service := services.NewService(database)
	handlers, err := handlers.SetRequestHandlers(service)
	if err != nil {
		logrus.Error(err)
		return
	}
	server := new(Server)
	port := config.GetPortEnv()
	logrus.Info("Starting AAAAAAAAAAAAAAAAAAAAAAAAA application on the port ", port)
	err = server.Run(port, handlers)
	if err != nil {
		logrus.Error(err)
		return
	}
}
