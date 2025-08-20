package main

import (
	"fmt"
	"log"
	"os"

  "github.com/lpernett/godotenv"
  "github.com/gin-gonic/gin"
  "api-user-golang/src/controller/routes"
)

func main() {
  err := godotenv.Load()	
  if err != nil {
    log.Fatal("Error loading .env file")
  } 
  fmt.Println(os.Getenv("TEST"))

  router := gin.Default()

  routes.InitRoutes(&router.RouterGroup)
  if err := router.Run(":8080"); err != nil {
    log.Fatalf("Failed to start server: %v", err)
  }
}
