package web

import (
	"HW1-OnlineCoursePlatform/internal/config"
	"HW1-OnlineCoursePlatform/internal/services"
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

const (
	STUDENTS                      = "/students"
	STUDENTS_ID                   = "/students/{id}"
	INSTRUCTORS                   = "/instructors"
	INSTRUCTORS_ID                = "/instructors/{id}"
	COURSES_GIVEN_BY_INSTRUCTOR   = "/courses/instructor/{id}"
	COURSES                       = "/courses"
	COURSES_ID                    = "/courses/{id}"
	COURSES_ID_AND_INSTRUCTOR     = "/courses/{courseId}/instructor/{instructorId}"
	ENROLLMENT                    = "/enrollments"
	ENROLLMENT_ID                 = "/enrollments/{id}"
	ENROLLMENT_STUDENT_AND_COURSE = "/enrollments/student/{studentId}/course/{courseId}"
)

func Routes(conf *config.Conf) {
	route := mux.NewRouter()

	// Students endpoints, Add Student, Get Student, Update Student, Delete Student
	route.HandleFunc(STUDENTS_ID, services.GetStudentById).Methods(http.MethodGet)
	route.HandleFunc(STUDENTS, services.GetStudents).Methods(http.MethodGet)
	route.HandleFunc(STUDENTS, services.AddStudent).Methods(http.MethodPost)
	route.HandleFunc(STUDENTS_ID, services.UpdateStudent).Methods(http.MethodPut)
	route.HandleFunc(STUDENTS_ID, services.DeleteStudent).Methods(http.MethodDelete)

	// Instructors endpoints, Add Instructor, Get Instructor, Update Instructor, Delete Instructor
	route.HandleFunc(INSTRUCTORS_ID, services.GetInstructorById).Methods(http.MethodGet)
	route.HandleFunc(INSTRUCTORS, services.GetInstructors).Methods(http.MethodGet)
	route.HandleFunc(COURSES_GIVEN_BY_INSTRUCTOR, services.GetInstructorCourses).Methods(http.MethodGet)
	route.HandleFunc(INSTRUCTORS, services.AddInstructor).Methods(http.MethodPost)
	route.HandleFunc(INSTRUCTORS_ID, services.UpdateInstructor).Methods(http.MethodPut)
	route.HandleFunc(INSTRUCTORS_ID, services.DeleteInstructor).Methods(http.MethodDelete)

	// Courses endpoints, Add Course, Get Course
	route.HandleFunc(COURSES_ID, services.GetCourseById).Methods(http.MethodGet)
	route.HandleFunc(COURSES, services.GetCourses).Methods(http.MethodGet)
	route.HandleFunc(COURSES_ID, services.AddCourse).Methods(http.MethodPost)
	route.HandleFunc(COURSES_ID, services.UpdateCourse).Methods(http.MethodPut)
	route.HandleFunc(COURSES_ID_AND_INSTRUCTOR, services.UpdateCourse).Methods(http.MethodPut)

	// Enrollment endpoints, Add Enrollment, Get Enrollment
	route.HandleFunc(ENROLLMENT_ID, services.GetEnrollmentById).Methods(http.MethodGet)
	route.HandleFunc(ENROLLMENT, services.GetEnrollments).Methods(http.MethodGet)
	route.HandleFunc(ENROLLMENT_STUDENT_AND_COURSE, services.AddEnrollment).Methods(http.MethodPost)
	// TODO: Update wouldn't work correctly, improve it.
	route.HandleFunc(ENROLLMENT_ID, services.UpdateEnrollment).Methods(http.MethodPut)

	// Start the server and handle any potential errors
	server := &http.Server{
		Addr:    conf.App.Port, // Address to listen on
		Handler: route,         // The router to handle incoming requests
	}

	// Log fatal error if http.ListenAndServe fails
	log.Fatal(server.ListenAndServe())
}
