package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/jpirolla/CRUD_golang/src/controller"
)

func InitRoutes(
	r *gin.RouterGroup,
	userController controller.UserControllerInterface) {

	// como quero receber um parametro, coloco : na variavel userId na rota
	r.GET("/getUserById/:userId", userController.FindUserByID)
	r.GET("/getUserByEmail/:userEmail", userController.FindUserByEmail)
	r.POST("/createUser", userController.CreateUser)
	r.PUT("/updateUser/:userId", userController.UpdateUser)
	r.DELETE("/deletUser/:userId", userController.DeleteUser)
}
