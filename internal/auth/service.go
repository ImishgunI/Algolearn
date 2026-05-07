package auth

import (
	"Algolearn/internal/infrastructure/db"
	"Algolearn/internal/users"
	"context"
	"crypto/rand"
	"encoding/base64"
	"fmt"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type RegisterInput struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type LoginInput struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthService struct {
	repo     db.UserRepository
	sessions db.SessionRepository
}

func NewService(repo db.UserRepository, session db.SessionRepository) *AuthService {
	return &AuthService{
		repo:     repo,
		sessions: session,
	}
}

func (s *AuthService) Register(ctx context.Context, input RegisterInput) error {
	if input.Email == "" || input.Password == "" {
		return fmt.Errorf("email and password are required")
	}

	if len(input.Password) < 6 {
		return fmt.Errorf("password must be at least 6 characters")
	}

	hash, err := users.HashPassword(input.Password)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	user := &users.User{
		Name:         input.Name,
		Surname:      input.Surname,
		Email:        input.Email,
		PasswordHash: hash,
		Role:         "user",
	}

	if err := s.repo.CreateUser(ctx, user); err != nil {
		return fmt.Errorf("create user: %w", err)
	}

	return nil
}

func (s *AuthService) SignIn(ctx context.Context, input LoginInput) (string, string, error) {
	if input.Email == "" || input.Password == "" {
		return "", "", fmt.Errorf("invalid input")
	}

	user, err := s.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return "", "", fmt.Errorf("user not found")
	}

	if !users.CheckPasswordHash(input.Password, user.PasswordHash) {
		return "", "", fmt.Errorf("invalid credentials")
	}

	access, err := s.createAccessToken(user)
	if err != nil {
		return "", "", err
	}

	refresh, err := s.createRefreshToken()
	if err != nil {
		return "", "", err
	}

	err = s.sessions.CreateSession(ctx, user.ID, refresh, time.Now().Add(7*24*time.Hour))
	if err != nil {
		return "", "", err
	}

	return access, refresh, nil
}

func (s *AuthService) createAccessToken(user *users.User) (string, error) {
	claims := jwt.MapClaims{
		"user_id": user.ID,
		"role":    user.Role,
		"exp":     time.Now().Add(time.Hour * 1).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	return token.SignedString([]byte("secret"))
}

func (s *AuthService) createRefreshToken() (string, error) {
	b := make([]byte, 32)
	_, err := rand.Read(b)
	if err != nil {
		return "", err
	}
	return base64.StdEncoding.EncodeToString(b), nil
}

func (s *AuthService) Refresh(ctx context.Context, refreshToken string) (string, error) {
	userID, err := s.sessions.GetSession(ctx, refreshToken)
	if err != nil {
		return "", fmt.Errorf("invalid session")
	}

	user, err := s.repo.GetByID(ctx, userID)
	if err != nil {
		return "", err
	}

	return s.createAccessToken(user)
}
