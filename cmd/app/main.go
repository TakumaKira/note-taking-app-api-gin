package main

import (
	"github.com/TakumaKira/note-taking-app-api-gin/internal/transport/rest"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/notes", rest.GetNotes)

	router.Run("localhost:8080")
}
