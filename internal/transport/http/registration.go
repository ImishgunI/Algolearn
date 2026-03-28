package http

import (
	"Algolearn/internal/auth"

	"github.com/gofiber/fiber/v3"
)

type Registration struct{}

type UserRegistrationInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

func (r *Registration) SignUp(c fiber.Ctx) error {
	var userInfo UserRegistrationInfo
	if err := c.Bind().Body(&userInfo); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	auth.CreateToken()
	// hash password
	// go to db
	// error handling
	return c.SendStatus(fiber.StatusCreated)
}
