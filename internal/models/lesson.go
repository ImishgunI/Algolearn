package models

type Lesson struct {
	Title      string `json:"title" db:"title"`
	Category   string `json:"category" db:"category"`
	Difficulty string `json:"difficulty" db:"difficulty"`
	Content    string `json:"content" db:"content"`
}

type AdminLessons struct {
	Title      string `json:"title" db:"title"`
	Category   string `json:"category" db:"category"`
	Difficulty string `json:"difficulty" db:"difficulty"`
	Content    string `json:"content" db:"content"`
	ID         int    `json:"id"`
}

type LessonCreator struct {
	Title      string `json:"title" db:"title"`
	Category   string `json:"category" db:"category"`
	Difficulty string `json:"difficulty" db:"difficulty"`
	Content    string `json:"content" db:"content"`
}

func NewLesson() *Lesson {
	return &Lesson{}
}
