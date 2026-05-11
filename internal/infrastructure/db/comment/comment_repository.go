package comment

import (
	"Algolearn/internal/infrastructure/db/sql"
	"Algolearn/internal/learning/comment"
	"context"
	"time"
)

type Repository struct {
	p *sql.Postgres
}

func NewCommentRepository(db *sql.Postgres) *Repository {
	return &Repository{
		p: db,
	}
}

func (r *Repository) Create(ctx context.Context, c *comment.Comment) error {
	c.CreatedAt = time.Now().UTC()
	_, err := r.p.Pool.Exec(ctx,
		"INSERT INTO comments (lesson_id, user_id, body, created_at) VALUES ($1, $2, $3, $4)",
		c.LessonID, c.UserID, c.Body, c.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) GetByLesson(ctx context.Context, lessonID int) ([]comment.Comment, error) {
	rows, err := r.p.Pool.Query(ctx, `SELECT c.id, c.lesson_id, c.user_id, u.user_name, c.body, c.created_at from comments c
		JOIN users u ON u.id = c.user_id
		WHERE lesson_id = $1
		ORDER BY c.created_at ASC`, lessonID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var comments []comment.Comment
	for rows.Next() {
		var c comment.Comment
		err := rows.Scan(&c.ID, &c.LessonID, &c.UserID, &c.UserName, &c.Body, &c.CreatedAt)
		if err != nil {
			return nil, err
		}
		comments = append(comments, c)
	}
	return comments, nil
}

func (r *Repository) Delete(ctx context.Context, id int) error {
	_, err := r.p.Pool.Exec(ctx, `DELETE FROM comments WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
