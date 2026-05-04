package http

import (
	"Algolearn/internal/execution/manager"

	"github.com/gofiber/fiber/v3"
)

type ExecutionHandler struct {
	manager *manager.Manager
}

func NewExecutionHandler(m *manager.Manager) *ExecutionHandler {
	return &ExecutionHandler{
		manager: m,
	}
}

type ExecuteRequest struct {
	Algorithm string `json:"algorithm"`
	Data      []int  `json:"data"`
}

func (h *ExecutionHandler) Execute(c fiber.Ctx) error {
	var req ExecuteRequest

	if err := c.Bind().Body(&req); err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	steps, err := h.manager.Execute(req.Algorithm, req.Data)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"steps": steps,
	})
}
