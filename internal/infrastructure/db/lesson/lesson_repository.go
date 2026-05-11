package lesson

import (
	"Algolearn/internal/infrastructure/db/sql"
	"Algolearn/internal/learning/lesson"
	"context"
	"time"
)

type Repository struct {
	p *sql.Postgres
}

func NewLessonRepo(db *sql.Postgres) *Repository {
	return &Repository{
		p: db,
	}
}

func (r *Repository) GetByCourse(ctx context.Context, courseID int) ([]lesson.Lesson, error) {
	rows, err := r.p.Pool.Query(ctx,
		"SELECT id, course_id, title, theory, algorithm_type, order_index, created_at FROM lessons WHERE course_id = $1 ORDER BY order_index",
		courseID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var lessons []lesson.Lesson
	for rows.Next() {
		var l lesson.Lesson
		err := rows.Scan(&l.ID, &l.CourseID, &l.Title, &l.Theory, &l.AlgorithmType, &l.OrderIndex, &l.CreatedAt)
		if err != nil {
			return nil, err
		}
		lessons = append(lessons, l)
	}
	return lessons, nil
}

func (r *Repository) GetByID(ctx context.Context, id int) (*lesson.Lesson, error) {
	var l lesson.Lesson
	err := r.p.Pool.QueryRow(ctx,
		"SELECT id, course_id, title, theory, algorithm_type, order_index, created_at from lessons WHERE id = $1",
		id).Scan(&l.ID, &l.CourseID, &l.Title, &l.Theory, &l.AlgorithmType, &l.OrderIndex, &l.CreatedAt)

	if err != nil {
		return nil, err
	}
	return &l, nil
}

func (r *Repository) Create(ctx context.Context, lesson *lesson.Lesson) error {
	lesson.CreatedAt = time.Now().UTC()
	_, err := r.p.Pool.Exec(ctx, `
	INSERT INTO lessons (course_id, title, theory, algorithm_type, order_index, created_at)
	VALUES ($1, $2, $3, $4, $5, $6)`, lesson.CourseID, lesson.Title, lesson.Theory, lesson.AlgorithmType, lesson.OrderIndex, lesson.CreatedAt)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) Update(ctx context.Context, id int, lesson *lesson.Lesson) error {
	lesson.CreatedAt = time.Now().UTC()
	_, err := r.p.Pool.Exec(ctx,
		`UPDATE lessons
		SET course_id = $1, title = $2, theory = $3, algorithm_type = $4, order_index = $5, created_at = $6
		WHERE id = $7`, lesson.CourseID, lesson.Title, lesson.Theory, lesson.AlgorithmType, lesson.OrderIndex, lesson.CreatedAt, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) Delete(ctx context.Context, id int) error {
	_, err := r.p.Pool.Exec(ctx, `DELETE FROM lessons WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}
