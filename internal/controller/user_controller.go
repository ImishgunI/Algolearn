package controller

import (
	"algolearn/internal/database"
	"algolearn/internal/models"
	"algolearn/internal/services"
	"context"
	"encoding/json"
	"log"
	"net/http"
)

type Handler struct {
	Db *database.Database
}

type RegisterRequest struct {
	Name     string `json:"name"`
	LastName string `json:"lastName"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func NewHandler(db *database.Database) *Handler {
	return &Handler{Db: db}
}

func (h *Handler) RegisterController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Method Not Allowed")
		return
	}
	defer r.Body.Close()
	user := models.NewUser()
	req := &RegisterRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Status Bad Request")
		return
	}
	hashed, err := services.HashPassword(req.Password)
	if err != nil {
		log.Println(err)
	}
	user.Name = req.Name
	user.LastName = req.LastName
	user.Email = req.Email
	user.PasswordHash = hashed
	user.Role = "student"
	err = user.Register(context.Background(), h.Db)
	if err != nil {
		log.Println("Failed to register user:", err)
		http.Error(w, "Registration failed", http.StatusUnauthorized)
		return
	}
	log.Println("User Register Success")
	w.WriteHeader(http.StatusCreated)
}

func (h *Handler) LoginController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Method Not Allowed")
		return
	}
	defer r.Body.Close()
	var req struct {
		Email    string `json:"email"`
		Password string `json:"password"`
	}
	user := models.NewUser()
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Status Bad Request")
		return
	}
	user.Email = req.Email
	user.PasswordHash = req.Password
	ok, err := user.Login(context.Background(), h.Db)
	if err != nil || !ok {
		log.Println("Failed to login user:", err)
		http.Error(w, "Login failed", http.StatusUnauthorized)
		return
	}
	log.Println("User Login Success")
	w.WriteHeader(http.StatusOK)
}
