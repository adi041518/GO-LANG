package main

import (
	"example1.com/user-registartion-api/controllers"
	"example1.com/user-registartion-api/services"
	"github.com/gin-gonic/gin"
)

func main() {

	router := gin.Default()
	serv := services.NewServices()
	con := controllers.NewController(serv)
	routers := router.Group("/home")
	{
		routers.POST("/user", con.UserValidateController)
	}
	router.Run()

}
