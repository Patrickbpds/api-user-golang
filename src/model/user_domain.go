package model

import (
	rest_errors "api-user-golang/src/configuration/rest-errors"
	"crypto/md5"
	"encoding/hex"
)

func NewUserDomain(
	email, password, name string, 
	age int8,
	) UserDomainInterface {
	return &UserDomain{
		email, password, name, age,}
}


type UserDomain struct {
	Email    string
	Password string
	Name     string
	Age      int8
}

func (ud *UserDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.Password))
	ud.Password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_errors.RestError
	GetUser(string) (*UserDomain, *rest_errors.RestError)
	UpdateUser(string) *rest_errors.RestError
	DeleteUser(string) *rest_errors.RestError 
}

