package middleware

import "github.com/gofiber/fiber/v3"

func RequireRole(role string) fiber.Handler {
	return func(c fiber.Ctx) error {
		userRole := c.Locals("role")

		if userRole == nil || userRole != role {
			return c.SendStatus(403)
		}

		return c.Next()
	}
}
