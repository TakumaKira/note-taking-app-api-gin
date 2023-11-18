package main

import (
	"github.com/TakumaKira/note-taking-app-api-gin/internal/transport/rest"
	"github.com/gin-gonic/gin"
)

func setupRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/notes", rest.GetNotes)
	r.GET("/notes/:id", rest.GetNoteByID)
	r.POST("/notes", rest.PostNote)
	r.PUT("/notes/:id", rest.PutNote)
	r.DELETE("/notes/:id", rest.DeleteNote)
	return r
}

func main() {
	r := setupRouter()
	r.Run(":8080")
}
