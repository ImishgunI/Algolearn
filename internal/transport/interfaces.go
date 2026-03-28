package transport

import "github.com/gofiber/fiber/v3"

type Registrator interface {
	SignUp(c fiber.Ctx) error
}

type Authorizer interface {
	SignIn(c fiber.Ctx) error
}

type UserRegistrationInfo struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}
