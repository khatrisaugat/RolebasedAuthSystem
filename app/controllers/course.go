package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khatrisaugat/PatternPractise/app/models"
)

func GetAllCourses(c *gin.Context) {
	var courses []models.Course
	courses, err := models.GetAllCourses()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, courses)
	}
}

func GetCourse(c *gin.Context) {
	var course models.Course
	id := c.Param("id")
	err := course.GetCourse(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, course)
	}
}

func CreateCourse(c *gin.Context) {
	var course models.Course
	c.BindJSON(&course)
	err := course.SaveCourse()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, course)
	}
}

func UpdateCourse(c *gin.Context) {
	var course models.Course
	id := c.Param("id")
	err := course.GetCourse(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	var courseUpdate models.Course
	c.BindJSON(&courseUpdate)

	err = course.UpdateCourse(courseUpdate)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, course)
	}
}

func DeleteCourse(c *gin.Context) {
	var course models.Course
	id := c.Param("id")
	err := course.GetCourse(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	err = course.DeleteCourse()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
	}
}
