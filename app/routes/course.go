package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khatrisaugat/PatternPractise/app/controllers"
	"github.com/khatrisaugat/PatternPractise/app/middlewares"
)

func CourseRoutes(incomingRoutes *gin.Engine) {
	courseRoutes := incomingRoutes.Group("/courses")
	courseRoutes.Use(middlewares.JWTAuthMiddleware())
	courseRoutes.GET("/", controllers.GetAllCourses)
	courseRoutes.GET("/:id", controllers.GetCourse)
	courseRoutes.POST("/", middlewares.AdminAuth(), controllers.CreateCourse)
	courseRoutes.PUT("/:id", middlewares.AdminAuth(), controllers.UpdateCourse)
	courseRoutes.DELETE("/:id", middlewares.AdminAuth(), controllers.DeleteCourse)
}
