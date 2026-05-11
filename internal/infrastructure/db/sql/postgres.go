package sql

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Postgres struct {
	Pool *pgxpool.Pool
}

func NewPool(ctx context.Context, connectionString string) (*Postgres, error) {
	conn, err := pgxpool.New(ctx, connectionString)
	if err != nil {
		return nil, err
	}
	return &Postgres{
		Pool: conn,
	}, nil
}

func (p *Postgres) Close() {
	p.Pool.Close()
}
