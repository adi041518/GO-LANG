package main

import (
	"errors"

	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		ctx.Next()
		if len(ctx.Errors) > 0 {
			err := ctx.Errors.Last().Err
			ctx.JSON(200, gin.H{
				"success:": "false",
				"error":    err.Error(),
			})
		}
	}
}

func main() {
	router := gin.New()
	router.Use(ErrorHandler())
	router.GET("/", func(ctx *gin.Context) {
		con := true
		if con {
			ctx.Error(errors.New("this is an Error!!!"))
			return
		}
		ctx.JSON(200, gin.H{
			"result": "succsess",
		})
	})
	router.Run()
}
