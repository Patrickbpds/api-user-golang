package repository

import (
	rest_errors "api-user-golang/src/configuration/rest-errors"
	"api-user-golang/src/model"

	"go.mongodb.org/mongo-driver/mongo"
)

func NewUserRepository(
	database *mongo.Database,
) UserRepository {
	return &userRepository{
		database,
	}
}

type userRepository struct{
	databaseConnection *mongo.Database
}

type UserRepository interface {
	CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_errors.RestError)
}
