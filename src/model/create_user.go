package model

import (
	"api-user-golang/src/configuration/logger"
	rest_errors "api-user-golang/src/configuration/rest-errors"
	"fmt"

	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_errors.RestError {
	
	logger.Info("Init createUser model", zap.String("jornal_action", "create_user"))

	ud.EncryptPassword()

	fmt.Println("UserDomain:", ud)

	return nil
}