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
	FirstName string `json:"firstName"`
	LastName  string `json:"lastName"`
	Email     string `json:"email"`
	Password  string `json:"passwordHash"`
}

func NewHandler(db *database.Database) *Handler {
	return &Handler{Db: db}
}

func CloseBody(r *http.Request) {
	if err := r.Body.Close(); err != nil {
		log.Println(err)
		return
	}
}

func (h *Handler) RegisterController(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		log.Println("Method Not Allowed")
		return
	}
	defer CloseBody(r)
	user := models.NewUser()
	req := &RegisterRequest{}
	if err := json.NewDecoder(r.Body).Decode(req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println("Status Bad Request")
		return
	}
	log.Println(req)
	hashed, err := services.HashPassword(req.Password)
	if err != nil {
		log.Println(err)
	}
	user.FirstName = req.FirstName
	user.LastName = req.LastName
	user.Email = req.Email
	user.PasswordHash = hashed
	user.Role = "student"
	log.Println(user)
	ok, err := h.Db.UniqueEmail(context.Background(), user.Email)
	if err != nil || !ok {
		w.WriteHeader(http.StatusInternalServerError)
		log.Println(err)
		return
	}
	err = h.Db.Register(context.Background(), user)
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
	defer CloseBody(r)
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
	ok, err := h.Db.Login(context.Background(), user)
	if err != nil || !ok {
		log.Println("Failed to login user:", err)
		http.Error(w, "Login failed", http.StatusUnauthorized)
		return
	}
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(user); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		log.Println(err)
		return
	}
	log.Println("User Login Success")
}
