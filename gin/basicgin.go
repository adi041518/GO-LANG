package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/home", func(c *gin.Context) {
		c.String(200, "Hi hello !!!!!")
	})
	router.Run()

}
