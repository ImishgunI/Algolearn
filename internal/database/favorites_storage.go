package database

import (
	"algolearn/internal/models"
	"context"
)

type FavoritesRepository interface {
	AddFavoriteLesson(ctx context.Context, email string, title string) error
	GetFavoritesByEmail(ctx context.Context, email string) ([]models.FavoritesList, error)
}

func (d *Database) AddFavoriteLesson(ctx context.Context, email string, title string) error {
	type Result struct {
		id  int
		err error
	}
	user_id_ch := make(chan Result, 1)
	lesson_id_ch := make(chan Result, 1)
	go func() {
		var user_id int
		err := d.db.QueryRow(ctx, `SELECT id FROM users WHERE email=$1`, email).Scan(&user_id)
		user_id_ch <- Result{id: user_id, err: err}
	}()
	go func() {
		var lesson_id int
		err := d.db.QueryRow(ctx, `SELECT id FROM lessons WHERE title=$1`, title).Scan(&lesson_id)
		lesson_id_ch <- Result{id: lesson_id, err: err}
	}()
	userRes := <-user_id_ch
	if userRes.err != nil {
		return userRes.err
	}
	lessonRes := <-lesson_id_ch
	if lessonRes.err != nil {
		return lessonRes.err
	}
	_, err := d.db.Exec(ctx, `INSERT INTO favorites (user_id, lesson_id) VALUES ($1, $2)`, userRes.id, lessonRes.id)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) GetFavoritesByEmail(ctx context.Context, email string) ([]models.FavoritesList, error) {
	rows, err := d.db.Query(ctx, `
	SELECT f.lesson_id, l.title FROM favorites f
	JOIN lessons as l ON l.id = f.lesson_id
	JOIN users as u ON u.id = f.user_id
	WHERE u.email = $1
	`, email)
	if err != nil {
		return nil, err
	}
	var result []models.FavoritesList
	for rows.Next() {
		var fl models.FavoritesList
		if err := rows.Scan(&fl.LessonId, &fl.LessonTitle); err != nil {
			return nil, err
		}
		result = append(result, fl)
	}
	return result, nil
}
