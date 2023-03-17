package models

import (
	"github.com/khatrisaugat/PatternPractise/app/helpers"
)

type UserType struct {
	ID        int    `json:"id" gorm:"primaryKey"`
	User_Type string `json:"user_type" gorm:"size:255"`
}

func (ut *UserType) GetUserType(id string) error {
	helpers.DB.AutoMigrate(&UserType{})
	err := helpers.DB.Model(UserType{}).Where("id = ?", id).First(&ut).Error
	if err != nil {
		return err
	}
	return nil
}

func (ut *UserType) SaveUserType() error {
	helpers.DB.AutoMigrate(&UserType{})
	err := helpers.DB.Create(&ut).Error
	if err != nil {
		return err
	}
	return nil
}

func (ut *UserType) DeleteUserType() error {
	helpers.DB.AutoMigrate(&UserType{})
	err := helpers.DB.Delete(&ut).Error
	if err != nil {
		return err
	}
	return nil
}
