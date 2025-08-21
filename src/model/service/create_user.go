package service

import (
	"api-user-golang/src/configuration/logger"
	rest_errors "api-user-golang/src/configuration/rest-errors"
	"api-user-golang/src/model"
	"fmt"

	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_errors.RestError {
	
	logger.Info("Init createUser model", zap.String("jornal_action", "create_user"))

	userDomain.EncryptPassword()

	fmt.Println("UserDomain:", userDomain.GetPassword())

	return nil
}