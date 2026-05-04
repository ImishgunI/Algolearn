package users

type User struct {
	ID           int    `json:"id" db:"id"`
	Name         string `json:"name" db:"first_name"`
	Surname      string `json:"surname" db:"last_name"`
	Email        string `json:"email" db:"email"`
	PasswordHash string `json:"password_hash" db:"password_hash"`
	Role         string `json:"role" db:"role"`
}
