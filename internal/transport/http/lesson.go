package http

import (
	"Algolearn/internal/infrastructure/db"

	"github.com/gofiber/fiber/v3"
)

type LessonHandler struct {
	repo db.LessonRepository
}

func NewLessonHandler(repo db.LessonRepository) *LessonHandler {
	return &LessonHandler{repo: repo}
}

func (h *LessonHandler) GetByCourse(c fiber.Ctx) error {
	courseID := fiber.Params[int](c, "CourseID")
	if courseID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid course id"})
	}
	lessons, err := h.repo.GetByCourse(c.Context(), courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(lessons)
}

func (h *LessonHandler) GetByID(c fiber.Ctx) error {
	id := fiber.Params[int](c, "id")
	if id == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid lesson id"})
	}
	lesson, err := h.repo.GetByID(c.Context(), id)
	if err != nil {
		return c.Status(fiber.StatusNotFound).JSON(fiber.Map{"error": "lesson not found"})
	}
	return c.JSON(lesson)
}
