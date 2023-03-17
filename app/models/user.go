package models

import (
	"github.com/khatrisaugat/PatternPractise/app/helpers"
	"gorm.io/gorm"
)

type User struct {
	ID            int      `json:"id" gorm:"primaryKey"`
	Email         string   `json:"email" gorm:"unique;size:255"`
	EmailVerified bool     `json:"email_verified" gorm:"default:false"`
	Password      string   `json:"password" gorm:"not null;size:255"`
	Status        bool     `json:"status" gorm:"default:false"`
	UserTypeID    int      `json:"ut_id"`
	UserType      UserType `gorm:"foreignKey:UserTypeID; references:ID ;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (u *User) SaveUser() error {
	helpers.DB.AutoMigrate(&User{})
	err := helpers.DB.Create(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) BeforeSave(tx *gorm.DB) error {
	hashedPassword, err := helpers.HashPassword(u.Password)
	if err != nil {
		return err
	}
	u.Password = string(hashedPassword)
	return nil
}

func (u *User) GetUserWithTypeById(id string) error {
	helpers.DB.AutoMigrate(&User{})
	// err := helpers.DB.Model(User{}).Joins("JOIN user_types ON user_types.id=users.user_type_id").Where("users.id = ?", id).First(&u).Error
	err := helpers.DB.Model(User{}).Preload("UserType").Where("id = ?", id).First(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllUsers() ([]User, error) {
	var users []User
	helpers.DB.AutoMigrate(&User{})
	err := helpers.DB.Model(User{}).Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (u *User) DeleteUser() error {
	helpers.DB.AutoMigrate(&User{})
	err := helpers.DB.Delete(&u).Error
	if err != nil {
		return err
	}
	return nil
}

func (u *User) UpdateUser(ui User) error {
	helpers.DB.AutoMigrate(&Student{})
	err := helpers.DB.Model(&u).Updates(ui).Error
	// fmt.Println(s)
	if err != nil {
		return err
	}
	return nil
}
