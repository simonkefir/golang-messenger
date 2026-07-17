package core_http_response

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Error   string `json:"error" example:"full error text"`
	Message string `json:"message" example:"short human-readable message"`
}

type ValidationErrorResponse struct {
	Errors map[string]string `json:"errors"`
}

func JSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func Error(w http.ResponseWriter, status int, msg string) {
	JSON(w, status, ErrorResponse{Error: msg})
}

func ValidationError(w http.ResponseWriter, errs map[string]string) {
	JSON(w, http.StatusUnprocessableEntity, ValidationErrorResponse{Errors: errs})
}
