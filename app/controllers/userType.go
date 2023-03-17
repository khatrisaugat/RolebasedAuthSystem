package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khatrisaugat/PatternPractise/app/models"
)

func CreateUserType(c *gin.Context) {
	var userType models.UserType
	c.BindJSON(&userType)
	err := userType.SaveUserType()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User Type created successfully", "data": userType})
}

func DeleteUserType(c *gin.Context) {
	var userType models.UserType
	id := c.Param("id")
	err := userType.GetUserType(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	err = userType.DeleteUserType()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "User Type deleted successfully", "data": userType})
	}
}
