package handlers

import (
	"HW1-OnlineCoursePlatform/internal/models"
	"context"
	"fmt"
	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
	"gorm.io/gorm"
	"log"
	"time"
)

var db *gorm.DB
var neo4jDriver neo4j.DriverWithContext

func CloseDriver() {
	if neo4jDriver != nil {
		err := neo4jDriver.Close(context.Background())
		if err != nil {
			return
		}
	}
}

func withSession(fn func(neo4j.SessionWithContext) (interface{}, error)) (interface{}, error) {
	session := neo4jDriver.NewSession(context.Background(), neo4j.SessionConfig{})
	defer session.Close(context.Background())

	return fn(session)
}

// TODO: Don't know if this is okay
// Will change the name of this function and the parameter.
// If dbInit is nil, then it will initialize neo4j driver.
// Also, GetById will include a if check based on the driver type.
func InitDBHandler(dbInit *gorm.DB, neo4jDriverInit neo4j.DriverWithContext) {
	db = dbInit
	neo4jDriver = neo4jDriverInit
}

// GetById returns the entity with the given ID.
func GetById(Id string, entityName string) (interface{}, error) {
	if db != nil {
		return getPostgresById(Id, entityName)
	}
	return getNeo4jById(Id, entityName)
}

// getPostgresById returns the entity with the given ID from the Postgres database.
func getPostgresById(Id string, entityName string) (interface{}, error) {
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

// GetStudentByID retrieves a Student from Neo4j based on their ID
func GetStudentByID(id string) (*models.Student, error) {
	var student *models.Student

	query := `
		MATCH (s:Student {ID: $id})
		RETURN s {.*} AS student
	`

	_, err := withSession(func(session neo4j.SessionWithContext) (interface{}, error) {
		return session.ExecuteRead(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
			result, err := tx.Run(context.Background(), query, map[string]interface{}{"id": id})
			if err != nil {
				return nil, err
			}

			if result.Next(context.Background()) {
				record := result.Record()
				studentMap, exists := record.Get("student")
				if !exists {
					return nil, fmt.Errorf("student with ID %s not found", id)
				}

				// studentMap is a map[string]interface{} containing node properties
				props := studentMap.(map[string]interface{})

				// Convert RegistryDate to *time.Time if it's a dbtype.Date or dbtype.LocalDateTime
				var registryDate *time.Time

				// Parse the string into a time.Time object
				layout := time.RFC3339
				parsedTime, err := time.Parse(layout, props["RegistryDate"].(string))
				if err != nil {
					return nil, fmt.Errorf("Error parsing time:", err)
				}

				// Convert to *time.Time
				registryDate = &parsedTime

				// Populate the student struct
				student = &models.Student{
					ID:           props["ID"].(string),
					Name:         props["Name"].(string),
					Email:        props["Email"].(string),
					Password:     props["Password"].(string),
					RegistryDate: registryDate,
				}

				return student, nil
			}

			return nil, fmt.Errorf("student with ID %s not found", id)
		})
	})

	if err != nil {
		log.Printf("Error fetching student by ID: %v", err)
		return nil, err
	}

	return student, nil
}

// GetInstuctorById retrieves an Instructor from Neo4j based on their ID
func GetInstructorById(id string) (*models.Instructor, error) {
	var insturctor *models.Instructor

	query := `
		MATCH (i:Instructor {ID: $id})
		RETURN i {.*} AS instructor
	`

	_, err := withSession(func(session neo4j.SessionWithContext) (interface{}, error) {
		return session.ExecuteRead(context.Background(), func(tx neo4j.ManagedTransaction) (interface{}, error) {
			result, err := tx.Run(context.Background(), query, map[string]interface{}{"id": id})
			if err != nil {
				return nil, err
			}
			if result.Next(context.Background()) {
				record := result.Record()
				// TODO: Simplify this with combining them into single function.
				insturctorMap, exists := record.Get("insturctor")
				if !exists {
					return nil, fmt.Errorf("insturctor with ID %s not found", id)
				}

				// Use type assertion to extract the fields
				insturctorProps := insturctorMap.(neo4j.Node).Props

				insturctor = &models.Instructor{
					ID:           insturctorProps["ID"].(string),
					Name:         insturctorProps["Name"].(string),
					Email:        insturctorProps["Email"].(string),
					Password:     insturctorProps["Password"].(string),
					RegistryDate: (*time.Time)(insturctorProps["RegistryDate"].(*neo4j.Date)), // Adjust the type as necessary
					Expertise:    insturctorProps["Expertise"].(string),
				}
				return nil, nil
			}
			return nil, nil
		})
	})

	if err != nil {
		log.Printf("Error fetching student by ID: %v", err)
		return nil, err
	}

	return insturctor, nil
}

// GetStudentByID retrieves a Student from Neo4j based on their ID
func GetCourseById(id string) (*models.Student, error) {
	return nil, nil
}

// GetStudentByID retrieves a Student from Neo4j based on their ID
func GetEnrollmentById(id string) (*models.Student, error) {
	return nil, nil
}

func getNeo4jById(id string, entityName string) (interface{}, error) {
	// Determine the query based on entityName
	switch entityName {
	case "Student":
		return GetStudentByID(id)
	case "Instructor":
		return GetInstructorById(id)
	case "Course":
		return GetCourseById(id)
	case "Enrollment":
		return GetEnrollmentById(id)
	default:
		return nil, fmt.Errorf("unsupported entity: %s", entityName)
	}
}

func GetAll(entity interface{}) error {
	result := db.Find(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Insert function to add a new record to the database
func Insert(entity interface{}) error {
	if db != nil {
		return insertPostgres(entity)
	}
	return addStudentNeo4j(entity.(models.Student))
}

func insertPostgres(entity interface{}) error {
	result := db.Create(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func addStudentNeo4j(student models.Student) error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Convert *time.Time to a string if RegistryDate is not nil
	var registryDate string
	if student.RegistryDate != nil {
		registryDate = student.RegistryDate.Format(time.RFC3339) // Convert to ISO 8601 string
	}

	_, err := withSession(func(session neo4j.SessionWithContext) (interface{}, error) {
		query := `
            CREATE (s:Student {ID: $id, Name: $name, Email: $email, Password: $password, RegistryDate: $registryDate})
            RETURN s`

		return session.ExecuteWrite(ctx, func(tx neo4j.ManagedTransaction) (interface{}, error) {
			// Run the query with parameters
			_, err := tx.Run(ctx, query, map[string]any{
				"id":           student.ID,
				"name":         student.Name,
				"email":        student.Email,
				"password":     student.Password,
				"registryDate": registryDate, // String date format passed here
			})

			return nil, err
		})
	})

	return err
}

// Update function to update a record in the database
func Update(entity interface{}) error {
	result := db.Save(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

// Delete function to delete a record from the database
func Delete(entity interface{}) error {
	result := db.Delete(entity)
	if result.Error != nil {
		return result.Error
	}
	return nil
}

func GetCoursesByInstructorId(instructorId string) ([]models.Course, error) {
	var courses []models.Course
	result := db.Where("instructorid = ?", instructorId).Find(&courses)
	if result.Error != nil {
		return nil, result.Error
	}
	return courses, nil
}
