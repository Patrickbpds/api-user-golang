package main

import (
	"api-user-golang/src/controller"
	"api-user-golang/src/model/repository"
	"api-user-golang/src/model/service"

	"go.mongodb.org/mongo-driver/mongo"
)

func initDependencies(
		database *mongo.Database,
) controller.UserControllerInterface {
	repo := repository.NewUserRepository(database)
	service := service.NewUserDomainService(repo)
	return controller.NewUserControllerInterface(service)
}