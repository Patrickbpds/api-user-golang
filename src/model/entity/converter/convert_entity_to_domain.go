package converter

import (
	"api-user-golang/src/model"
	"api-user-golang/src/model/entity"
)


func ConverterEntityToDomain(entity entity.UserEntity) model.UserDomainInterface {
	domain := model.NewUserDomain(
		entity.Email, 
		entity.Password, 
		entity.Name, 
		entity.Age,
	)
	domain.SetId(entity.Id.Hex())
	return domain
}