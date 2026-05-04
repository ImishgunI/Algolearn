package auth

import (
	"Algolearn/internal/infrastructure/db"
	"Algolearn/internal/users"
	"context"
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
	Email    string
	Password string
}

type AuthService struct {
	repo db.UserRepository
}

func NewService(repo db.UserRepository) *AuthService {
	return &AuthService{repo: repo}
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

func (s *AuthService) SignIn(ctx context.Context, input LoginInput) (string, error) {
	// 1. валидация
	if input.Email == "" || input.Password == "" {
		return "", fmt.Errorf("email and password required")
	}

	// 2. найти пользователя
	user, err := s.repo.GetByEmail(ctx, input.Email)
	if err != nil {
		return "", fmt.Errorf("get user: %w", err)
	}

	// 3. проверить пароль
	if !users.CheckPasswordHash(input.Password, user.PasswordHash) {
		return "", fmt.Errorf("invalid credentials")
	}

	// 4. создать JWT
	token, err := s.createAccessToken(user)
	if err != nil {
		return "", fmt.Errorf("create token: %w", err)
	}

	return token, nil
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
