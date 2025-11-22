package database

import (
	"algolearn/internal/models"
	"context"
	"errors"
)

type AdminRepository interface {
	GetAllUsers(ctx context.Context) ([]models.AdminResponse, error)
	GetStatistics(ctx context.Context) (*models.AdminStatistics, error)
	GetLessonsForAdmin(ctx context.Context) ([]models.AdminLessons, error)
	GetLessonToEdit(ctx context.Context, id int) (*models.AdminLessons, error)
	GetUserById(ctx context.Context, id int) (*models.AdminResponse, error)
	GetAllComments(ctx context.Context) ([]models.CommentGetter, error)
	DeleteUser(ctx context.Context, id int) error
	DeleteLesson(ctx context.Context, id int) error
	DeleteComment(ctx context.Context, id int) error
	AddLesson(ctx context.Context, req *models.LessonCreator) error
	AddUser(ctx context.Context, req *models.UserCreator) error
	UpdateLesson(ctx context.Context, req *models.LessonCreator, id int) error
	UpdateUser(ctx context.Context, req *models.UserCreator, id int) error
}

func (d *Database) GetAllUsers(ctx context.Context) ([]models.AdminResponse, error) {
	rows, err := d.db.Query(ctx, `
		SELECT id, first_name, last_name, email, role 
		FROM users 
		WHERE role <> 'admin'
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	var result []models.AdminResponse
	for rows.Next() {
		var a models.AdminResponse
		if err := rows.Scan(&a.ID, &a.FirstName, &a.LastName, &a.Email, &a.Role); err != nil {
			return nil, err
		}
		result = append(result, a)
	}
	return result, nil
}

func (d *Database) GetStatistics(ctx context.Context) (*models.AdminStatistics, error) {
	var as models.AdminStatistics
	err := d.db.QueryRow(ctx, `
		SELECT 
			(SELECT COUNT(*) FROM users WHERE role <> 'admin') AS usersCount,
			(SELECT COUNT(*) FROM lessons) AS lessonsCount,
			(SELECT COUNT(*) FROM comments) AS commentsCount,
			(SELECT COUNT(*) FROM tasks) AS tasksCount
	`).Scan(&as.UsersCount, &as.LessonsCount, &as.CommentsCount, &as.TasksCount)
	if err != nil {
		return nil, err
	}
	return &as, nil
}

func (d *Database) GetLessonsForAdmin(ctx context.Context) ([]models.AdminLessons, error) {
	rows, err := d.db.Query(ctx, `
		SELECT id, title, category, difficulty, content FROM lessons
		ORDER BY id
	`)
	if err != nil {
		return nil, err
	}
	var result []models.AdminLessons
	for rows.Next() {
		var al models.AdminLessons
		if err := rows.Scan(&al.ID, &al.Title, &al.Category, &al.Difficulty, &al.Content); err != nil {
			return nil, err
		}
		result = append(result, al)
	}
	return result, nil
}

func (d *Database) GetLessonToEdit(ctx context.Context, id int) (*models.AdminLessons, error) {
	var al models.AdminLessons
	err := d.db.QueryRow(ctx, `
		SELECT id, title, category, difficulty, content FROM lessons WHERE id=$1
	`, id).Scan(&al.ID, &al.Title, &al.Category, &al.Difficulty, &al.Content)
	if err != nil {
		return nil, err
	}
	return &al, nil
}

func (d *Database) GetUserById(ctx context.Context, id int) (*models.AdminResponse, error) {
	var ar models.AdminResponse
	err := d.db.QueryRow(ctx, `SELECT id, first_name, last_name, email, role FROM users WHERE id = $1`, id).
		Scan(&ar.ID, &ar.FirstName, &ar.LastName, &ar.Email, &ar.Role)

	if err != nil {
		return nil, err
	}
	return &ar, nil
}

func (d *Database) GetAllComments(ctx context.Context) ([]models.CommentGetter, error) {
	rows, err := d.db.Query(ctx, `
	SELECT c.id, u.first_name, u.last_name, l.title AS lesson_title, c.content, c.created_at FROM comments c
	JOIN users u ON u.id = c.user_id
	JOIN lessons l ON l.id = c.lesson_id
	ORDER BY c.created_at
	`)
	if err != nil {
		return nil, err
	}

	var cg []models.CommentGetter
	for rows.Next() {
		var c models.CommentGetter
		if err := rows.Scan(&c.ID, &c.FirstName, &c.LastName, &c.LessonTitle, &c.Content, &c.CreatedAt); err != nil {
			return nil, err
		}
		cg = append(cg, c)
	}
	return cg, nil
}

func (d *Database) DeleteUser(ctx context.Context, id int) error {
	_, err := d.db.Exec(ctx, `DELETE FROM users WHERE id = $1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteLesson(ctx context.Context, id int) error {
	_, err := d.db.Exec(ctx, `
		DELETE FROM lessons WHERE id=$1
	`, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) DeleteComment(ctx context.Context, id int) error {
	_, err := d.db.Exec(ctx, `DELETE FROM comments WHERE id=$1`, id)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) AddLesson(ctx context.Context, req *models.LessonCreator) error {
	var exist bool
	err := d.db.QueryRow(ctx, `SELECT id FROM lessons WHERE title = $1`, req.Title).Scan(&exist)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("lessons already exist")
	}
	_, err = d.db.Exec(ctx, `
		INSERT INTO lessons (title, category, difficulty, content) VALUES ($1, $2, $3, $4)
	`, req.Title, req.Category, req.Difficulty, req.Content)

	if err != nil {
		return err
	}
	return nil
}

func (d *Database) UpdateLesson(ctx context.Context, req *models.LessonCreator, id int) error {
	_, err := d.db.Exec(ctx, `
	UPDATE lessons
	SET title=$1, category=$2, difficulty=$3, content=$4
	WHERE id=$5
	`, req.Title, req.Category, req.Difficulty, req.Content, id)

	if err != nil {
		return err
	}
	return nil
}

func (d *Database) AddUser(ctx context.Context, req *models.UserCreator) error {
	var exist bool
	err := d.db.QueryRow(ctx, `SELECT id FROM users WHERE email=$1`, req.Email).Scan(&exist)
	if err != nil {
		return err
	}
	if exist {
		return errors.New("User already exist")
	}
	_, err = d.db.Exec(ctx, `INSERT INTO users (first_name, last_name, email, role) VALUES ($1, $2, $3, $4)`, req.FirstName, req.LastName, req.Email, req.Role)
	if err != nil {
		return err
	}
	return nil
}

func (d *Database) UpdateUser(ctx context.Context, req *models.UserCreator, id int) error {
	_, err := d.db.Exec(ctx, `
		UPDATE users
		SET first_name = $1, last_name = $2, email = $3, role = $4
		WHERE id = $5
	`, req.FirstName, req.LastName, req.Email, req.Role, id)
	if err != nil {
		return err
	}
	return nil
}
