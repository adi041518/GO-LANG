package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.LoadHTMLGlob("html/*")
	router.GET("/", func(ctx *gin.Context) {
		ctx.HTML(200, "in.html", gin.H{
			"title": "Aditya",
		})
	})
	router.Run()
}
