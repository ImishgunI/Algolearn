package users

type User struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"user_name"`
	Surname      string `json:"surname" db:"user_surname"`
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	Role         string `json:"role" db:"user_role"`
}
