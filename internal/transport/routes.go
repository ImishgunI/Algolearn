package transport

import (
	"Algolearn/internal/middleware"
	"Algolearn/internal/transport/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
)

func Routes(reg *http.Registration, auth *http.Authorization, exec *http.ExecutionHandler) *fiber.App {
	app := fiber.New()
	app.Use(cors.New())
	api := app.Group("/api")
	api.Post("/register", reg.SignUp)
	api.Post("/login", auth.SignIn)
	api.Post("/refresh", auth.Refresh)

	api.Get("/me", middleware.JWTMiddleware("secret"), auth.Me)

	api.Post("/execute", exec.Execute)
	api.Get("/execution/:id", exec.Get)
	return app
}
