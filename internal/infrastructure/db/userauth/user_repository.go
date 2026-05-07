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
