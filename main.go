package main

import (
	"log"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/jpirolla/CRUD_golang/src/controller/routes"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// preciso instanciar um router
	// gin.New criará um novo roteador gin sem nenhum middleware anexado. O gin.Default criará um novo roteador gin com logger e e middlewares de recovery anexados por padrão.
	router := gin.Default()

	routes.InitRoutes(&router.RouterGroup)

	if err := router.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}
