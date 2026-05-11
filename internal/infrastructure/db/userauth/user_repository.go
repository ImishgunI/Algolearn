package userauth

import (
	"Algolearn/internal/infrastructure/db/sql"
	"Algolearn/internal/users"
	"context"
)

type Repository struct {
	p *sql.Postgres
}

func NewRepo(db *sql.Postgres) *Repository {
	return &Repository{
		p: db,
	}
}

func (r *Repository) CreateUser(ctx context.Context, user_data *users.User) error {
	const query = `INSERT INTO users (user_name, user_surname, email, password_hash, user_role)
		VALUES ($1, $2, $3, $4, $5)`

	_, err := r.p.Pool.Exec(ctx, query,
		user_data.Name, user_data.Surname, user_data.Email, user_data.PasswordHash, user_data.Role)

	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetByEmail(ctx context.Context, email string) (*users.User, error) {
	const query = `SELECT id, user_name, user_surname, email, password_hash, user_role
					FROM users
					WHERE email = $1`
	var user_info users.User
	err := r.p.Pool.QueryRow(ctx, query, email).Scan(&user_info.ID, &user_info.Name, &user_info.Surname, &user_info.Email, &user_info.PasswordHash, &user_info.Role)
	if err != nil {
		return nil, err
	}
	return &user_info, nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*users.User, error) {
	const query = `SELECT id, user_name, user_surname, email, password_hash, user_role
					FROM users
					WHERE id = $1`
	var user_info users.User
	err := r.p.Pool.QueryRow(ctx, query, id).Scan(&user_info.ID, &user_info.Name, &user_info.Surname, &user_info.Email, &user_info.PasswordHash, &user_info.Role)
	if err != nil {
		return nil, err
	}
	return &user_info, nil
}

func (r *Repository) UpdateProfile(ctx context.Context, userID int, name, surname, email string) error {
	_, err := r.p.Pool.Exec(ctx,
		`UPDATE users
		SET user_name = $1, user_surname = $2, email = $3
		WHERE user_id = $4`, name, surname, email, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) UpdatePassword(ctx context.Context, userID int, hash string) error {
	_, err := r.p.Pool.Exec(ctx,
		`UPDATE users
		SET password_hash = $1
		WHERE user_id = $2`, hash, userID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetAll(ctx context.Context) ([]users.User, error) {
	rows, err := r.p.Pool.Query(ctx, `SELECT id, user_name, user_surname, email, password_hash, user_role FROM users`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var us []users.User
	for rows.Next() {
		var u users.User
		if err := rows.Scan(&u.ID, &u.Name, &u.Surname, &u.Email, &u.PasswordHash, &u.Role); err != nil {
			return nil, err
		}
		us = append(us, u)
	}
	return us, nil
}
func (r *Repository) UpdateRole(ctx context.Context, userID int, role string) error {
	_, err := r.p.Pool.Exec(ctx,
		`UPDATE users
		SET role = $1
		WHERE user_id = $2`, role, userID)
	if err != nil {
		return err
	}
	return nil
}
