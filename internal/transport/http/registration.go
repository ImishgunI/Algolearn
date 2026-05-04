package http

import (
	"Algolearn/internal/auth"

	"github.com/gofiber/fiber/v3"
)

type Registration struct {
	service *auth.AuthService
}

func NewRegistration(service *auth.AuthService) *Registration {
	return &Registration{
		service: service,
	}
}

func (r *Registration) SignUp(c fiber.Ctx) error {
	var userInfo auth.RegisterInput

	if err := c.Bind().Body(&userInfo); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	err := r.service.Register(c.Context(), userInfo)
	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}

	return c.SendStatus(fiber.StatusCreated)
}
