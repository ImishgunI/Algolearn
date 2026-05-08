package http

import (
	"Algolearn/internal/auth"
	"log/slog"

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
		slog.Error("Не удалось спарсить данные в input: ", "err", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	access, refresh, err := a.service.SignIn(c.Context(), input)
	if err != nil {
		slog.Error("Ошибка SignIn: ", "err", err.Error())
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.JSON(fiber.Map{
		"access_token":  access,
		"refresh_token": refresh,
	})
}

func (a *Authorization) Refresh(c fiber.Ctx) error {
	var body struct {
		RefreshToken string `json:"refresh_token"`
	}

	if err := c.Bind().Body(&body); err != nil {
		slog.Error("Не удалось спарсить данные в body: ", "err", err)
		return c.SendStatus(fiber.StatusBadRequest)
	}

	access, err := a.service.Refresh(c.Context(), body.RefreshToken)
	if err != nil {
		slog.Error("Ошибка создания токена: ", "err", err)
		return c.SendStatus(fiber.StatusUnauthorized)
	}

	return c.JSON(fiber.Map{
		"access_token": access,
	})
}

func (h *Authorization) Me(c fiber.Ctx) error {
	userID := c.Locals("user_id")
	user, err := h.service.GetByID(c.Context(), int(userID.(int)))
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}
	return c.JSON(fiber.Map{
		"user_id":      user.ID,
		"user_name":    user.Name,
		"user_surname": user.Surname,
		"email":        user.Email,
		"role":         user.Role,
	})
}
