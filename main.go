package main

import (
	"log"

	"github.com/carvalhocaio/go-crud/src/configuration/logger"
	"github.com/carvalhocaio/go-crud/src/controller"
	"github.com/carvalhocaio/go-crud/src/controller/routes"
	"github.com/carvalhocaio/go-crud/src/model/service"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	logger.Info("About to start user application")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Init dependencies
	service := service.NewUserDomainService()
	userController := controller.NewUserControllerInterface(service)

	router := gin.Default()
	routes.InitRoutes(&router.RouterGroup, userController)
	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
