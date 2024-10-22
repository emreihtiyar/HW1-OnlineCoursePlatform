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
	Email        string     `json:"email" gorm:"unique"`
	Password     string     `json:"password"`
	RegistryDate *time.Time `json:"registrydate,omitempty" gorm:"column:registrydate"`
}

// Instructor Struct for Instructor
type Instructor struct {
	ID           string     `gorm:"primaryKey"`
	Name         string     `json:"name"`
	Email        string     `json:"email" gorm:"unique"`
	Password     string     `json:"password"`
	RegistryDate *time.Time `json:"registrydate,omitempty" gorm:"column:registrydate"`
	Expertise    string     `json:"expertise" gorm:"column:areaofexpertise"`
}

// Course Struct for Course
type Course struct {
	ID           string     `gorm:"primaryKey"`
	Title        string     `json:"title"`
	Description  string     `json:"description"`
	Date         *time.Time `json:"registrydate,omitempty" gorm:"column:coursedate"`
	InstructorID string     `gorm:"column:instructorid"`
}

// Enrollment struct for Enrollment table. This struct is used to represent a row in the Enrollment table
type Enrollment struct {
	ID             string     `gorm:"primaryKey"`
	StudentID      string     `json:"student_id" gorm:"column:studentid"`
	CourseID       string     `json:"course_id" gorm:"column:courseid"`
	EnrollmentDate *time.Time `json:"registrydate,omitempty" gorm:"column:enrollmentdate"`
	Grade          int        `json:"grade"`
}

type InstructorCourses struct {
	Instructor Instructor `json:"instructor"`
	Courses    []Course   `json:"courses"`
}

// Neo4jStudent represents the student node in Neo4j
type Neo4jStudent struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	RegistryDate *time.Time `json:"registryDate"`
}

// Neo4jInstructor represents the instructor node in Neo4j
type Neo4jInstructor struct {
	ID           string     `json:"id"`
	Name         string     `json:"name"`
	Email        string     `json:"email"`
	Password     string     `json:"password"`
	RegistryDate *time.Time `json:"registryDate"`
	Expertise    string     `json:"expertise"`
}

// Neo4jCourse represents the course node in Neo4j
type Neo4jCourse struct {
	ID          string     `json:"id"`
	Title       string     `json:"title"`
	Description string     `json:"description"`
	Date        *time.Time `json:"date"`
}

// Neo4jEnrollment represents the enrollment relationship in Neo4j
type Neo4jEnrollment struct {
	ID             string     `json:"id"`
	StudentID      string     `json:"studentId"`
	CourseID       string     `json:"courseId"`
	EnrollmentDate *time.Time `json:"enrollmentDate"`
	Grade          int        `json:"grade"`
}
