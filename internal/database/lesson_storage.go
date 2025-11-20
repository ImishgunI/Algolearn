package database

import (
	"algolearn/internal/models"
	"context"
)

type LessonsRepository interface {
	GetLessons(ctx context.Context) ([]*models.Lesson, error)
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
