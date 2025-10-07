package main

import (
	"crudproject/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	routers := router.Group("/crud")
	{
		routers.POST("/create", controllers.CreateHandler)
		routers.GET("/all", controllers.Readhandler)
		routers.PUT("/update/:id", controllers.Updatehandler)
		routers.DELETE("/delete/:id", controllers.DeleteHandler)
	}
	router.Run(":8888")
}
