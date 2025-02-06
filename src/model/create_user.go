package model

import (
	"fmt"

	"github.com/jpirolla/CRUD_golang/src/configuration/logger"
	"github.com/jpirolla/CRUD_golang/src/configuration/rest_err"
	"go.uber.org/zap"
)

func (ud *UserDomain) CreateUser() *rest_err.RestErr {

	logger.Info("Init createUser controller", zap.String("journey", "create user"))

	ud.EncryptPassword()
	fmt.Println(ud)

	return nil
}
