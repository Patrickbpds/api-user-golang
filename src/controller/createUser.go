package controller

import (
	"api-user-golang/src/configuration/logger"
	"api-user-golang/src/configuration/validation"
	"api-user-golang/src/controller/dtos"
	"api-user-golang/src/model"
	"api-user-golang/src/view"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	UserDomainInterface model.UserDomainInterface
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Initializing CreateUser controller",
		zap.String("jornal_action", "create_user"),
	)
	var userRequest dtos.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user request: %v", err,
			zap.String("jornal_action", "create_user"),
		)
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Status, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)
	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Status, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("jornal_action", "create_user"),
	)
	c.JSON(200, view.ConvertDomainToResponse(domain))
}
