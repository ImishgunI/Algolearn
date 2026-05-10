package transport

import (
	"Algolearn/internal/middleware"
	"Algolearn/internal/transport/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func Routes(reg *http.Registration, auth *http.Authorization, exec *http.ExecutionHandler, ls *http.LessonHandler, ch *http.CommentHandler, fh *http.FavoriteHandler) *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	app.Post("/register", reg.SignUp)
	app.Post("/login", auth.SignIn)
	app.Post("/refresh", auth.Refresh)

	app.Get("/me", middleware.JWTMiddleware("secret"), auth.Me)

	app.Post("/execute", exec.Execute)
	app.Get("/execution/:id", exec.Get)

	app.Get("/courses/:courseID/lessons", ls.GetByCourse)
	app.Get("/lessons/:id", ls.GetByID)

	app.Get("/lessons/:lessonID/comments", ch.GetByLesson)
	app.Post("/lessons/:lessonID/comments", middleware.JWTMiddleware("secret"), ch.Create)

	app.Get("/lessons/:lessonID/favorite", middleware.JWTMiddleware("secret"), fh.Status)
	app.Post("/lessons/:lessonID/favorite", middleware.JWTMiddleware("secret"), fh.Toggle)
	return app
}
