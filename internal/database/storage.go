package database

import (
	"algolearn/internal/models"
	"algolearn/internal/services"
	"context"
	"fmt"
)

func (d *Database) Register(context context.Context, user *models.User) error {
	var existID uint64
	err := d.db.QueryRow(context, "SELECT id FROM users WHERE email=$1", user.Email).Scan(&existID)
	if err != nil && existID > 0 {
		return fmt.Errorf("user already exists %d", existID)
	}
	_, err = d.db.Exec(context,
		"INSERT INTO users (first_name, last_name, email, password_hash, role) VALUES ($1, $2, $3, $4, $5)",
		user.FirstName, user.LastName, user.Email, user.PasswordHash, user.Role)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) Login(context context.Context, user *models.User) (bool, error) {
	var storedHash string
	err := d.db.QueryRow(context, "SELECT password_hash FROM users WHERE email=$1 ORDER BY id DESC LIMIT 1", user.Email).Scan(&storedHash)
	if err != nil {
		return false, fmt.Errorf("user not found")
	}
	ok, err := services.CheckPasswordHash(user.PasswordHash, storedHash)
	if !ok {
		return false, fmt.Errorf("%v", err)
	}
	return true, nil
}

func (d *Database) UniqueEmail(context context.Context, email string) (bool, error) {
	var id uint64
	err := d.db.QueryRow(context, "SELECT id FROM users WHERE email=$1", email).Scan(&id)
	if err != nil && id > 0 {
		return false, err
	}
	return true, nil
}
