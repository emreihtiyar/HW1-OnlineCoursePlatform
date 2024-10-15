package handlers

import (
	"HW1-OnlineCoursePlatform/internal/models"
	"fmt"
	"gorm.io/gorm"
)

var db *gorm.DB

// TODO: Don't know if this is okay
func InitPostgresHandler(dbInit *gorm.DB) {
	db = dbInit
}

// GetById returns the entity with the given ID.
func GetById(Id string, entityName string) (interface{}, error) {
	switch entityName {
	case "Student":
		var student models.Student
		result := db.Where("id = ?", Id).Find(&student)
		if result.RowsAffected > 0 {
			return student, nil
		}
		return nil, nil
	case "Course":
		var course models.Course
		result := db.Where("id = ?", Id).Find(&course)
		if result.RowsAffected > 0 {
			return course, nil
		}
		return nil, nil
	case "Instructor":
		var instructor models.Instructor
		result := db.Where("id = ?", Id).Find(&instructor)
		if result.RowsAffected > 0 {
			return instructor, nil
		}
		return nil, nil
	case "Enrollment":
		var enrollment models.Enrollment
		result := db.Where("id = ?", Id).Find(&enrollment)
		if result.RowsAffected > 0 {
			return enrollment, nil
		}
		return nil, nil
	}
	return nil, fmt.Errorf("Entity not found")
}

// Insert function to add a new record to the database
func Insert(entity interface{}) error {
	result := db.Create(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}
