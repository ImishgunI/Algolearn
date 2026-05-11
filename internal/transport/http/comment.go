package http

import (
	"Algolearn/internal/infrastructure/db"
	"Algolearn/internal/learning/comment"

	"github.com/gofiber/fiber/v3"
)

type CommentHandler struct {
	repo     db.CommentRepository
	userRepo db.UserRepository
}

func NewCommentHandler(repo db.CommentRepository, userRepo db.UserRepository) *CommentHandler {
	return &CommentHandler{repo: repo, userRepo: userRepo}
}

func (h *CommentHandler) Create(c fiber.Ctx) error {
	userID, ok := c.Locals("user_id").(int)
	if !ok {
		return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{"error": "unauthorized"})
	}
	lessonID := fiber.Params[int](c, "lessonID")
	if lessonID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid lesson id"})
	}
	var input struct {
		Body string `json:"body"`
	}
	if err := c.Bind().Body(&input); err != nil || input.Body == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "body is required"})
	}
	user, err := h.userRepo.GetByID(c.Context(), userID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": "user not found"})
	}
	cmt := &comment.Comment{
		LessonID: lessonID,
		UserID:   userID,
		UserName: user.Name,
		Body:     input.Body,
	}
	if err := h.repo.Create(c.Context(), cmt); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(cmt)
}

func (h *CommentHandler) GetByLesson(c fiber.Ctx) error {
	lessonID := fiber.Params[int](c, "lessonID")
	if lessonID == 0 {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid lesson id"})
	}
	comments, err := h.repo.GetByLesson(c.Context(), lessonID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(comments)
}
