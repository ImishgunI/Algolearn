package http

import "github.com/gofiber/fiber/v3"

func Routes(app *fiber.App) {
	app.Get("/health", HealthCheck)
}
