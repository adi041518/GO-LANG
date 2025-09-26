package main

import (
	"github.com/gin-gonic/gin"
)

type Use struct {
	Name string `form:"name" binding:"required"`
	City string `form:"city" binding:"required"`
}

func main() {
	router := gin.Default()

	router.Any("/", func(ctx *gin.Context) {
		var u Use
		ctx.Bind(&u)
		ctx.JSON(200, gin.H{
			"name": u.Name,
			"city": u.City,
		})
	})
	router.Run(":8888")
}
