package mockdata

import (
	"github.com/TakumaKira/note-taking-app-api-gin/internal/models"
)

var Notes = []models.Note{
	{ID: "1", Title: "Note 1", Content: "This is a note #1.", CreatedAt: "2023-11-17T10:14:08.861Z"},
	{ID: "2", Title: "Note 2", Content: "This is a note #2.", CreatedAt: "2023-11-17T10:14:08.861Z"},
	{ID: "3", Title: "Note 3", Content: "This is a note #3.", CreatedAt: "2023-11-17T10:14:08.861Z"},
}
