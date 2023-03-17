package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khatrisaugat/PatternPractise/app/controllers"
	"github.com/khatrisaugat/PatternPractise/app/middlewares"
)

func StudentRoutes(incomingRoutes *gin.Engine) {
	studentRoutes := incomingRoutes.Group("/students")
	studentRoutes.POST("/", controllers.CreateStudent)
	studentRoutes.Use(middlewares.JWTAuthMiddleware())
	studentRoutes.GET("/", controllers.GetAllStudents)
	studentRoutes.GET("/:id", controllers.GetStudent)
	studentRoutes.PUT("/:id", middlewares.SelfCheck(), controllers.UpdateStudent)
	studentRoutes.DELETE("/:id", middlewares.SelfCheck(), controllers.DeleteStudent)
	studentRoutes.GET("/courses/:id", middlewares.SelfCheck(), controllers.GetCoursesOfStudent)
	studentRoutes.POST("/courses/:id", middlewares.SelfCheck(), controllers.EnrollStudent)
	studentRoutes.DELETE("/courses/:id", middlewares.SelfCheck(), controllers.RemoveStudentFromCourse)

}
