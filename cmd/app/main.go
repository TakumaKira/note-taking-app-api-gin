package main

import (
	"github.com/TakumaKira/note-taking-app-api-gin/internal/transport/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/notes", rest.GetNotes)
	router.GET("/notes/:id", rest.GetNoteByID)
	router.POST("/notes", rest.PostNote)
	router.PUT("/notes/:id", rest.PutNote)
	router.DELETE("/notes/:id", rest.DeleteNote)

	router.Run("localhost:8080")
}
