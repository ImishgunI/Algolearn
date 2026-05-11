package http

import (
	"Algolearn/internal/infrastructure/db"
	"Algolearn/internal/users"

	"github.com/gofiber/fiber/v3"
)

type ProfileHandler struct {
	repo         db.UserRepository
	favoriteRepo db.FavoriteRepository
	commentRepo  db.CommentRepository
}

func NewProfileHandler(repo db.UserRepository, favRepo db.FavoriteRepository) *ProfileHandler {
	return &ProfileHandler{repo: repo, favoriteRepo: favRepo}
}

func (h *ProfileHandler) UpdateProfile(c fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	var input struct {
		Name    string `json:"user_name"`
		Surname string `json:"user_surname"`
		Email   string `json:"email"`
	}
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}
	if input.Name == "" || input.Surname == "" || input.Email == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "fields required"})
	}
	err := h.repo.UpdateProfile(c.Context(), userID, input.Name, input.Surname, input.Email)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "updated"})
}

func (h *ProfileHandler) UpdatePassword(c fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	var input struct {
		OldPassword string `json:"old_password"`
		NewPassword string `json:"new_password"`
	}
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}
	user, err := h.repo.GetByID(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "user not found"})
	}
	if !users.CheckPasswordHash(input.OldPassword, user.PasswordHash) {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "old password is incorrect"})
	}
	hash, err := users.HashPassword(input.NewPassword)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "hashing failed"})
	}
	if err := h.repo.UpdatePassword(c.Context(), userID, hash); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "password updated"})
}

func (h *ProfileHandler) GetStats(c fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	favCount, err := h.favoriteRepo.CountByUser(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{
		"algorithms_executed": 0, // позже привяжешь
		"favorites_count":     favCount,
	})
}

func (h *ProfileHandler) GetFavorites(c fiber.Ctx) error {
	userID := c.Locals("user_id").(int)
	favorites, err := h.favoriteRepo.GetFavoriteLessons(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(favorites)
}
