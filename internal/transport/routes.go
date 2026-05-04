package transport

import (
	"Algolearn/internal/middleware"
	"Algolearn/internal/transport/http"

	"github.com/gofiber/fiber/v3"
)

func Routes(reg *http.Registration, auth *http.Authorization, exec *http.ExecutionHandler) *fiber.App {
	app := fiber.New()

	app.Post("/registration", reg.SignUp)
	app.Post("/login", auth.SignIn)
	app.Post("/refresh", auth.Refresh)

	app.Get("/me", middleware.JWTMiddleware("secret"), func(c fiber.Ctx) error {
		userID := c.Locals("user_id")

		return c.JSON(fiber.Map{
			"user_id": userID,
		})
	})

	app.Post("/execute", exec.Execute)
	app.Get("/execution/:id", exec.Get)
	return app
}
