package controller

import (
	"algolearn/internal/database"
	"algolearn/internal/models"
	"log"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
)

type CommentHandler struct {
	db *database.Database
}

func NewCommentHandler(db *database.Database) *CommentHandler {
	return &CommentHandler{db: db}
}

func (h *CommentHandler) AddComment(c *gin.Context) {
	var req models.CommentRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "bad request",
			"message": "Failed to parse. Invalid Body",
		})
		return
	}
	err := h.db.AddComment(c.Request.Context(), &req)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to add comment",
		})
		return
	}
}

func (h *CommentHandler) GetComments(c *gin.Context) {
	title := c.Query("lesson_title")
	code := strings.Compare(title, "")
	if code != 0 {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid query",
		})
		return
	}
	var comments []models.Comment
	comments, err := h.db.GetComments(c.Request.Context(), title)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get comments",
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})
}
