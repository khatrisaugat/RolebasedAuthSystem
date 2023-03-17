package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khatrisaugat/PatternPractise/app/models"
)

func EnrollStudent(c *gin.Context) {
	var enrolledCourse models.EnrolledCourse
	c.BindJSON(&enrolledCourse)
	err := enrolledCourse.SaveEnrolledCourse()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, enrolledCourse)
	}
}

func RemoveStudentFromCourse(c *gin.Context) {
	var enrolledCourse models.EnrolledCourse
	id := c.Param("id")
	err := enrolledCourse.GetEnrolledCourse(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	err = enrolledCourse.DeleteEnrolledCourse()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, enrolledCourse)
	}
}

func GetCoursesOfStudent(c *gin.Context) {
	var enrolledCourse []models.EnrolledCourse
	id := c.Param("id")
	enrolledCourse, err := models.GetAllEnrolledCoursesForStudent(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, enrolledCourse)
	}
}
