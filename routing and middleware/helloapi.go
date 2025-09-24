package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	router.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"res": "Success",
		})
	})
	router.GET("/hello/:name", func(ctx *gin.Context) {
		c := ctx.Param("name")
		ctx.String(200, c)
	})
	router.Run(":8520")
}
