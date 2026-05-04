package sessions

import (
	"Algolearn/internal/infrastructure/db/sql"
	"context"
	"fmt"
	"time"
)

type Session struct {
	p *sql.Postgres
}

func NewSession(db *sql.Postgres) *Session {
	return &Session{
		p: db,
	}
}

func (s *Session) CreateSession(ctx context.Context, userID int, token string, expiresAt time.Time) error {
	const query = `
		INSERT INTO sessions (user_id, refresh_token, expires_at)
		VALUES ($1, $2, $3)
	`

	_, err := s.p.Pool.Exec(ctx, query, userID, token, expiresAt)
	if err != nil {
		return fmt.Errorf("create session: %w", err)
	}

	return nil
}

func (s *Session) GetSession(ctx context.Context, token string) (int, error) {
	const query = `
		SELECT user_id
		FROM sessions
		WHERE refresh_token = $1 AND expires_at > now()
	`

	var userID int

	err := s.p.Pool.QueryRow(ctx, query, token).Scan(&userID)
	if err != nil {
		return 0, fmt.Errorf("get session: %w", err)
	}

	return userID, nil
}

func (s *Session) DeleteSession(ctx context.Context, token string) error {
	const query = `
		DELETE FROM sessions WHERE refresh_token = $1
	`

	_, err := s.p.Pool.Exec(ctx, query, token)
	if err != nil {
		return fmt.Errorf("delete session: %w", err)
	}

	return nil
}
