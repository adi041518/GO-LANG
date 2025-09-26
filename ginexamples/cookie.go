package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		Cookie, err := ctx.Cookie("usercookie")
		if err != nil {
			Cookie = "NOTSET"
			ctx.SetCookie("usercookie", "ADITYA", 3600, "/", "localhost", false, true)
		}
		ctx.String(200, Cookie)
	})
	router.Run(":8989")
}
