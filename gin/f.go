package main

import (
	"example.com/dataaccess/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	r := router.Group("/home")
	{

		r.GET("/", controllers.TotalMembers)
	}
	router.Run(":8002")
}
