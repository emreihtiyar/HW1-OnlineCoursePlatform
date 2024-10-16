package services

import (
	"HW1-OnlineCoursePlatform/internal/handlers"
	"HW1-OnlineCoursePlatform/internal/models"
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"time"
)

// encode writes the given parameter to the response writer.
func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

// decode returns the given parameter.
func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

// GetURIParam function to extract a parameter from the request's URI
func GetURIParam(r *http.Request, param string) string {
	return r.URL.Query().Get(param)
}

// getId function to extract the ID from the request's URL
func getId(w http.ResponseWriter, r *http.Request, parameter string) string {
	// Extracts the variables from the request
	vars := mux.Vars(r)
	id := vars[parameter] // Get the ID from the URL
	return id
}

func GetById(w http.ResponseWriter, r *http.Request, entityName string) {
	id := getId(w, r, "id")
	entity, err := handlers.GetById(id, entityName)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if entity == nil {
		http.Error(w, "Entity not found", http.StatusNotFound)
		return
	}
	encode(w, r, http.StatusOK, entity)
}

func GetStudents(w http.ResponseWriter, r *http.Request) {
	var students []models.Student
	err := handlers.GetAll(&students)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encode(w, r, http.StatusOK, students)
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	GetById(w, r, "Student")
}

func AddStudent(w http.ResponseWriter, r *http.Request) {
	// Decode JSON body into a Student object
	student, err := decode[models.Student](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	student.ID = models.GenerateULID()
	if student.RegistryDate == nil {
		now := time.Now()           // Get the current time
		student.RegistryDate = &now // Assign the address of now to registryDate
	}
	err = handlers.Insert(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responding with the inserted student data or ID
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Inserted student with ID: %s\n", student.ID)
}

func UpdateStudent(w http.ResponseWriter, r *http.Request) {
	student, err := decode[models.Student](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	student.ID = getId(w, r, "id")
	err = handlers.Update(student)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responding with the updated student data or ID
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated student with ID: %s\n", student.ID)
}

func DeleteStudent(w http.ResponseWriter, r *http.Request) {
	id := getId(w, r, "id")
	err := handlers.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responding with the deleted student ID
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted student with ID: %s\n", id)
}

func GetInstructors(w http.ResponseWriter, r *http.Request) {
	var instructors []models.Instructor
	err := handlers.GetAll(&instructors)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encode(w, r, http.StatusOK, instructors)
}

func GetInstructorById(w http.ResponseWriter, r *http.Request) {
	GetById(w, r, "Instructor")
}

func GetInstructorCourses(w http.ResponseWriter, r *http.Request) {
	id := getId(w, r, "id")
	courses, errCourses := handlers.GetCoursesByInstructorId(id)
	if errCourses != nil {
		http.Error(w, errCourses.Error(), http.StatusInternalServerError)
		return
	}

	instructor, err := handlers.GetById(id, "Instructor")
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	if instructor == nil {
		http.Error(w, "Instructor not found", http.StatusNotFound)
		return
	}

	var instructorCourses models.InstructorCourses
	instructorCourses.Instructor = instructor.(models.Instructor)
	instructorCourses.Courses = courses
	encode(w, r, http.StatusOK, instructorCourses)
}

func AddInstructor(w http.ResponseWriter, r *http.Request) {
	// Decode JSON body into a Student object
	instructor, err := decode[models.Instructor](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	instructor.ID = models.GenerateULID()
	if instructor.RegistryDate == nil {
		now := time.Now()              // Get the current time
		instructor.RegistryDate = &now // Assign the address of now to registryDate
	}
	err = handlers.Insert(instructor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responding with the inserted student data or ID
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Inserted instructor with ID: %s\n", instructor.ID)
}

func UpdateInstructor(w http.ResponseWriter, r *http.Request) {
	instructor, err := decode[models.Instructor](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	instructor.ID = getId(w, r, "id")
	err = handlers.Update(instructor)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responding with the updated student data or ID
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated instructor with ID: %s\n", instructor.ID)
}

func DeleteInstructor(w http.ResponseWriter, r *http.Request) {
	id := getId(w, r, "id")
	err := handlers.Delete(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responding with the deleted student ID
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Deleted instructor with ID: %s\n", id)
}

func GetCourseById(w http.ResponseWriter, r *http.Request) {
	GetById(w, r, "Course")
}

func GetCourses(w http.ResponseWriter, r *http.Request) {
	var courses []models.Course
	err := handlers.GetAll(&courses)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encode(w, r, http.StatusOK, courses)
}

func AddCourse(w http.ResponseWriter, r *http.Request) {
	course, err := decode[models.Course](r)
	instructorId := getId(w, r, "id")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	course.ID = models.GenerateULID()
	course.InstructorID = instructorId
	if course.Date == nil {
		now := time.Now()  // Get the current time
		course.Date = &now // Assign the address of now to registryDate
	}
	err = handlers.Insert(course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responding with the inserted student data or ID
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Inserted course with ID: %s\n", course.ID)
}

func UpdateCourse(w http.ResponseWriter, r *http.Request) {
	course, err := decode[models.Course](r)
	courseId := getId(w, r, "courseId")

	if courseId != "" {
		instructorId := getId(w, r, "instructorId")
		course.InstructorID = instructorId
		course.ID = courseId
	} else {
		course.ID = getId(w, r, "id")
	}

	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	err = handlers.Update(course)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responding with the updated student data or ID
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated course with ID: %s\n", course.ID)
}

func GetEnrollmentById(w http.ResponseWriter, r *http.Request) {
	GetById(w, r, "Enrollment")
}

func GetEnrollments(w http.ResponseWriter, r *http.Request) {
	var enrollments []models.Enrollment
	err := handlers.GetAll(&enrollments)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	encode(w, r, http.StatusOK, enrollments)
}

func AddEnrollment(w http.ResponseWriter, r *http.Request) {
	enrollment, err := decode[models.Enrollment](r)
	studentId := getId(w, r, "studentId")
	courseId := getId(w, r, "courseId")
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	enrollment.ID = models.GenerateULID()
	enrollment.StudentID = studentId
	enrollment.CourseID = courseId
	if enrollment.EnrollmentDate == nil {
		now := time.Now()                // Get the current time
		enrollment.EnrollmentDate = &now // Assign the address of now to registryDate
	}
	err = handlers.Insert(enrollment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responding with the inserted student data or ID
	w.WriteHeader(http.StatusCreated)
	fmt.Fprintf(w, "Inserted enrollment with ID: %s\n", enrollment.ID)
}

func UpdateEnrollment(w http.ResponseWriter, r *http.Request) {
	enrollment, err := decode[models.Enrollment](r)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	enrollment.ID = getId(w, r, "id")
	err = handlers.Update(enrollment)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	// Responding with the updated student data or ID
	w.WriteHeader(http.StatusOK)
	fmt.Fprintf(w, "Updated enrollment with ID: %s\n", enrollment.ID)
}
