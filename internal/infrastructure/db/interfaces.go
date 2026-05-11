package db

import (
	"Algolearn/internal/learning/comment"
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
}

type SessionRepository interface {
	CreateSession(ctx context.Context, userID int, token string, expiresAt time.Time) error
	GetSession(ctx context.Context, token string) (int, error)
	DeleteSession(ctx context.Context, token string) error
}

type LessonRepository interface {
	GetByCourse(ctx context.Context, courseID int) ([]lesson.Lesson, error)
	GetByID(ctx context.Context, id int) (*lesson.Lesson, error)
}

type CommentRepository interface {
	Create(ctx context.Context, c *comment.Comment) error
	GetByLesson(ctx context.Context, lessonID int) ([]comment.Comment, error)
}

type FavoriteRepository interface {
	Add(ctx context.Context, userID int, lessonID int) error
	Remove(ctx context.Context, userID int, lessonID int) error
	IsFavorited(ctx context.Context, userID int, lessonID int) (bool, error)
	GetUserFavorites(ctx context.Context, userID int) ([]int, error)
	CountByUser(ctx context.Context, userID int) (int, error)
	GetFavoriteLessons(ctx context.Context, userID int) ([]favorite.LessonInfo, error)
}
