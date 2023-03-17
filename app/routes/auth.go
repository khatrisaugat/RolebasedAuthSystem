package routes

import (
	"github.com/gin-gonic/gin"
	"github.com/khatrisaugat/PatternPractise/app/controllers"
	"github.com/khatrisaugat/PatternPractise/app/middlewares"
)

func AuthRoutes(incomingRoutes *gin.Engine) {
	authRoutes := incomingRoutes.Group("/auth")
	authRoutes.POST("/register", controllers.RegisterUser)
	authRoutes.POST("/login", controllers.LoginUser)
	authRoutes.GET("/current-user", middlewares.JWTAuthMiddleware(), controllers.CurrentUser)
}
