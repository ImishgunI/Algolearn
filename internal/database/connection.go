package database

import (
	"algolearn/internal/config"
	"context"
	"log"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	db *pgx.Conn
}

func Connect(ctx context.Context) *Database {

	conn, err := pgx.Connect(ctx, config.GetString("POSTGRES_URL"))
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to database")
	return &Database{
		db: conn,
	}
}

func Close(ctx context.Context, conn *Database) {
	err := conn.db.Close(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully closed database")
}
