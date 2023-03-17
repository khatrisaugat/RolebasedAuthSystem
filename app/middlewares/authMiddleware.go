package middlewares

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/khatrisaugat/PatternPractise/app/helpers"
	"github.com/khatrisaugat/PatternPractise/app/services"
)

func JWTAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		err := helpers.TokenValid(c)
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			c.Abort()
			return
		}
		c.Next()
	}
}

func SuperUserAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id, err := helpers.ExtractTokenId(c)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			c.Abort()
			return
		}
		role, err := services.GetUserType(fmt.Sprint(user_id))
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			c.Abort()
			return
		}
		if role.User_Type != "super admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			c.Abort()
			return
		}
		c.Next()
	}
}

func AdminAuth() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id, err := helpers.ExtractTokenId(c)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			c.Abort()
			return
		}
		role, err := services.GetUserType(fmt.Sprint(user_id))
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			c.Abort()
			return
		}
		if role.User_Type != "admin" {
			if role.User_Type != "super admin" {
				c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
				c.Abort()
				return
			}
		}
		c.Next()
	}
}

func SelfCheck() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id, err := helpers.ExtractTokenId(c)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			c.Abort()
			return
		}
		id := c.Param("id")
		role, err := services.GetUserType(fmt.Sprint(user_id))
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			c.Abort()
			return
		}
		if role.User_Type != "admin" || fmt.Sprint(user_id) != id || role.User_Type != "super admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}

func SelfAndSuperAdminOnly() gin.HandlerFunc {
	return func(c *gin.Context) {
		user_id, err := helpers.ExtractTokenId(c)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			c.Abort()
			return
		}
		id := c.Param("id")
		role, err := services.GetUserType(fmt.Sprint(user_id))
		fmt.Println(role.User_Type)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			c.Abort()
			return
		}
		intId, err := strconv.Atoi(id)
		if err != nil {
			c.JSON(http.StatusForbidden, gin.H{"error": "Bad request"})
			c.Abort()
			return
		}
		fmt.Println(user_id)
		if user_id != uint64(intId) && role.User_Type != "super admin" {
			c.JSON(http.StatusForbidden, gin.H{"error": "You are not authorized to access this resource"})
			c.Abort()
			return
		}

		c.Next()
	}
}
