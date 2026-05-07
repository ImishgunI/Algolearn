package transport

import (
	"Algolearn/internal/middleware"
	"Algolearn/internal/transport/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func Routes(reg *http.Registration, auth *http.Authorization, exec *http.ExecutionHandler) *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	app.Post("/register", reg.SignUp)
	app.Post("/login", auth.SignIn)
	app.Post("/refresh", auth.Refresh)

	app.Get("/me", middleware.JWTMiddleware("secret"), auth.Me)

	app.Post("/execute", exec.Execute)
	app.Get("/execution/:id", exec.Get)
	return app
}
