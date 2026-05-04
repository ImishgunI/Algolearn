package auth

import (
	"Algolearn/internal/infrastructure/db"
	"Algolearn/internal/users"
	"context"
	"fmt"
)

type RegisterInput struct {
	Name     string `json:"name"`
	Surname  string `json:"surname"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

type AuthService struct {
	repo db.UserCreator
}

func NewService(repo db.UserCreator) *AuthService {
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
