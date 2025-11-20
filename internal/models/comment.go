package models

import "time"

type Comment struct {
	UserID    int       `db:"user_id"`
	LessonID  int       `db:"lesson_id"`
	Content   string    `db:"content" json:"content"`
	CreatedAt time.Time `db:"created_at"`
}

type CommentRequest struct {
	Email       string `json:"email"`
	LessonTitle string `json:"lesson_title"`
	Content     string `json:"content" db:"content"`
}

type CommentResponse struct {
	FirstName string    `json:"first_name"`
	LastName  string    `json:"last_name"`
	Content   string    `json:"content"`
	CreatedAt time.Time `json:"created_at"`
}
