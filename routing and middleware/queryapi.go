package main

import (
	"log"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

func RequestHeader() gin.HandlerFunc {
	return func(c *gin.Context) {
		start := time.Now()
		c.Next()
		duration := time.Since(start)
		log.Printf("Request %s %s took %v", c.Request.Method, c.Request.URL.Path, duration)
	}
}

func main() {
	router := gin.Default()
	accouts := gin.Accounts{
		"Aditya": "123456",
	}
	router.Use(RequestHeader())
	routers := router.Group("/hello", gin.BasicAuth(accouts))
	{
		routers.GET("/", func(ctx *gin.Context) {
			a := ctx.Query("a")
			b := ctx.Query("b")
			aint, err1 := strconv.Atoi(a)
			if err1 != nil {
				ctx.String(404, "Not a int")
				return
			}
			bint, err2 := strconv.Atoi(b)
			if err2 != nil {
				ctx.String(404, "Not a int")
				return
			}
			ctx.JSON(200, gin.H{
				"a":   aint,
				"b":   bint,
				"sum": aint + bint,
			})
		})
	}
	router.Run(":8520")
}
