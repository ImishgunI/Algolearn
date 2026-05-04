package db

import (
	"Algolearn/internal/users"
	"context"
	"time"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *users.User) error
	GetByEmail(ctx context.Context, email string) (*users.User, error)
	GetByID(ctx context.Context, id int) (*users.User, error)
}

type SessionRepository interface {
	CreateSession(ctx context.Context, userID int, token string, expiresAt time.Time) error
	GetSession(ctx context.Context, token string) (int, error)
	DeleteSession(ctx context.Context, token string) error
}
