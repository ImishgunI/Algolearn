package db

import (
	"Algolearn/internal/users"
	"context"
)

type UserCreator interface {
	CreateUser(ctx context.Context, user *users.User) error
}
