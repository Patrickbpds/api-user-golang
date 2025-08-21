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
	return &userDomain{
		email, password, name, age,}
}


type userDomain struct {
	email    string
	password string
	name     string
	age      int8
}

func (ud *userDomain) EncryptPassword() {
	hash := md5.New()
	defer hash.Reset()
	hash.Write([]byte(ud.password))
	ud.password = hex.EncodeToString(hash.Sum(nil))
}

type UserDomainInterface interface {
	CreateUser() *rest_errors.RestError
	GetUser(string) (*userDomain, *rest_errors.RestError)
	UpdateUser(string) *rest_errors.RestError
	DeleteUser(string) *rest_errors.RestError 
}

