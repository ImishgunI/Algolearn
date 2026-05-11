package favorite

type Favorite struct {
	UserID   int `json:"user_id"`
	LessonID int `json:"lesson_id"`
}

type LessonInfo struct {
	LessonID int    `json:"lesson_id"`
	Title    string `json:"title"`
}
