package users

type User struct {
	Name         string `json:"name"`
	Surname      string `json:"surname"`
	Email        string `json:"email"`
	PasswordHash string `json:"password_hash"`
	Role         string `json:"role"`
}
