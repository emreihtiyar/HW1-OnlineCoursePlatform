package models

import (
	"github.com/oklog/ulid"
	"math/rand"
	"time"
)

// GenerateULID generates a new ULID as a string
func GenerateULID() string {
	t := time.Now().UTC()
	entropy := ulid.Monotonic(rand.New(rand.NewSource(t.UnixNano())), 0)
	return ulid.MustNew(ulid.Timestamp(t), entropy).String()
}

// Student Struct for Student
type Student struct {
	ID           string     `gorm:"primaryKey"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	RegistryDate *time.Time `json:"registrydate,omitempty" gorm:"column:registrydate"`
}

// Instructor Struct for Instructor
type Instructor struct {
	ID           string     `gorm:"primaryKey"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	RegistryDate *time.Time `json:"registrydate,omitempty" gorm:"column:registrydate"`
	Expertise    string     `json:"expertise"`
}

// Course Struct for Course
type Course struct {
	ID           string     `gorm:"primaryKey"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Date         *time.Time `json:"registrydate,omitempty" gorm:"column:registrydate"`
	InstructorID string
}

// Enrollment struct for Enrollment table. This struct is used to represent a row in the Enrollment table
type Enrollment struct {
	ID             string     `gorm:"primaryKey"`
	StudentID      string     `json:"student_id"`
	CourseID       string     `json:"course_id"`
	EnrollmentDate *time.Time `json:"registrydate,omitempty" gorm:"column:registrydate"`
	Grade          int        `json:"grade"`
}
