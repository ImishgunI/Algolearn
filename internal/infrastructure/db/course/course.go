package course

import (
	"Algolearn/internal/infrastructure/db/sql"
	"Algolearn/internal/learning/course"
	"context"
)

type Repository struct {
	p *sql.Postgres
}

func NewCourseRepository(db *sql.Postgres) *Repository {
	return &Repository{
		p: db,
	}
}

func (r *Repository) GetAll(ctx context.Context) ([]course.Course, error) {
	rows, err := r.p.Pool.Query(ctx, `SELECT id, title, description, created_at FROM courses`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var cr []course.Course
	for rows.Next() {
		var c course.Course
		if err := rows.Scan(&c.ID, &c.Title, &c.Description, &c.CreatedAt); err != nil {
			return nil, err
		}
		cr = append(cr, c)
	}
	return cr, nil
}
func (r *Repository) Create(ctx context.Context, title, description string) (*course.Course, error) {
	var c course.Course
	c.Title = title
	c.Description = description
	err := r.p.Pool.QueryRow(ctx,
		`INSERT INTO courses (title, description, created_at) VALUES ($1, $2, NOW()) RETURNING id, created_at`,
		title, description).Scan(&c.ID, &c.CreatedAt)
	if err != nil {
		return nil, err
	}
	return &c, nil
}
func (r *Repository) Update(ctx context.Context, id int, title, description string) error {
	_, err := r.p.Pool.Exec(ctx,
		`UPDATE courses
		SET title = $1, description = $2
		WHERE id = $3`, title, description, id)
	if err != nil {
		return err
	}
	return nil
}
func (r *Repository) Delete(ctx context.Context, id int) error {
	_, err := r.p.Pool.Exec(ctx, `DELETE FROM courses WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Count(ctx context.Context) (int, error) {
	var c int
	err := r.p.Pool.QueryRow(ctx, `SELECT COUNT(*) FROM courses`).Scan(&c)
	if err != nil {
		return 0, err
	}
	return c, nil
}
