package database

import (
	"algolearn/internal/models"
	"context"
	"time"
)

type LessonsRepository interface {
	GetLessons(ctx context.Context) ([]*models.Lesson, error)
	SetLessonComplete(ctx context.Context, email, lessonTitle string) error
}

func (d *Database) GetLessons(ctx context.Context) ([]models.Lesson, error) {
	rows, err := d.db.Query(ctx, `SELECT title, category, difficulty, content FROM lessons`)
	if err != nil {
		return nil, err
	}
	var result []models.Lesson
	for rows.Next() {
		var l models.Lesson
		if err := rows.Scan(&l.Title, &l.Category, &l.Difficulty, &l.Content); err != nil {
			return nil, err
		}
		result = append(result, l)
	}
	return result, nil
}

func (d *Database) SetLessonComplete(ctx context.Context, email, lessonTitle string) error {
	var user_id int
	err := d.db.QueryRow(ctx, `SELECT id FROM users where email = $1`, email).Scan(&user_id)

	if err != nil {
		return err
	}
	completed_at := time.Now()
	_, err = d.db.Exec(ctx, `INSERT INTO completed_lessons (user_id, lesson_title, completed_at) VALUES ($1, $2, $3)`, user_id, lessonTitle, completed_at)
	if err != nil {
		return err
	}
	return nil
}
