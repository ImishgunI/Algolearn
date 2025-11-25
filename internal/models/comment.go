package models

import "time"

type Comment struct {
	CreatedAt time.Time `db:"created_at"`
	Content   string    `db:"content" json:"content"`
	UserID    int       `db:"user_id"`
	LessonID  int       `db:"lesson_id"`
}

type CommentRequest struct {
	Email       string `json:"email"`
	LessonTitle string `json:"lesson_title"`
	Content     string `json:"content" db:"content"`
}

type CommentResponse struct {
	CreatedAt time.Time `json:"created_at"`
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Content   string    `json:"content"`
}

type CommentGetter struct {
	CreatedAt   time.Time `json:"created_at"`
	FirstName   string    `json:"first_name"`
	LastName    string    `json:"last_name"`
	LessonTitle string    `json:"lesson_title"`
	Content     string    `json:"content"`
	ID          int       `json:"id"`
}
