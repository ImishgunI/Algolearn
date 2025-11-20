package database

import (
	"algolearn/internal/models"
	"context"
	"errors"
	"time"
)

type CommentRepository interface {
	AddComment(ctx context.Context, req *models.CommentRequest) error
	GetComments(ctx context.Context) ([]models.Comment, error)
}

func (d *Database) AddComment(ctx context.Context, req *models.CommentRequest) error {
	var comment models.Comment
	err := d.db.QueryRow(ctx, `
		SELECT u.id FROM users as u WHERE email=$1
	`, req.Email).Scan(&comment.UserID)
	if err != nil {
		return errors.New("ID not found in users")
	}
	err = d.db.QueryRow(ctx, `SELECT id FROM lessons WHERE title=$1`, req.LessonTitle).Scan(&comment.LessonID)
	if err != nil {
		return errors.New("ID not found in lessons")
	}
	comment.Content = req.Content
	comment.CreatedAt = time.Now()
	_, err = d.db.Exec(ctx, `INSERT INTO comments (user_id, lesson_id, content, created_at) VALUES ($1, $2, $3, $4)`,
		comment.UserID, comment.LessonID, comment.Content, comment.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) GetComments(ctx context.Context, title string) ([]models.Comment, error) {
	rows, err := d.db.Query(ctx, `
		SELECT c.content, c.created_at FROM comments
		JOIN lessons l on l.id = c.lesson_id
		WHERE l.title=$1
	`, title)
	if err != nil {
		return nil, err
	}
	var result []models.Comment
	for rows.Next() {
		var c models.Comment
		if err := rows.Scan(&c.Content, &c.CreatedAt); err != nil {
			return nil, err
		}
		result = append(result, c)
	}
	return result, nil
}
