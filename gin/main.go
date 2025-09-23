package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Student struct {
	Name  string `json:"name"`
	Id    string `json: "id"`
	Marks string `json:"marks"`
	Dept  string `json:"dept"`
}

func main() {
	var dummy Student
	students := make([]Student, 1)
	for i := 0; i < 1; i++ {
		fmt.Println("Enter the Students name,id,marks,dept:::")
		fmt.Scan(&dummy.Name)
		fmt.Scan(&dummy.Id)
		fmt.Scan(&dummy.Marks)
		fmt.Scan(&dummy.Dept)
		students = append(students, dummy)
	}
	fmt.Println(dummy.Name)
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"students": students,
		})
	})
	router.Run(":8000")

}
