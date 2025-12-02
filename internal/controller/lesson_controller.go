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
	SetDone(c *gin.Context)
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

func (h *LessonHandler) SetDone(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email",
		})
		return
	}
	var req struct {
		Lesson_title string `json:"lesson_title"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body",
		})
		return
	}
	err := h.db.SetLessonComplete(c.Request.Context(), email, req.Lesson_title)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
