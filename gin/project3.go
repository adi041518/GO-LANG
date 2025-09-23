package main

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

type Studen struct {
	Name  string `json:"name"`
	Id    int    `json:"id"`
	Marks string `json:"marks"`
	Dept  string `json:"dept"`
}

var nextid = 1
var mp = make(map[int]Studen)

func main() {
	router := gin.Default()
	routers := router.Group("/home")
	{
		// POST /home/
		routers.POST("/", func(ctx *gin.Context) {
			var s Studen

			if err := ctx.BindJSON(&s); err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}

			s.Id = nextid
			nextid++
			mp[s.Id] = s

			ctx.JSON(200, s)
		})

		// GET /home/:id
		routers.GET("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			pid, err := strconv.Atoi(id)
			if err != nil {
				ctx.String(400, "Invalid ID")
				return
			}

			check, exists := mp[pid]
			if !exists {
				ctx.String(404, "Student not found")
				return
			}

			ctx.JSON(200, check)
		})

		// PUT /home/:id
		routers.PUT("/:id", func(ctx *gin.Context) {
			id := ctx.Param("id")
			pid, err := strconv.Atoi(id)
			if err != nil {
				ctx.String(400, "Invalid ID")
				return
			}

			_, exists := mp[pid]
			if !exists {
				ctx.String(404, "Student not found")
				return
			}

			var s Studen
			if err := ctx.BindJSON(&s); err != nil {
				ctx.JSON(400, gin.H{"error": err.Error()})
				return
			}

			// Preserve the original ID
			s.Id = pid

			mp[pid] = s
			ctx.JSON(200, s)
		})
	}
	router.Run()
}
