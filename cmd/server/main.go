package main

import (
	"algolearn/internal/config"
	"algolearn/internal/database"
	"algolearn/internal/routes"
	"context"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rs/cors"
)

func main() {
	r := gin.Default()
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
	routes.SetRoutes(r, db)
	log.Println("Connecting to localhost")
	handler := c.Handler(r)
	if err := http.ListenAndServe(":"+config.GetString("PORT"), handler); err != nil {
		log.Fatal(err)
	}
	defer database.Close(ctx, db)
}
