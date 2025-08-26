package service

import (
	rest_errors "api-user-golang/src/configuration/rest-errors"
	"api-user-golang/src/model"
	"api-user-golang/src/model/repository"
)


func NewUserDomainService(userRepository repository.UserRepository,
) UserDomainService {
	return &userDomainService{userRepository}
}

type userDomainService struct {
	userRepository repository.UserRepository

}

type UserDomainService interface {
	CreateUser(model.UserDomainInterface) (model.UserDomainInterface, *rest_errors.RestError)
	UpdateUser(string, model.UserDomainInterface) *rest_errors.RestError
	GetUser(string) (*model.UserDomainInterface, *rest_errors.RestError)
	DeleteUser(string) *rest_errors.RestError
}
