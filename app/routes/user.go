package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khatrisaugat/PatternPractise/app/controllers"
	"github.com/khatrisaugat/PatternPractise/app/middlewares"
)

func UserRoutes(incomingRoutes *gin.Engine) {
	userRoutes := incomingRoutes.Group("/users")
	// userRoutes.POST("/", controllers.CreateUser)
	userRoutes.Use(middlewares.JWTAuthMiddleware())
	userRoutes.GET("/", middlewares.AdminAuth(), controllers.GetAllUsers)
	userRoutes.GET("/:id", middlewares.AdminAuth(), controllers.GetUser)
	userRoutes.PUT("/:id", middlewares.SelfAndSuperAdminOnly(), controllers.UpdateUser)
	userRoutes.DELETE("/:id", middlewares.SelfAndSuperAdminOnly(), controllers.DeleteUser)
}
