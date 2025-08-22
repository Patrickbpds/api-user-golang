package main

import (
	"context"
	"log"

	"api-user-golang/src/configuration/database/mongodb"
	"api-user-golang/src/configuration/logger"
	"api-user-golang/src/controller/routes"

	"github.com/gin-gonic/gin"
	"github.com/lpernett/godotenv"
)

func main() {
	logger.Info("Starting API User Service...")
	
	err := godotenv.Load()

	database, err := mongodb.NewMongoDbConnection(context.Background())
	if err != nil {
		log.Fatalf(
			"Error trying to connect to database, error=%s \n",
			err.Error())
		return
	}

	userController := initDependencies(database)

	router := gin.Default()
	gin.SetMode(gin.ReleaseMode)
	routes.InitRoutes(&router.RouterGroup, userController)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
