package models

import "github.com/khatrisaugat/PatternPractise/app/helpers"

type EnrolledCourse struct {
	ID        int     `json:"id" gorm:"primaryKey"`
	StudentID int     `json:"student_id"`
	CourseID  int     `json:"course_id"`
	Student   Student `gorm:"foreignKey:StudentID; references:ID ;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
	Course    Course  `gorm:"foreignKey:CourseID; references:ID ;constraint:OnUpdate:CASCADE,OnDelete:SET NULL;"`
}

func (ec *EnrolledCourse) SaveEnrolledCourse() error {
	helpers.DB.AutoMigrate(&EnrolledCourse{})
	err := helpers.DB.Create(&ec).Error
	if err != nil {
		return err
	}
	return nil
}

// func (ec *EnrolledCourse) UpdateEnrolledCourse(ui EnrolledCourse) error {
// 	helpers.DB.AutoMigrate(&EnrolledCourse{})
// 	err := helpers.DB.Model(&ec).Updates(ui).Error
// 	// fmt.Println(s)
// 	if err != nil {
// 		return err
// 	}
// 	return nil
// }

func (ec *EnrolledCourse) DeleteEnrolledCourse() error {
	helpers.DB.AutoMigrate(&EnrolledCourse{})
	err := helpers.DB.Delete(&ec).Error
	if err != nil {
		return err
	}
	return nil
}

func (ec *EnrolledCourse) GetEnrolledCourse(id string) error {
	helpers.DB.AutoMigrate(&EnrolledCourse{})
	err := helpers.DB.Model(EnrolledCourse{}).Where("id = ?", id).First(&ec).Error
	if err != nil {
		return err
	}
	return nil
}

func GetAllEnrolledCoursesForStudent(sid string) ([]EnrolledCourse, error) {
	var enrolledCourses []EnrolledCourse
	helpers.DB.AutoMigrate(&EnrolledCourse{})
	err := helpers.DB.Model(EnrolledCourse{}).Where("sid=?", sid).Find(&enrolledCourses).Error
	if err != nil {
		return nil, err
	}
	return enrolledCourses, nil
}
