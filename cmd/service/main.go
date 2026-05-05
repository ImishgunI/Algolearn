package main

import (
	"context"
	"log"
	"os/signal"
	"syscall"
	"time"

	"Algolearn/internal/auth"
	"Algolearn/internal/execution/manager"
	"Algolearn/internal/execution/storage"
	"Algolearn/internal/execution/worker"
	"Algolearn/internal/infrastructure/cache"
	"Algolearn/internal/infrastructure/db/sessions"
	"Algolearn/internal/infrastructure/db/sql"
	"Algolearn/internal/infrastructure/db/userauth"
	"Algolearn/internal/transport"
	"Algolearn/internal/transport/http"
)

func main() {
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	psql, err := sql.NewPool(ctx, "")
	if err != nil {
		log.Fatalf("db error: %+v", err)
	}
	defer psql.Close()

	repo := userauth.NewRepo(psql)
	session := sessions.NewSession(psql)

	service := auth.NewService(repo, session)

	regHandler := http.NewRegistration(service)

	authHandler := http.NewAuthorization(service)

	rdb := cache.NewRedisClient("redis:6379")
	defer rdb.Close()
	storage := storage.New(rdb)
	execManager := manager.New(storage)
	execHandler := http.NewExecutionHandler(execManager)

	worker := worker.New(storage)

	go worker.Start(ctx)

	app := transport.Routes(regHandler, authHandler, execHandler)

	go func() {
		if err := app.Listen(":8000"); err != nil {
			log.Fatalf("listen error %+v\n", err)
		}
	}()

	<-ctx.Done()
	log.Println("Shutdown signal received")

	shutdownCtx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := app.ShutdownWithContext(shutdownCtx); err != nil {
		log.Printf("shutdown error %+v\n", err)
	}
}
