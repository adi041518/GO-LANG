package main

import (
	"github.com/gin-gonic/gin"
)

type Users struct {
	Name string `json:"name"`
	Id   string `json:"id`
}

func main() {
	router := gin.Default()

	var u Users

	router.POST("/todo", func(ctx *gin.Context) {
		ctx.BindJSON(&u)
		ctx.JSON(200, gin.H{
			"name": u.Name,
			"id":   u.Id,
		})
	})
	router.PUT("/todo/:no", func(ctx *gin.Context) {
		number := ctx.Param("no")
		ctx.BindJSON(&u)
		ctx.JSON(200, gin.H{
			"Number": number,
			"name":   u.Name,
			"id":     u.Id,
		})

	})
	router.Run(":8745")
}
