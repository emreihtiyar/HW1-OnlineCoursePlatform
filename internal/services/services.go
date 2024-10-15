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
func getId(w http.ResponseWriter, r *http.Request) string {
	// Extracts the variables from the request
	vars := mux.Vars(r)
	id := vars["id"] // Get the ID from the URL
	return id
}

func GetStudentById(w http.ResponseWriter, r *http.Request) {
	id := getId(w, r)
	entity, err := handlers.GetById(id, "Student")
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
