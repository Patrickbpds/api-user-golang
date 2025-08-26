package repository

import (
	"api-user-golang/src/configuration/logger"
	rest_errors "api-user-golang/src/configuration/rest-errors"
	"api-user-golang/src/model"
	"context"
	"os"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.uber.org/zap"
	"api-user-golang/src/model/entity/converter"
)
var (
	COLLECTION_NAME = "MONGO_COLLECTION"
)

func (ur *userRepository) CreateUser(
		userDomain model.UserDomainInterface,
	) (model.UserDomainInterface, *rest_errors.RestError) {

		logger.Info("Creating user in repository", zap.String("jornal_action", "createUser"))
		
		collection_name := os.Getenv(COLLECTION_NAME)

		collection := ur.databaseConnection.Collection(collection_name)
	
		value := converter.ConverterDomainToEntity(userDomain)

		result, err := collection.InsertOne(context.Background(), value)
		if err != nil {
		logger.Error("Error trying to create user",
			err,
			zap.String("journey", "createUser"))
		return nil, rest_errors.NewInternalServerError(err.Error())
	}

		value.Id = result.InsertedID.(primitive.ObjectID)

		logger.Info(
		"CreateUser repository executed successfully",
		zap.String("userId", value.Id.Hex()),
		zap.String("journey", "createUser"))

		return converter.ConverterEntityToDomain(*value), nil
	}