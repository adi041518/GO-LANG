package main

import (
	"github.com/gin-gonic/gin"
)

type User struct {
	Name string `json:"name"`
	Id   int    `json:"id"`
}

func main() {
	router := gin.Default()

	router.GET("/adi/:id/:name", func(c *gin.Context) {
		id := c.Param("id")
		name := c.Param("name")
		c.JSON(200, gin.H{
			"id":   id,
			"name": name,
		})
	})
	router.POST("/json", func(c *gin.Context) {
		var u User
		c.BindJSON(&u)
		c.JSON(200, gin.H{
			"Name": u.Name,
			"Id":   u.Id,
		})
	})

	router.Run()
}
