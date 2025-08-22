package repository

import (
	"api-user-golang/src/configuration/logger"
	rest_errors "api-user-golang/src/configuration/rest-errors"
	"api-user-golang/src/model"
	"context"
	"os"
)
var (
	COLLECTION_NAME = "MONGO_COLLECTION"
)

func (ur *userRepository) CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_errors.RestError) {

		logger.Info("Creating user in repository")
		
		collection_name := os.Getenv(COLLECTION_NAME)

		collection := ur.databaseConnection.Collection(collection_name)
		
		value, err := userDomain.GetJSONValue();
		if err != nil {
			return nil, rest_errors.NewInternalServerError("Error getting JSON value for user domain")
		}
		collection.InsertOne(context.Background(), value)

		result , err := collection.InsertOne(context.Background(), value)
		if err != nil {
			return nil, rest_errors.NewInternalServerError("Error getting JSON value for user domain")
		}

		userDomain.SetId(result.InsertedID.(string))
		logger.Info("User created successfully in repository")
		
		return userDomain, nil
	}