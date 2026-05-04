package userauth

import (
	"Algolearn/internal/infrastructure/db/sql"
	"Algolearn/internal/users"
	"context"
)

type Registration struct {
	p *sql.Postgres
}

func NewReg(db *sql.Postgres) *Registration {
	return &Registration{
		p: db,
	}
}

func (p *Registration) CreateUser(ctx context.Context, user_data *users.User) error {
	const query = `INSERT INTO users (user_name, user_surname, email, password_hash, role)
		VALUES ($1, $2, $3, $4, $5)`

	_, err := p.p.Pool.Exec(ctx, query,
		user_data.Name, user_data.Surname, user_data.Email, user_data.PasswordHash, user_data.Role)

	if err != nil {
		return err
	}
	return nil
}
