package http

import (
	"Algolearn/internal/auth"

	"github.com/gofiber/fiber/v3"
)

type Authorization struct {
	service *auth.AuthService
}

func NewAuthorization(service *auth.AuthService) *Authorization {
	return &Authorization{
		service: service,
	}
}

func (a *Authorization) SignIn(c fiber.Ctx) error {
	var input auth.LoginInput

	if err := c.Bind().Body(&input); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	token, err := a.service.SignIn(c.Context(), input)
	if err != nil {
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.JSON(fiber.Map{
		"access_token": token,
	})
}
