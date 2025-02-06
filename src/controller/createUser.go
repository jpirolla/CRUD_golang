package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jpirolla/CRUD_golang/src/configuration/logger"
	"github.com/jpirolla/CRUD_golang/src/configuration/validation"
	"github.com/jpirolla/CRUD_golang/src/controller/model/request"
	"github.com/jpirolla/CRUD_golang/src/controller/model/response"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
)

func CreateUser(c *gin.Context) {
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

	response := response.UserResponse{
		ID:    "test",
		Email: userRequest.Email,
		Name:  userRequest.Name,
		Age:   userRequest.Age,
	}

	logger.Info("User created successfully",
		zapcore.Field{
			Key:    "journey",
			String: "create user",
		},
	)

	c.JSON(http.StatusOK, response)

}
