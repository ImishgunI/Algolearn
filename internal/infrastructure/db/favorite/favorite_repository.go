package favorite

import (
	"Algolearn/internal/infrastructure/db/sql"
	"Algolearn/internal/learning/favorite"
	"context"
)

type Repository struct {
	p *sql.Postgres
}

func NewFavoriteRepo(db *sql.Postgres) *Repository {
	return &Repository{
		p: db,
	}
}

func (r *Repository) Add(ctx context.Context, userID int, lessonID int) error {
	_, err := r.p.Pool.Exec(ctx, "INSERT INTO favorites (user_id, lesson_id) VALUES ($1, $2)", userID, lessonID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) Remove(ctx context.Context, userID int, lessonID int) error {
	_, err := r.p.Pool.Exec(ctx, "DELETE FROM favorites WHERE user_id = $1 AND lesson_id = $2", userID, lessonID)
	if err != nil {
		return err
	}
	return nil
}

func (r *Repository) IsFavorited(ctx context.Context, userID int, lessonID int) (bool, error) {
	var exist bool
	err := r.p.Pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM favorites WHERE user_id = $1 AND lesson_id = $2)", userID, lessonID).Scan(&exist)
	if err != nil {
		return false, err
	}
	return exist, nil
}

func (r *Repository) GetUserFavorites(ctx context.Context, userID int) ([]int, error) {
	rows, err := r.p.Pool.Query(ctx, "SELECT lesson_id FROM favorites WHERE user_id = $1", userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var ids []int
	for rows.Next() {
		var id int
		err := rows.Scan(&id)
		if err != nil {
			return nil, err
		}
		ids = append(ids, id)
	}
	return ids, nil
}

func (r *Repository) GetFavoriteLessons(ctx context.Context, userID int) ([]favorite.LessonInfo, error) {
	const query = `SELECT f.lesson_id, l.title FROM favorites f
					JOIN lessons l ON f.lesson_id = l.id
					WHERE user_id = $1`

	rows, err := r.p.Pool.Query(ctx, query, userID)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var li []favorite.LessonInfo
	for rows.Next() {
		var l favorite.LessonInfo
		if err := rows.Scan(&l.LessonID, &l.Title); err != nil {
			return nil, err
		}
		li = append(li, l)
	}
	return li, nil
}

func (r *Repository) CountByUser(ctx context.Context, userID int) (int, error) {
	const query = `SELECT COUNT(lesson_id) FROM favorites WHERE user_id = $1`
	var count int
	err := r.p.Pool.QueryRow(ctx, query, userID).Scan(&count)
	if err != nil {
		return 0, err
	}
	return count, nil
}
