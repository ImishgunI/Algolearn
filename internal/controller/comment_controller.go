package controller

import (
	"algolearn/internal/database"
	"algolearn/internal/models"
	"log"
	"net/http"

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
	comment := &models.Comment{}
	comment, err := h.db.AddComment(c.Request.Context(), &req)
	if err != nil {
		log.Println("ID not found in users or lessons")
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to add comment",
		})
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"content":    comment.Content,
		"created_at": comment.CreatedAt,
	})
}
