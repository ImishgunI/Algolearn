package main

import (
	"algolearn/internal/config"
	"algolearn/internal/controller"
	"algolearn/internal/database"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/rs/cors"
)

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 20*time.Second)
	db := database.Connect(ctx)
	defer cancel()
	c := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:3000", "http://127.0.0.1:3000"},
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"},
		AllowCredentials: true,
		Debug:            true,
	})
	h := controller.NewHandler(db)
	mux := http.NewServeMux()
	mux.HandleFunc("/register", h.RegisterController)
	mux.HandleFunc("/login", h.LoginController)
	log.Println("Connecting to localhost")
	handler := c.Handler(mux)
	if err := http.ListenAndServe(":"+config.GetString("PORT"), handler); err != nil {
		log.Fatal(err)
	}
	defer database.Close(ctx, db)
}
