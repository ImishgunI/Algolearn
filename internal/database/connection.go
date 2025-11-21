package database

import (
	"context"
	"log"
	"os"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Database struct {
	db *pgxpool.Pool
}

func Connect(ctx context.Context) *Database {

	pool, err := pgxpool.New(ctx, os.Getenv("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to database")
	return &Database{
		db: pool,
	}
}

func Close(conn *Database) {
	conn.db.Close()
	log.Println("Successfully closed database")
}
