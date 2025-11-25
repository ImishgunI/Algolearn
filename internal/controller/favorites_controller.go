package controller

import (
	"algolearn/internal/database"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type FavoriteHandler struct {
	repo *database.Database
}

func NewFavoriteHandler(db *database.Database) *FavoriteHandler {
	return &FavoriteHandler{
		repo: db,
	}
}

func (h *FavoriteHandler) CreateFavorite(c *gin.Context) {
	var req struct {
		LessonTitle string `json:"lesson_title"`
		Email       string `json:"email"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body",
		})
		return
	}
	err := h.repo.AddFavoriteLesson(c.Request.Context(), req.Email, req.LessonTitle)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *FavoriteHandler) GetFavorites(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email",
		})
		return
	}
	list, err := h.repo.GetFavoritesByEmail(c.Request.Context(), email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"favorites": list,
	})
}

func (h *FavoriteHandler) DeleteFavorite(c *gin.Context) {
	id := c.Param("lessonId")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid query parametr",
		})
		return
	}
	lessonId, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "Failed to convert string to int",
		})
		return
	}
	err = h.repo.DeleteFavoriteByID(c.Request.Context(), lessonId)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
