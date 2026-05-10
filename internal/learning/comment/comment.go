package comment

import "time"

type Comment struct {
	ID        int       `json:"id"`
	LessonID  int       `json:"lesson_id"`
	UserID    int       `json:"user_id"`
	UserName  string    `json:"user_name"`
	Body      string    `json:"body"`
	CreatedAt time.Time `json:"created_at"`
}
