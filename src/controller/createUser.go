package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpirolla/CRUD_golang/src/configuration/logger"
	"github.com/jpirolla/CRUD_golang/src/configuration/validation"
	"github.com/jpirolla/CRUD_golang/src/controller/model/request"
	"github.com/jpirolla/CRUD_golang/src/model"
	"github.com/jpirolla/CRUD_golang/src/view"
	"go.uber.org/zap"
)

func (uc *userControllerInterface) CreateUser(c *gin.Context) {
	logger.Info("Init createUser controller",
		zap.String("journey", "create user"),
	)
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		logger.Error("Error trying to validate user info", err, zap.String("journey", "create user"))
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	domain := model.NewUserDomain(
		userRequest.Email,
		userRequest.Password,
		userRequest.Name,
		userRequest.Age,
	)

	// controler nao vai ter acesso aos dados, mas sim às funções
	// que vão acessar os dados
	//r etorna um rest Error
	if err := uc.service.CreateUser(domain); err != nil {
		c.JSON(err.Code, err)
		return
	}

	logger.Info("User created successfully",
		zap.String("journey", "createUser"))

	c.JSON(http.StatusOK, view.ConvertDomainToResponse(domain))

}
