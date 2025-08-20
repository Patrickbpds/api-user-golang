package controller

import (
	rest_errors "api-user-golang/src/configuration/rest-errors"
	"api-user-golang/src/controller/dtos"
	"fmt"

	"github.com/gin-gonic/gin"
)

func CreateUser (c *gin.Context) {
	var userRequest dtos.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		restErr := rest_errors.NewBadRequestError(
			fmt.Sprintf("There are invalid fields in the request body, error=%s\n", err.Error()))
		c.JSON(restErr.Status, restErr)
	}
	fmt.Println("User request received: ", userRequest)
}