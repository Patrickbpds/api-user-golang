package model

type UserDomainInterface interface {
	GetEmail() string
	GetPassword() string
	GetName() string
	GetAge() int8
	GetId() string

	SetId(id string)
	EncryptPassword()
}

func NewUserDomain(
	email, password, name string, 
	age int8,
	) UserDomainInterface {
	return &userDomain{
		email:    email,
		password: password,
		name:     name,
		age:      age,
	}
}