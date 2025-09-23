// main.go
package main

import (
	"exam.com/API-project/controllers"
	"exam.com/API-project/services"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	studentService := services.NewStudentService()
	studentController := controllers.NewStudentController(studentService)

	r.POST("/students", studentController.CreateStudentController)
	r.GET("/students/:id", studentController.GetStudentByID)
	r.GET("/students", studentController.GetAllStudents)
	r.PUT("/students/:id", studentController.UpdateStudent)
	r.GET("/students/:id/:dept", studentController.UpdateByDept)
	r.Run(":8000")
}
