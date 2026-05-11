package db

import (
	"Algolearn/internal/learning/comment"
	"Algolearn/internal/learning/course"
	"Algolearn/internal/learning/favorite"
	"Algolearn/internal/learning/lesson"
	"Algolearn/internal/users"
	"context"
	"time"
)

type UserRepository interface {
	CreateUser(ctx context.Context, user *users.User) error
	GetByEmail(ctx context.Context, email string) (*users.User, error)
	GetByID(ctx context.Context, id int) (*users.User, error)
	UpdateProfile(ctx context.Context, userID int, name, surname, email string) error
	UpdatePassword(ctx context.Context, userID int, hash string) error
	GetAll(ctx context.Context) ([]users.User, error)
	UpdateRole(ctx context.Context, userID int, role string) error
}

type SessionRepository interface {
	CreateSession(ctx context.Context, userID int, token string, expiresAt time.Time) error
	GetSession(ctx context.Context, token string) (int, error)
	DeleteSession(ctx context.Context, token string) error
}

type LessonRepository interface {
	GetByCourse(ctx context.Context, courseID int) ([]lesson.Lesson, error)
	GetByID(ctx context.Context, id int) (*lesson.Lesson, error)
	Create(ctx context.Context, lesson *lesson.Lesson) error
	Update(ctx context.Context, id int, lesson *lesson.Lesson) error
	Delete(ctx context.Context, id int) error
}

type CommentRepository interface {
	Create(ctx context.Context, c *comment.Comment) error
	GetByLesson(ctx context.Context, lessonID int) ([]comment.Comment, error)
	Delete(ctx context.Context, id int) error
}

type FavoriteRepository interface {
	Add(ctx context.Context, userID int, lessonID int) error
	Remove(ctx context.Context, userID int, lessonID int) error
	IsFavorited(ctx context.Context, userID int, lessonID int) (bool, error)
	GetUserFavorites(ctx context.Context, userID int) ([]int, error)
	CountByUser(ctx context.Context, userID int) (int, error)
	GetFavoriteLessons(ctx context.Context, userID int) ([]favorite.LessonInfo, error)
}

type CourseRepository interface {
	GetAll(ctx context.Context) ([]course.Course, error)
	Create(ctx context.Context, title, description string) (*course.Course, error)
	Update(ctx context.Context, id int, title, description string) error
	Delete(ctx context.Context, id int) error
}
