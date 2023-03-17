package models

import "github.com/khatrisaugat/PatternPractise/app/helpers"

type Course struct {
	ID                int     `json:"id" gorm:"primaryKey"`
	Title             string  `json:"title" gorm:"size:255"`
	CreditHour        float64 `json:"credit_hour"`
	CourseDescription string  `json:"course_description" gorm:"size:255"`
}

func (c *Course) SaveCourse() error {
	helpers.DB.AutoMigrate(&Course{})
	err := helpers.DB.Create(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *Course) UpdateCourse(ui Course) error {
	helpers.DB.AutoMigrate(&Course{})
	err := helpers.DB.Model(&c).Updates(ui).Error
	// fmt.Println(s)
	if err != nil {
		return err
	}
	return nil
}

func (c *Course) DeleteCourse() error {
	helpers.DB.AutoMigrate(&Course{})
	err := helpers.DB.Delete(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func (c *Course) GetCourse(id string) error {
	helpers.DB.AutoMigrate(&Course{})
	err := helpers.DB.Model(Course{}).Where("id = ?", id).First(&c).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllCourses() ([]Course, error) {
	var courses []Course
	helpers.DB.AutoMigrate(&Course{})
	err := helpers.DB.Model(Course{}).Find(&courses).Error
	if err != nil {
		return nil, err
	}
	return courses, nil
}
