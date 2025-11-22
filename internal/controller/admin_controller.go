package controller

import (
	"algolearn/internal/database"
	"algolearn/internal/models"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type AdminHandler struct {
	db *database.Database
}

func NewAdminHandler(db *database.Database) *AdminHandler {
	return &AdminHandler{
		db: db,
	}
}

func (h *AdminHandler) GetAllUsers(c *gin.Context) {
	users, err := h.db.GetAllUsers(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get all users from database",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"users": users,
	})
}

func (h *AdminHandler) GetStatistics(c *gin.Context) {
	adminStatistics, err := h.db.GetStatistics(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to get statistics",
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"usersCount":    adminStatistics.UsersCount,
		"lessonsCount":  adminStatistics.LessonsCount,
		"commentsCount": adminStatistics.CommentsCount,
		"tasksCount":    adminStatistics.TasksCount,
	})
}

func (h *AdminHandler) GetLessonsForAdmin(c *gin.Context) {
	adminLessons, err := h.db.GetLessonsForAdmin(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"lessons": adminLessons,
	})
}

func (h *AdminHandler) GetLessonToEdit(c *gin.Context) {
	id := c.Param("lessonId")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid param",
		})
		return
	}
	lessonID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to convert string to int",
		})
		return
	}
	al, err := h.db.GetLessonToEdit(c.Request.Context(), lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         al.ID,
		"title":      al.Title,
		"category":   al.Category,
		"difficulty": al.Difficulty,
		"content":    al.Content,
	})
}

func (h *AdminHandler) DeleteLesson(c *gin.Context) {
	id := c.Param("lessonId")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid param",
		})
		return
	}
	lessonID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to convert string to int",
		})
		return
	}
	err = h.db.DeleteLesson(c.Request.Context(), lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func (h *AdminHandler) AddLesson(c *gin.Context) {
	req := models.LessonCreator{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body",
		})
		return
	}

	err := h.db.AddLesson(c.Request.Context(), &req)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *AdminHandler) UpdateLesson(c *gin.Context) {
	id := c.Param("lessonId")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid param",
		})
		return
	}
	lessonID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to convert string to int",
		})
		return
	}
	req := models.LessonCreator{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body",
		})
		return
	}
	err = h.db.UpdateLesson(c.Request.Context(), &req, lessonID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func (h *AdminHandler) AddUser(c *gin.Context) {
	req := models.UserCreator{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body",
		})
		return
	}
	err := h.db.AddUser(c.Request.Context(), &req)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Status(http.StatusCreated)
}

func (h *AdminHandler) UpdateUser(c *gin.Context) {
	id := c.Param("userId")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid param",
		})
		return
	}
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to convert string to int",
		})
		return
	}
	req := models.UserCreator{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid body",
		})
		return
	}
	err = h.db.UpdateUser(c.Request.Context(), &req, userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func (h *AdminHandler) GetUserById(c *gin.Context) {
	id := c.Param("userId")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid param",
		})
		return
	}
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to convert string to int",
		})
		return
	}
	ar, err := h.db.GetUserById(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"id":         ar.ID,
		"first_name": ar.FirstName,
		"last_name":  ar.LastName,
		"email":      ar.Email,
		"role":       ar.Role,
	})
}

func (h *AdminHandler) DeleteUser(c *gin.Context) {
	id := c.Param("userId")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid param",
		})
		return
	}
	userID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to convert string to int",
		})
		return
	}
	err = h.db.DeleteUser(c.Request.Context(), userID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func (h *AdminHandler) GetAllComments(c *gin.Context) {
	comments, err := h.db.GetAllComments(c.Request.Context())
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"comments": comments,
	})
}

func (h *AdminHandler) DeleteComment(c *gin.Context) {
	id := c.Param("commentId")
	if id == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": "invalid param",
		})
		return
	}
	commentID, err := strconv.Atoi(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": "Failed to convert string to int",
		})
		return
	}
	err = h.db.DeleteComment(c.Request.Context(), commentID)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"message": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}
