package controller

import (
	"algolearn/internal/database"
	"algolearn/internal/models"
	"algolearn/internal/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

type Handler struct {
	Db *database.Database
}

type RegisterRequest struct {
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"passwordHash"`
}

func NewUserHandler(db *database.Database) *Handler {
	return &Handler{Db: db}
}

func (h *Handler) RegisterController(c *gin.Context) {
	user := models.NewUser()
	req := &RegisterRequest{}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "bad request",
			"message": err.Error(),
		})
		return
	}
	hashed, err := services.HashPassword(req.Password)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "StatusInternalServerError",
			"message": "Failed to hash password",
		})
		return
	}
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.PasswordHash = hashed
	user.Role = services.CheckRoleForNewUsers(req.Email)
	err = h.Db.Register(c.Request.Context(), user)
	if err != nil {
		log.Println("Failed to register user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "StatusInternalServerError",
			"message": "Failed to register user",
		})
		return
	}
	log.Println("User Register Success")
	c.Status(http.StatusCreated)
}

func (h *Handler) LoginController(c *gin.Context) {
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	user := models.NewUser()
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"code":    "bad request",
			"message": err.Error(),
		})
		return
	}
	user.Email = req.Email
	user.PasswordHash = req.Password
	user, err := h.Db.Login(c.Request.Context(), user)
	if err != nil {
		log.Println("Failed to login user:", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"code":    "StatusInternalServerError",
			"message": err.Error(),
		})
	}
	c.JSON(http.StatusOK, gin.H{
		"first_name": user.FirstName,
		"last_name":  user.LastName,
		"email":      user.Email,
		"role":       user.Role,
	})
	log.Println("User Login Success")
}
