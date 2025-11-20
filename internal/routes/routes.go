package routes

import (
	"algolearn/internal/controller"
	"algolearn/internal/database"

	"github.com/gin-gonic/gin"
)

func SetRoutes(r *gin.Engine, db *database.Database) {
	h := controller.NewUserHandler(db)
	hc := controller.NewCommentHandler(db)
	hl := controller.NewLessonHandler(db)
	r.POST("/register", h.RegisterController)
	r.POST("/login", h.LoginController)
	r.POST("/newComment", hc.AddComment)
	r.GET("/lessons", hl.GetLessons)
	r.GET("/comments/", hc.GetComments)
}
