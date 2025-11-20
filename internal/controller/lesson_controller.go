package controller

import (
	"algolearn/internal/database"
	"algolearn/internal/models"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type LessonFactory interface {
	GetLessons(c *gin.Context)
}

type LessonHandler struct {
	db *database.Database
}

func NewLessonHandler(db *database.Database) *LessonHandler {
	return &LessonHandler{
		db: db,
	}
}

func (h *LessonHandler) GetLessons(c *gin.Context) {
	ls := []models.Lesson{}
	ls, err := h.db.GetLessons(c.Request.Context())
	if err != nil {
		log.Printf("%v", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get lessons",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"lessons": ls,
	})
}
