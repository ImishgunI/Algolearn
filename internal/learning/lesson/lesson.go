package lesson

import "time"

type Lesson struct {
	ID            int       `json:"id"`
	CourseID      int       `json:"course_id"`
	Title         string    `json:"title"`
	Theory        string    `json:"theory"`
	AlgorithmType string    `json:"algorithm_type"`
	OrderIndex    int       `json:"order_index"`
	CreatedAt     time.Time `json:"created_at"`
}
