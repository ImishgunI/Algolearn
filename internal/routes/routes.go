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
	ha := controller.NewAdminHandler(db)
	r.POST("/register", h.RegisterController)
	r.POST("/login", h.LoginController)
	r.POST("/newComment", hc.AddComment)
	r.POST("/admin/add/lessons", ha.AddLesson)
	r.POST("/admin/add/users", ha.AddUser)
	r.GET("/lessons", hl.GetLessons)
	r.GET("/comments/", hc.GetComments)
	r.GET("/admin/users", ha.GetAllUsers)
	r.GET("/admin/stats", ha.GetStatistics)
	r.GET("/admin/lessons", ha.GetLessonsForAdmin)
	r.GET("/admin/lessons/:lessonId", ha.GetLessonToEdit)
	r.GET("/admin/users/:userId", ha.GetUserById)
	r.GET("/admin/comments", ha.GetAllComments)
	r.DELETE("/admin/lessons/:lessonId", ha.DeleteLesson)
	r.DELETE("/admin/users/:userId", ha.DeleteUser)
	r.PUT("/admin/update/lessons/:lessonId", ha.UpdateLesson)
	r.PUT("/admin/update/users/:userId", ha.UpdateUser)
}
