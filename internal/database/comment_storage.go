package database

import (
	"algolearn/internal/models"
	"context"
	"time"
)

type CommentRepository interface {
	AddComment(ctx context.Context, req *models.CommentRequest) (models.Comment, error)
}

func (d *Database) AddComment(ctx context.Context, req *models.CommentRequest) (*models.Comment, error) {
	var comment models.Comment
	err := d.db.QueryRow(ctx, `
		SELECT u.id FROM users as u WHERE email=$1
		UNION SELECT l.id FROM lessons as l WHERE title=$2
	`, req.Email, req.LessonTitle).Scan(&comment.UserID, &comment.LessonID)
	if err != nil {
		return nil, err
	}
	comment.Content = req.Content
	comment.CreatedAt = time.Now()
	_, err = d.db.Exec(ctx, `INSERT INTO comments (user_id, lesson_id, content, created_at) VALUES ($1, $2, $3, $4)`,
		comment.UserID, comment.LessonID, comment.Content, comment.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &comment, nil
}
