package service

import (
	"api-user-golang/src/configuration/logger"
	rest_errors "api-user-golang/src/configuration/rest-errors"
	"api-user-golang/src/model"

	"go.uber.org/zap"
)

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) (model.UserDomainInterface, *rest_errors.RestError) {
	
	logger.Info("Init createUser model", zap.String("jornal_action", "create_user"))

	userDomain.EncryptPassword()

	userDomainRepository, err := ud.userRepository.CreateUser(userDomain)
	if err != nil { return nil, err }

	return userDomainRepository, nil
}