package transport

import (
	"Algolearn/internal/transport/http"

	"github.com/gofiber/fiber/v3"
)

func Routes(reg *http.Registration) *fiber.App {
	app := fiber.New()
	app.Post("/registration", reg.SignUp)
	return app
}
