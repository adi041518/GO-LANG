package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/hello", func(ctx *gin.Context) {
		cp := ctx.Copy()
		go func() {
			time.Sleep(5 * time.Second)
			fmt.Println("url is:", cp.Request.URL)
		}()
		go func() {
			time.Sleep(2 * time.Second)
			fmt.Println("url irequest is:", cp.Request.URL)
		}()
		ctx.JSON(200, gin.H{
			"message": "received",
		})
	})
	router.Run(":8000")

}
