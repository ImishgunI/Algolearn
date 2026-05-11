package http

import (
	"Algolearn/internal/infrastructure/db"
	"Algolearn/internal/learning/lesson"
	"strconv"

	"github.com/gofiber/fiber/v3"
)

type AdminHandler struct {
	courseRepo  db.CourseRepository
	lessonRepo  db.LessonRepository
	userRepo    db.UserRepository
	commentRepo db.CommentRepository
	// executionRepo db.ExecutionRepository
}

func NewAdminHandler(cr db.CourseRepository, lr db.LessonRepository, ur db.UserRepository, cmr db.CommentRepository) *AdminHandler {
	return &AdminHandler{
		courseRepo:  cr,
		lessonRepo:  lr,
		userRepo:    ur,
		commentRepo: cmr,
	}
}

// Курсы
func (h *AdminHandler) ListCourses(c fiber.Ctx) error {
	courses, err := h.courseRepo.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(courses)
}

func (h *AdminHandler) CreateCourse(c fiber.Ctx) error {
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}
	if input.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "title is required"})
	}
	course, err := h.courseRepo.Create(c.Context(), input.Title, input.Description)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.Status(fiber.StatusCreated).JSON(course)
}

func (h *AdminHandler) UpdateCourse(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var input struct {
		Title       string `json:"title"`
		Description string `json:"description"`
	}
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}
	if input.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "title is required"})
	}
	if err := h.courseRepo.Update(c.Context(), id, input.Title, input.Description); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "updated"})
}

func (h *AdminHandler) DeleteCourse(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.courseRepo.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "deleted"})
}

// Уроки
func (h *AdminHandler) ListLessons(c fiber.Ctx) error {
	courseID, err := strconv.Atoi(c.Params("courseID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid course id"})
	}
	lessons, err := h.lessonRepo.GetByCourse(c.Context(), courseID)
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(lessons)
}

func (h *AdminHandler) CreateLesson(c fiber.Ctx) error {
	courseID, err := strconv.Atoi(c.Params("courseID"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid course id"})
	}
	var input struct {
		Title         string `json:"title"`
		Theory        string `json:"theory"`
		AlgorithmType string `json:"algorithm_type"`
		OrderIndex    int    `json:"order_index"`
	}
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}
	if input.Title == "" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "title is required"})
	}

	l := lesson.Lesson{
		CourseID:      courseID,
		Title:         input.Title,
		Theory:        input.Theory,
		AlgorithmType: input.AlgorithmType,
		OrderIndex:    input.OrderIndex,
	}
	if err := h.lessonRepo.Create(c.Context(), &l); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	// объект l должен пополниться полем ID после вставки (если репозиторий возвращает его через RETURNING)
	return c.Status(fiber.StatusCreated).JSON(l)
}
func (h *AdminHandler) UpdateLesson(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var input struct {
		Title         string `json:"title"`
		Theory        string `json:"theory"`
		AlgorithmType string `json:"algorithm_type"`
		OrderIndex    int    `json:"order_index"`
		CourseID      int    `json:"course_id"`
	}
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}

	updated := lesson.Lesson{
		ID:            id,
		CourseID:      input.CourseID,
		Title:         input.Title,
		Theory:        input.Theory,
		AlgorithmType: input.AlgorithmType,
		OrderIndex:    input.OrderIndex,
	}
	if err := h.lessonRepo.Update(c.Context(), id, &updated); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "updated"})
}

func (h *AdminHandler) DeleteLesson(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.lessonRepo.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "deleted"})
}

// Пользователи
func (h *AdminHandler) ListUsers(c fiber.Ctx) error {
	users, err := h.userRepo.GetAll(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(users)
}
func (h *AdminHandler) UpdateUserRole(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	var input struct {
		Role string `json:"role"`
	}
	if err := c.Bind().Body(&input); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid input"})
	}
	if input.Role != "user" && input.Role != "admin" {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "role must be user or admin"})
	}
	if err := h.userRepo.UpdateRole(c.Context(), id, input.Role); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "role updated"})
}

// Комментарии
func (h *AdminHandler) ListComments(c fiber.Ctx) error {
	comments, err := h.commentRepo.GetAllComments(c.Context())
	if err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(comments)
}

func (h *AdminHandler) DeleteComment(c fiber.Ctx) error {
	id, err := strconv.Atoi(c.Params("id"))
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{"error": "invalid id"})
	}
	if err := h.commentRepo.Delete(c.Context(), id); err != nil {
		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"error": err.Error()})
	}
	return c.JSON(fiber.Map{"message": "deleted"})
}

// Статистика
func (h *AdminHandler) Stats(c fiber.Ctx) error {
	usersCount, _ := h.userRepo.Count(c.Context())
	coursesCount, _ := h.courseRepo.Count(c.Context())
	lessonsCount, _ := h.lessonRepo.Count(c.Context())
	commentsCount, _ := h.commentRepo.Count(c.Context())

	return c.JSON(fiber.Map{
		"users":    usersCount,
		"courses":  coursesCount,
		"lessons":  lessonsCount,
		"comments": commentsCount,
		// добавишь executions позже
	})
}
