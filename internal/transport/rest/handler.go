package rest

import (
	"net/http"

	"github.com/TakumaKira/note-taking-app-api-gin/internal/mockdata"
	"github.com/TakumaKira/note-taking-app-api-gin/internal/models"
	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mockdata.Notes)
}

func GetNoteByID(c *gin.Context) {
	id := c.Param("id")

	for _, n := range mockdata.Notes {
		if n.ID == id {
			c.IndentedJSON(http.StatusOK, n)
			return
		}
	}
	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note not found"})
}

func PostNote(c *gin.Context) {
	var newNote models.Note

	if err := c.BindJSON(&newNote); err != nil {
		return
	}

	mockdata.Notes = append(mockdata.Notes, newNote)
	c.IndentedJSON(http.StatusCreated, newNote)
}

func PutNote(c *gin.Context) {
	id := c.Param("id")
	var updatedNote models.Note

	if err := c.BindJSON(&updatedNote); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "invalid request"})
		return
	}

	for i, n := range mockdata.Notes {
		if n.ID == id {
			mockdata.Notes[i] = updatedNote
			c.IndentedJSON(http.StatusOK, updatedNote)
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note not found"})
}

func DeleteNote(c *gin.Context) {
	id := c.Param("id")

	for i, n := range mockdata.Notes {
		if n.ID == id {
			mockdata.Notes = append(mockdata.Notes[:i], mockdata.Notes[i+1:]...)
			c.IndentedJSON(http.StatusOK, gin.H{"message": "note deleted"})
			return
		}
	}

	c.IndentedJSON(http.StatusNotFound, gin.H{"message": "note not found"})
}
