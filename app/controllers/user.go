package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khatrisaugat/PatternPractise/app/models"
)

func CreateUser(c *gin.Context) {
	var user models.User
	c.ShouldBindJSON(&user)
	err := user.SaveUser()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""
	c.JSON(200, gin.H{"data": user})
}

func GetAllUsers(c *gin.Context) {
	var users []models.User
	// Get all users from db
	users, err := models.GetAllUsers()
	for _, u := range users {
		u.Password = ""
	}
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}

	c.JSON(200, gin.H{"data": users})
}

func GetUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	err := user.GetUserWithTypeById(id)
	user.Password = ""
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": user})
}

func UpdateUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	err := user.GetUserWithTypeById(id)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var updateInput models.User
	c.ShouldBindJSON(&updateInput)

	err = user.UpdateUser(updateInput)
	fmt.Println(updateInput)
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	user.Password = ""
	c.JSON(200, gin.H{"data": user})
}

func DeleteUser(c *gin.Context) {
	var user models.User
	id := c.Param("id")
	err := user.GetUserWithTypeById(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	err = user.DeleteUser()
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	c.JSON(200, gin.H{"data": "User deleted successfully"})
}
