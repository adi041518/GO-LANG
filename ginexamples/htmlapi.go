package main

import (
	"github.com/gin-gonic/gin"
)

type form struct {
	Colors []string `form:"colors[]"`
}

func main() {
	router := gin.Default()

	router.LoadHTMLGlob("html/*.html")

	router.POST("/", func(ctx *gin.Context) {
		var f form
		ctx.Bind(&f)
		ctx.JSON(200, gin.H{
			"result": f.Colors,
		})
	})
	router.Run(":8585")
}
