package favorite

import (
	"Algolearn/internal/infrastructure/db/sql"
	"context"
)

// type FavoriteRepository interface {
// 	Add(ctx context.Context, userID int, lessonID int) error
// 	Remove(ctx context.Context, userID int, lessonID int) error
// 	IsFavorited(ctx context.Context, userID int, lessonID int) (bool, error)
// 	GetUserFavorites(ctx context.Context, userID int) ([]int, error)
// }

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
	var exist int
	err := r.p.Pool.QueryRow(ctx, "SELECT EXISTS(SELECT 1 FROM favorites WHERE user_id = $1 AND lesson_id = $2)", userID, lessonID).Scan(&exist)
	if err != nil {
		return false, err
	}
	if exist == 1 {
		return true, nil
	} else {
		return false, nil
	}
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
