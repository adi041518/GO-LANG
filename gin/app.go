package main

import (
	"example.com/dataaccess/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	books := router.Group("/home")
	{
		books.GET("/", controllers.ViewStudents)
		books.POST("/:id/:name", controllers.AddStudnets)
	}
	router.Run(":7777")
}
