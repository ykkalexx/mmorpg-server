package handlers

import (
	"encoding/json"
	"net/http"
)

type apiFunction func(http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func WriteJSON(w http.ResponseWriter, status int, value any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(value)
}

func makeHTTPHandleFunc(fn apiFunction) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if err := fn(w, r); err != nil {
			WriteJSON(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}