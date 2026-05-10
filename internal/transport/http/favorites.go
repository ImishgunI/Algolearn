package http

import (
	"Algolearn/internal/infrastructure/db"

	"github.com/gofiber/fiber/v3"
)

type FavoriteHandler struct {
	repo db.FavoriteRepository
}

func NewFavoriteHandler(repo db.FavoriteRepository) *FavoriteHandler {
	return &FavoriteHandler{
		repo: repo,
	}
}

func (h *FavoriteHandler) Toggle(c fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	lessonID := fiber.Params[int](c, "lessonID")
	if lessonID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid lesson id"})
	}
	isFav, _ := h.repo.IsFavorited(c.Context(), userID, lessonID)
	if isFav {
		h.repo.Remove(c.Context(), userID, lessonID)
		return c.JSON(fiber.Map{"favorited": false})
	} else {
		h.repo.Add(c.Context(), userID, lessonID)
		return c.JSON(fiber.Map{"favorited": true})
	}
}

func (h *FavoriteHandler) Status(c fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	lessonID := fiber.Params[int](c, "lessonID")
	if lessonID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid lesson id"})
	}
	isFav, _ := h.repo.IsFavorited(c.Context(), userID, lessonID)
	return c.JSON(fiber.Map{"favorited": isFav})
}
