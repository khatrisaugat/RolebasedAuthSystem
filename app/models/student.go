package models

import (
	"github.com/khatrisaugat/PatternPractise/app/helpers"
)

type Student struct {
	ID           int    `json:"id" gorm:"primaryKey"`
	FirstName    string `json:"first_name" gorm:"size:255"`
	LastName     string `json:"last_name" gorm:"size:255"`
	Address      string `json:"address" gorm:"size:255"`
	Contact      string `json:"contact" gorm:"size:255"`
	GuardianName string `json:"guardian_name" gorm:"size:255"`
	Dob          string `json:"dob" gorm:"type:date"`
	UserId       int    `json:"user_id"`
	User         User   `gorm:"foreignKey:UserId; references:ID ;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (s *Student) SaveStudent() error {
	helpers.DB.AutoMigrate(&Student{})
	err := helpers.DB.Create(&s).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Student) UpdateStudent(ui Student) error {
	helpers.DB.AutoMigrate(&Student{})
	err := helpers.DB.Model(&s).Updates(ui).Error
	// fmt.Println(s)
	if err != nil {
		return err
	}
	return nil
}

func (s *Student) DeleteStudent() error {
	helpers.DB.AutoMigrate(&Student{})
	err := helpers.DB.Delete(&s).Error
	if err != nil {
		return err
	}
	return nil
}

func (s *Student) GetStudent(id string) error {
	helpers.DB.AutoMigrate(&Student{})
	err := helpers.DB.Model(Student{}).Where("id = ?", id).First(&s).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllStudents() ([]Student, error) {
	var students []Student
	helpers.DB.AutoMigrate(&Student{})
	err := helpers.DB.Find(&students).Error
	if err != nil {
		return nil, err
	}
	return students, nil
}
