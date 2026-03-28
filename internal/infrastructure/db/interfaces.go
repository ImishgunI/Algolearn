package db

import (
	"Algolearn/internal/transport"
	"context"
)

type UserCreator interface {
	CreateUser(ctx context.Context, user_data *transport.UserRegistrationInfo) error
}
