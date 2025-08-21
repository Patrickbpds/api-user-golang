package main

import (
	"fmt"
	"log"
	"os"

	"api-user-golang/src/configuration/logger"
	"api-user-golang/src/controller"
	"api-user-golang/src/controller/routes"
	"api-user-golang/src/model/service"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func main() {
  logger.Info("Starting API User Service...")
  err := godotenv.Load()	
  if err != nil {
    log.Fatal("Error loading .env file")
  } 
  fmt.Println(os.Getenv("TEST"))

  service := service.NewUserDomainService()
  userController := controller.NewUserControllerInterface(service)

  router := gin.Default()

  routes.InitRoutes(&router.RouterGroup, userController)
  if err := router.Run(":8080"); err != nil {
    log.Fatalf("Failed to start server: %v", err)
  }
}
