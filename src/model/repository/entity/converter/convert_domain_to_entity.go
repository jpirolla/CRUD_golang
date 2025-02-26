package converter

import (
	"github.com/jpirolla/CRUD_golang/src/model"
	"github.com/jpirolla/CRUD_golang/src/model/repository/entity"
)

func ConvertDomainToEntity(
	domain model.UserDomainInterface,
) *entity.UserEntity {
	return &entity.UserEntity{
		Email:    domain.GetEmail(),
		Password: domain.GetPassword(),
		Name:     domain.GetName(),
		Age:      domain.GetAge(),
	}
}
