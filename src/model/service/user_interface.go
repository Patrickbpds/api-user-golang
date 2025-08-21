package service

import (
	rest_errors "api-user-golang/src/configuration/rest-errors"
	"api-user-golang/src/model"
)


func NewUserDomainService() UserDomainService {
	return &userDomainService{}
}

type userDomainService struct {

}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) *rest_errors.RestError
	UpdateUser(string, model.UserDomainInterface) *rest_errors.RestError
	GetUser(string) (*model.UserDomainInterface, *rest_errors.RestError)
	DeleteUser(string) *rest_errors.RestError
}
