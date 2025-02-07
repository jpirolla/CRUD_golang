package service

import (
	"fmt"

	"github.com/jpirolla/CRUD_golang/src/configuration/logger"
	"github.com/jpirolla/CRUD_golang/src/configuration/rest_err"
	"github.com/jpirolla/CRUD_golang/src/model"
	"go.uber.org/zap"
)

// para dependencias externas

func (ud *userDomainService) CreateUser(
	userDomain model.UserDomainInterface,
) *rest_err.RestErr {

	logger.Info("Init createUser controller", zap.String("journey", "createUser"))

	userDomain.EncryptPassword()
	fmt.Println(userDomain.GetPassword())

	return nil
}
