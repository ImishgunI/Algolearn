package http

import (
	"Algolearn/internal/auth"
	"Algolearn/internal/infrastructure/db"
	"Algolearn/internal/transport"

	"github.com/gofiber/fiber/v3"
)

type Registration struct {
	uc *db.UserCreator
}

func (r *Registration) SignUp(c fiber.Ctx) error {
	var userInfo transport.UserRegistrationInfo
	if err := c.Bind().Body(&userInfo); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}
	auth.CreateToken()
	// hash password
	// go to db
	// error handling
	return c.SendStatus(fiber.StatusCreated)
}
