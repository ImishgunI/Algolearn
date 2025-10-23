package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

type DConfig struct {
	Port string
}

type DBConfig struct {
	Host     string
	Port     string
	User     string
	Password string
	DBName   string
}

func Config() *DConfig {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error loading .env file")
	}
	return &DConfig{Port: os.Getenv("PORT")}
}

func DbConfig() *DBConfig {
	if err := godotenv.Load(".env"); err != nil {
		log.Fatal("Error: db .env file not found")
	}
	return &DBConfig{
		Host:     os.Getenv("DBHOST"),
		Port:     os.Getenv("DBPORT"),
		User:     os.Getenv("DBUSER"),
		Password: os.Getenv("DBPASSWORD"),
		DBName:   os.Getenv("DBNAME"),
	}
}
