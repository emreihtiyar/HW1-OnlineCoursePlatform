package services

import (
	"encoding/json"
	"fmt"
	"github.com/gorilla/mux"
	"internal/handlers"
	"internal/models"
	"net/http"
)

// encode
func encode[T any](w http.ResponseWriter, r *http.Request, status int, v T) error {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json: %w", err)
	}
	return nil
}

// decode
func decode[T any](r *http.Request) (T, error) {
	var v T
	if err := json.NewDecoder(r.Body).Decode(&v); err != nil {
		return v, fmt.Errorf("decode json: %w", err)
	}
	return v, nil
}

// handleParameters returns the given parameter.
func handleParameters(r *http.Request, parameter string) string {
	vars := mux.Vars(r)
	param := vars[parameter]
	return param
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	if !isRequestMethodValid(w, r, http.MethodGet) {
		return
	}
	request := parseRequest(w, readBody(w, r))
	// TODO: This conversion can be achieved with pointers probably.
	handlers.Get(models.Login{
		Username: request.Username,
		Password: request.Password,
	})

}

func AddUser(w http.ResponseWriter, r *http.Request) {
	if !isRequestMethodValid(w, r, http.MethodPost) {
		return
	}
	request := parseRequest(w, readBody(w, r))
	// TODO: This does not insert.
	handlers.Insert(models.Login{
		Username: request.Username,
		Password: request.Password,
	})
}
