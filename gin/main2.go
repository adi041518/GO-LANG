package main

import (
	"exam.com/API-project/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", controllers.Printdisplay)
	router.Run()
}
