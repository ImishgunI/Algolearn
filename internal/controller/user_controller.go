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

func (h *Handler) GetDataController(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email",
		})
		return
	}
	u, err := h.Db.GetUserData(c.Request.Context(), &email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"first_name": u.FirstName,
		"last_name":  u.LastName,
		"email":      u.Email,
		"role":       u.Role,
	})
}

func (h *Handler) UpdateUserData(c *gin.Context) {
	var req struct {
		Username string `json:"username"`
		Lastname string `json:"lastname"`
		Email    string `json:"email"`
		Password string `json:"password"`
	}

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid body",
		})
		return
	}
	ud := models.UserUpdate{}
	ud.FirstName = req.Username
	ud.LastName = req.Lastname
	ud.Email = req.Email
	if req.Password != "" {
		psh, err := services.HashPassword(req.Password)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{
				"error": err.Error(),
			})
			return
		}
		ud.PasswordHash = psh
	} else {
		ud.PasswordHash = ""
	}

	err := h.Db.UpdateUserData(c.Request.Context(), &ud)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.Status(http.StatusOK)
}

func (h *Handler) GetProgress(c *gin.Context) {
	email := c.Query("email")
	if email == "" {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "invalid email",
		})
		return
	}
	completed, favorites, err := h.Db.GetUserProgress(c.Request.Context(), &email)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": err.Error(),
		})
		return
	}
	c.JSON(http.StatusOK, gin.H{
		"completed": completed,
		"favorites": favorites,
	})
}
