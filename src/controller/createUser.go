package controller

import (
	"api-user-golang/src/configuration/validation"
	"api-user-golang/src/controller/dtos"
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
)

func CreateUser(c *gin.Context) {
	var userRequest dtos.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("There are invalid fields in the request body, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)
		c.JSON(errRest.Status, errRest)
		return
	}
	fmt.Printf("UserRequest: %+v\n", userRequest)
	response := dtos.UserResponse{
		Id:    "12345",
		Name:  userRequest.Name,
		Email: userRequest.Email,
		Age:   userRequest.Age,
	}
	c.JSON(200, response)
}
