package database

import (
	"algolearn/internal/config"
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
)

type Database struct {
	Db *pgx.Conn
}

func Connect(ctx context.Context, config config.DBConfig) *Database {
	postgresUrl := fmt.Sprintf("postgresql://%s:%s@%s:%s/%s?search_path=public", config.User, config.Password, config.Host, config.Port, config.DBName)
	conn, err := pgx.Connect(ctx, postgresUrl)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Successfully connected to database")
	return &Database{
		Db: conn,
	}
}

func Close(ctx context.Context, conn *Database) {
	err := conn.Db.Close(ctx)
	if err != nil {
		log.Println(err)
	}
	log.Println("Successfully closed database")
}
