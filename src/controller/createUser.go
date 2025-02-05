package controller

import (
	"fmt"
	"log"

	"github.com/gin-gonic/gin"
	"github.com/jpirolla/CRUD_golang/src/configuration/validation"
	"github.com/jpirolla/CRUD_golang/src/controller/model/request"
)

func CreateUser(c *gin.Context) {
	log.Println("Init createUser controller")
	var userRequest request.UserRequest

	if err := c.ShouldBindJSON(&userRequest); err != nil {
		log.Printf("Error trying to marshal object, error=%s\n", err.Error())
		errRest := validation.ValidateUserError(err)

		c.JSON(errRest.Code, errRest)
		return
	}

	fmt.Println((userRequest))
}
