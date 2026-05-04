package db

import (
	"Algolearn/internal/users"
	"context"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *users.User) error
	GetByEmail(ctx context.Context, email string) (*users.User, error)
}
