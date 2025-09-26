package main

import (
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"res": "success",
		})
	})
	s := &http.Server{
		Addr:         ":8000",
		Handler:      router,
		ReadTimeout:  2 * time.Second,
		WriteTimeout: 2 * time.Second,
	}
	s.ListenAndServe()
}
