package main

import (
	"github.com/infinitss13/Cakes-store"
	"github.com/infinitss13/Cakes-store/config"
	_ "github.com/infinitss13/Cakes-store/docs"
	"github.com/infinitss13/Cakes-store/handlers"
	"github.com/infinitss13/Cakes-store/repository"
	"github.com/infinitss13/Cakes-store/services"
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
	configDb := config.NewDBConfig()
	database, err := repository.NewDatabase(configDb)
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
	server := new(cakes_store_user_service.Server)
	port := config.GetPortEnv()
	logrus.Info("Starting application on the port ", port)
	err = server.Run(port, handlers)
	if err != nil {
		logrus.Error(err)
		return
	}
}
