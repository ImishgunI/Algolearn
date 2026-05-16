package http

import (
	"Algolearn/internal/execution/custom"
	"log/slog"

	"github.com/gofiber/fiber/v3"
)

type CustomExecutionHandler struct{}

func NewCustomExecutionHandler() *CustomExecutionHandler {
	return &CustomExecutionHandler{}
}

type customExecuteRequest struct {
	Code  string `json:"code"`
	Input []int  `json:"input"`
}

func (h *CustomExecutionHandler) Execute(c fiber.Ctx) error {
	var req customExecuteRequest
	if err := c.Bind().Body(&req); err != nil {
		slog.Error("Не удалось сбиндить тело запроса со структурой customExecuteRequest", "err", err)
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid request"})
	}

	if req.Code == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "code is required"})
	}

	steps, err := custom.RunUserCode(req.Code, req.Input)
	if err != nil {
		slog.Error("Ошибка запуска пользовательского кода", "err", err)
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}

	return c.JSON(fiber.Map{"steps": steps})
}
