package service

import (
	"github.com/jpirolla/CRUD_golang/src/configuration/rest_err"
	"github.com/jpirolla/CRUD_golang/src/model"
)

func (*userDomainService) FindUser(string) (
	*model.UserDomainInterface, *rest_err.RestErr) {
	return nil, nil
}
