package middleware

import "github.com/gofiber/fiber/v3"

func RequireRole(roles ...string) fiber.Handler {
	return func(c fiber.Ctx) error {
		role, ok := c.Locals("role").(string)
		if !ok {
			return c.SendStatus(fiber.StatusUnauthorized)
		}
		for _, r := range roles {
			if role == r {
				return c.Next()
			}
		}
		return c.SendStatus(fiber.StatusForbidden)
	}
}
