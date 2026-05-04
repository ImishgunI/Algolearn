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

	execID, err := h.manager.Execute(c.Context(), req.Algorithm, req.Data)
	if err != nil {
		return c.SendStatus(fiber.StatusBadRequest)
	}

	return c.JSON(fiber.Map{
		"execution_id": execID,
	})
}

func (h *ExecutionHandler) Get(c fiber.Ctx) error {
	execID := c.Params("id")

	steps, err := h.manager.Get(c.Context(), execID)
	if err != nil {
		return c.SendStatus(fiber.StatusNotFound)
	}

	return c.JSON(fiber.Map{
		"steps": steps,
	})
}
