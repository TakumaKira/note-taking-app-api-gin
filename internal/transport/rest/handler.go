package rest

import (
	"net/http"

	"github.com/TakumaKira/note-taking-app-api-gin/internal/mockdata"
	"github.com/gin-gonic/gin"
)

func GetNotes(c *gin.Context) {
	c.IndentedJSON(http.StatusOK, mockdata.Notes)
}
