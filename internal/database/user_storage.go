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
	GetUserData(ctx context.Context, email *string) (*models.UserCreator, error)
	UpdateUserData(ctx context.Context, userdata *models.UserUpdate) error
	GetUserProgress(ctx context.Context, email *string) (int, int, error)
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

func (d *Database) GetUserData(ctx context.Context, email *string) (*models.UserCreator, error) {
	var u models.UserCreator
	err := d.db.QueryRow(ctx, `SELECT first_name, last_name, email, role FROM users WHERE email=$1`, email).Scan(&u.FirstName, &u.LastName, &u.Email, &u.Role)
	if err != nil {
		return nil, err
	}
	return &u, nil
}

func (d *Database) UpdateUserData(ctx context.Context, userdata *models.UserUpdate) error {
	if userdata.PasswordHash == "" {
		_, err := d.db.Exec(ctx, `
		UPDATE users
		SET first_name = $1, last_name = $2
		WHERE email = $3
		`, userdata.FirstName, userdata.LastName, userdata.Email)
		if err != nil {
			return err
		}
	} else {
		_, err := d.db.Exec(ctx, `
		UPDATE users
		SET first_name = $1, last_name = $2, password_hash = $3
		WHERE email = $4
		`, userdata.FirstName, userdata.LastName, userdata.PasswordHash, userdata.Email)
		if err != nil {
			return err
		}
	}
	return nil
}

func (d *Database) GetUserProgress(ctx context.Context, email *string) (int, int, error) {
	var done_lessons int
	err := d.db.QueryRow(ctx, ` 
		SELECT COUNT(*) as done_lessons from completed_lessons as cl
		JOIN users u ON u.id = cl.user_id
		WHERE u.email = $1
	`, email).Scan(&done_lessons)
	if err != nil {
		return 0, 0, nil
	}
	var favorites_lessons int
	err = d.db.QueryRow(ctx, `
	SELECT COUNT(*) as favorites_lessons FROM favorites as f
	JOIN users u ON u.id = f.user_id
	WHERE u.email = $1
	`, email).Scan(&favorites_lessons)
	if err != nil {
		return 0, 0, nil
	}
	return done_lessons, favorites_lessons, nil
}
