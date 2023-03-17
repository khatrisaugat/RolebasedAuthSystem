package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"github.com/khatrisaugat/PatternPractise/app/helpers"
	"github.com/khatrisaugat/PatternPractise/app/routes"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	helpers.SetupDb()
}

func main() {
	router := gin.Default()
	router.GET("/", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "Welcome"})
	})
	//auth routes
	routes.AuthRoutes(router)

	//user type routes
	routes.UserTypeRoutes(router)

	//user routes
	routes.UserRoutes(router)

	//student routes
	routes.StudentRoutes(router)

	//course routes
	routes.CourseRoutes(router)

	fmt.Println("Server is running on port " + os.Getenv("PORT"))
	router.Run(":" + os.Getenv("PORT"))
}
