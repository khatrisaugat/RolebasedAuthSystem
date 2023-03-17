package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khatrisaugat/PatternPractise/app/models"
)

func GetAllStudents(c *gin.Context) {
	fmt.Println("GetAllStudents")
	var students []models.Student
	students, err := models.GetAllStudents()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, students)
	}
}

func GetStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	err := student.GetStudent(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

func CreateStudent(c *gin.Context) {
	var student models.Student
	c.BindJSON(&student)
	err := student.SaveStudent()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

func UpdateStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	err := student.GetStudent(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	var updateInput models.Student
	// c.BindJSON(&student)
	c.ShouldBindJSON(&updateInput)
	// fmt.Println(c.Request.Body)
	// fmt.Println("updateInput")
	// fmt.Println(updateInput)
	err = student.UpdateStudent(updateInput)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}

func DeleteStudent(c *gin.Context) {
	var student models.Student
	id := c.Param("id")
	err := student.GetStudent(id)
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	}
	err = student.DeleteStudent()
	if err != nil {
		c.AbortWithStatus(http.StatusNotFound)
	} else {
		c.JSON(http.StatusOK, student)
	}
}
