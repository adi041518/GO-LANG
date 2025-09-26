package main

import (
	"github.com/gin-gonic/gin"
)

type person struct {
	Name string `json:"name" form:"name" xml:"name"`
	Id   int    `json:"id" form:"id" xml:"id"`
}

func main() {
	router := gin.Default()
	routers := router.Group("/home")
	{
		routers.POST("/json", func(ctx *gin.Context) {
			var p person
			ctx.Bind(&p)
			ctx.JSON(200, gin.H{
				"res": p,
			})
		})
		routers.POST("/xml", func(ctx *gin.Context) {
			var p person
			ctx.Bind(&p)
			ctx.JSON(200, gin.H{
				"res": p,
			})
		})
		routers.GET("/form", func(ctx *gin.Context) {
			var p person
			ctx.Bind(&p)
			ctx.JSON(200, gin.H{
				"res": p,
			})
		})
	}
	router.Run()
}
