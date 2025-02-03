package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpirolla/CRUD_golang/src/controller"
)

func InitRoutes(r *gin.RouterGroup) {
	// como quero receber um parametro, coloco : na variavel userId na rota
	r.GET("/getUserById/:userId", controller.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", controller.FindUserByEmail)
	r.POST("/createUser", controller.CreateUser)
	r.PUT("/updateUser/:userId", controller.UpdateUser)
	r.DELETE("/deletUser/:userId", controller.DeleteUser)
}
