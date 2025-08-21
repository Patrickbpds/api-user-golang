package controller

import (
	"api-user-golang/src/configuration/logger"
	"api-user-golang/src/configuration/validation"
	"api-user-golang/src/controller/dtos"

	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

func CreateUser(c *gin.Context) {
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

	response := dtos.UserResponse{
		Id:    "12345",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}
	logger.Info("User created successfully",
		zap.String("jornal_action", "create_user"),
	)
	c.JSON(200, response)
}
