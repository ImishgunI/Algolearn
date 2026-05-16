package transport

import (
	"Algolearn/internal/middleware"
	"Algolearn/internal/transport/http"

	"github.com/gofiber/fiber/v3"
	"github.com/gofiber/fiber/v3/middleware/cors"
	"github.com/gofiber/fiber/v3/middleware/logger"
)

func Routes(reg *http.Registration, auth *http.Authorization,
	exec *http.ExecutionHandler, ls *http.LessonHandler,
	ch *http.CommentHandler, fh *http.FavoriteHandler, ph *http.ProfileHandler,
	adminHandler *http.AdminHandler, custom *http.CustomExecutionHandler) *fiber.App {
	app := fiber.New()
	app.Use(logger.New())
	app.Use(cors.New())
	app.Post("/register", reg.SignUp)
	app.Post("/login", auth.SignIn)
	app.Post("/refresh", auth.Refresh)

	app.Get("/me", middleware.JWTMiddleware("secret"), auth.Me)

	app.Post("/execute", exec.Execute)
	app.Get("/execution/:id", exec.Get)
	app.Post("/execute-custom", middleware.JWTMiddleware("secret"), custom.Execute)

	app.Get("/courses/:courseID/lessons", ls.GetByCourse)
	app.Get("/lessons/:id", ls.GetByID)

	app.Get("/lessons/:lessonID/comments", ch.GetByLesson)
	app.Post("/lessons/:lessonID/comments", middleware.JWTMiddleware("secret"), ch.Create)

	app.Get("/lessons/:lessonID/favorite", middleware.JWTMiddleware("secret"), fh.Status)
	app.Post("/lessons/:lessonID/favorite", middleware.JWTMiddleware("secret"), fh.Toggle)

	app.Get("/me", middleware.JWTMiddleware("secret"), auth.Me)
	app.Put("/profile", middleware.JWTMiddleware("secret"), ph.UpdateProfile)
	app.Put("/profile/password", middleware.JWTMiddleware("secret"), ph.UpdatePassword)
	app.Get("/profile/stats", middleware.JWTMiddleware("secret"), ph.GetStats)
	app.Get("/profile/favorites", middleware.JWTMiddleware("secret"), ph.GetFavorites)

	app.Get("/courses", adminHandler.ListCourses)

	adminGroup := app.Group("/admin", middleware.JWTMiddleware("secret"), middleware.RequireRole("admin"))

	adminGroup.Get("/courses", adminHandler.ListCourses)
	adminGroup.Post("/courses", adminHandler.CreateCourse)
	adminGroup.Put("/courses/:id", adminHandler.UpdateCourse)
	adminGroup.Delete("/courses/:id", adminHandler.DeleteCourse)

	adminGroup.Get("/courses/:courseID/lessons", adminHandler.ListLessons)
	adminGroup.Post("/courses/:courseID/lessons", adminHandler.CreateLesson)
	adminGroup.Put("/lessons/:id", adminHandler.UpdateLesson)
	adminGroup.Delete("/lessons/:id", adminHandler.DeleteLesson)

	adminGroup.Get("/users", adminHandler.ListUsers)
	adminGroup.Put("/users/:id/role", adminHandler.UpdateUserRole)

	adminGroup.Get("/comments", adminHandler.ListComments)
	adminGroup.Delete("/comments/:id", adminHandler.DeleteComment)

	adminGroup.Get("/stats", adminHandler.Stats)
	return app
}
