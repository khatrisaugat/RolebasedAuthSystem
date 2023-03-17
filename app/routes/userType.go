package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khatrisaugat/PatternPractise/app/controllers"
	"github.com/khatrisaugat/PatternPractise/app/middlewares"
)

func UserTypeRoutes(incomingRoutes *gin.Engine) {
	userTypeRoutes := incomingRoutes.Group("/user-types")
	// incomingRoutes.Use(middlewares.JWTAuthMiddleware())
	// userTypeRoutes.Use(middlewares.AdminAuth())
	userTypeRoutes.Use(middlewares.SuperUserAuth())
	userTypeRoutes.POST("/", controllers.CreateUserType)
	userTypeRoutes.DELETE("/:id", controllers.DeleteUserType)
}
