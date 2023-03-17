package services

import (
	"errors"

	"github.com/khatrisaugat/PatternPractise/app/models"
)

type RegisterUserInput struct {
	Email      string `json:"email" binding:"required"`
	Password   string `json:"password" binding:"required"`
	UserTypeID int    `json:"ut_id"`
}

type LoginUserInput struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

func GetUserType(id string) (models.UserType, error) {
	u := models.User{}
	err := u.GetUserWithTypeById(id)
	if err != nil {
		return models.UserType{}, err
	}
	return u.UserType, nil
}

func CheckStatus(id string) error {
	u := models.User{}
	err := u.GetUserWithTypeById(id)
	if err != nil {
		return err
	}
	if !u.Status {
		err = errors.New("user is not active")
		return err
	}
	return nil
}

func CheckIfSuperAdmin(id string) (bool, error) {
	u := models.UserType{}
	err := u.GetUserType(id)
	if err != nil {
		return false, err
	}
	if u.User_Type != "super admin" {
		return false, nil
	}
	return true, nil
}
