package database

import (
	"algolearn/internal/models"
	"algolearn/internal/services"
	"context"
	"errors"
	"fmt"
)

type UserRepository interface {
	Register(context context.Context, user *models.User) error
	Login(context context.Context, user *models.User) (bool, error)
}

func (d *Database) Register(context context.Context, user *models.User) error {
	var existID bool
	_ = d.db.QueryRow(context, "SELECT id FROM users WHERE email=$1", user.Email).Scan(&existID)
	if existID {
		return fmt.Errorf("user already exists %t", existID)
	}
	_, err := d.db.Exec(context,
		"INSERT INTO users (first_name, last_name, email, password_hash, role) VALUES ($1, $2, $3, $4, $5)",
		user.FirstName, user.LastName, user.Email, user.PasswordHash, user.Role)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) Login(context context.Context, user *models.User) (*models.User, error) {
	var storedHash string
	err := d.db.QueryRow(context, `SELECT first_name, last_name, password_hash, role FROM users WHERE email=$1`, user.Email).
		Scan(&user.FirstName, &user.LastName, &storedHash, &user.Role)
	if err != nil {
		return nil, errors.New("Failed to select datas from users")
	}
	ok, err := services.CheckPasswordHash(user.PasswordHash, storedHash)
	if !ok {
		return nil, fmt.Errorf("%v", err)
	}
	return user, nil
}
