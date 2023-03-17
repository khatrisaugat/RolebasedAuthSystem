package controllers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/khatrisaugat/PatternPractise/app/helpers"
	"github.com/khatrisaugat/PatternPractise/app/models"
	"github.com/khatrisaugat/PatternPractise/app/services"
	"golang.org/x/crypto/bcrypt"
)

func RegisterUser(c *gin.Context) {
	var input services.RegisterUserInput
	err := c.ShouldBindJSON(&input)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	isSuperAdmin, err := services.CheckIfSuperAdmin(fmt.Sprint(input.UserTypeID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	if isSuperAdmin {
		c.JSON(http.StatusBadRequest, gin.H{"error": "you can't register as super admin"})
		return
	}

	u := &models.User{
		Email:      input.Email,
		Password:   input.Password,
		UserTypeID: input.UserTypeID,
	}

	err = u.SaveUser()
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": u})
}

func LoginUser(c *gin.Context) {
	var input services.LoginUserInput
	err := c.ShouldBindJSON(&input)
	fmt.Println(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	u := models.User{}
	err = helpers.DB.Model(models.User{}).Where("email = ?", input.Email).First(&u).Error

	// fmt.Println(u)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err = services.CheckStatus(fmt.Sprintf("%v", u.ID))
	if err != nil {
		c.JSON(http.StatusForbidden, gin.H{"error": err.Error()})
		return
	}

	err = helpers.VerifyPassword(u.Password, input.Password)
	// fmt.Println(err)
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login credentials. Please try again"})
		return
	}
	token, err := helpers.GenerateToken(uint(u.ID))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func CurrentUser(c *gin.Context) {
	user_id, err := helpers.ExtractTokenId(c)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	u := models.User{}
	err = helpers.DB.Model(models.User{}).Where("id=?", user_id).First(&u).Error
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": u})
}
