package view

import (
	"github.com/jpirolla/CRUD_golang/src/controller/model/response"
	"github.com/jpirolla/CRUD_golang/src/model"
)

func ConvertDomainToResponse(
	userDomain model.UserDomainInterface,
) response.UserResponse {
	return response.UserResponse{
		ID:    "",
		Email: userDomain.GetEmail(),
		Name:  userDomain.GetName(),
		Age:   userDomain.GetAge(),
	}
}
