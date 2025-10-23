package models

import (
	"algolearn/internal/database"
	"algolearn/internal/services"
	"context"
	"fmt"
)

type User struct {
	Name         string
	LastName     string
	Email        string
	PasswordHash string
	Role         string
}

func NewUser() *User {
	return &User{}
}

func (u *User) Register(context context.Context, db *database.Database) error {
	var existID uint64
	err := db.Db.QueryRow(context, "SELECT id FROM users WHERE email=$1", u.Email).Scan(&existID)
	if err != nil && existID > 0 {
		return fmt.Errorf("user already exists %d", existID)
	}
	_, err = db.Db.Exec(context, "INSERT INTO users (first_name, last_name, email, password_hash, role) VALUES ($1, $2, $3, $4, $5)", u.Name, u.LastName, u.Email, u.PasswordHash, u.Role)
	if err != nil {
		return err
	}
	return nil
}

func (u *User) Login(context context.Context, db *database.Database) (bool, error) {
	var storedHash string
	err := db.Db.QueryRow(context, "SELECT password_hash FROM users WHERE email=$1", u.Email).Scan(&storedHash)
	if err != nil {
		return false, fmt.Errorf("user not found")
	}
	ok, err := services.CheckPasswordHash(u.PasswordHash, storedHash)
	if !ok {
		return false, fmt.Errorf("%v", err)
	}
	return true, nil
}
