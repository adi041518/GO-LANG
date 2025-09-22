package main

import (
	"your_project/controllers"
	"your_project/services"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	noteService := services.NewNoteService()
	notesController := controllers.NewNotesController(noteService)
	notesController.InitNotesControllerRoutes(router)

	router.Run(":8080")
}
