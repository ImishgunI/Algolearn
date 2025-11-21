package services

import (
	"golang.org/x/crypto/bcrypt"
)

func HashPassword(password string) (string, error) {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 9)
	return string(bytes), err
}

func CheckPasswordHash(password, hash string) (bool, error) {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	return err == nil, err
}

func CheckRoleForNewUsers(email string) string {
	var role string
	switch email {
	case "adminemail@gmail.com":
		role = "admin"
	case "manageremail@gmail.com":
		role = "manager"
	default:
		role = "student"
	}
	return role
}
