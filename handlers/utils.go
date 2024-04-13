package handlers

import (
	"encoding/json"
	"net/http"
)

type apiFunction func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func WriteJSON(w http.ResponseWriter, statusCode int, data interface{}) error {
    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(statusCode)
    return json.NewEncoder(w).Encode(data)
}
func makeHTTPHandleFunc(fn apiFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}