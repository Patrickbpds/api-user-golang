package view

import (
	"api-user-golang/src/controller/dtos"
	"api-user-golang/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) dtos.UserResponse {
	return dtos.UserResponse{
		Id:    "",
		Name:  userDomain.GetName(),
		Email: userDomain.GetEmail(),
		Age:   userDomain.GetAge(),
	}
}