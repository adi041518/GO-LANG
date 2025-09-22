package main

import (
	"github.com/gin-gonic/gin"
)

type Student struct {
	Name string `json:"name"`
}

func main() {
	router := gin.Default()
	var s Student
	router.POST("/", func(ctx *gin.Context) {
		ctx.BindJSON(&s)
		if s.Name == "Aditya" {
			ctx.String(200, "validation done thank you")
		}
	})
	router.Run(":8888")
}
"/services"