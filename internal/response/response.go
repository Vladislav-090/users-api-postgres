package response

import (
	"encoding/json"
	"net/http"
	"users-api-postgres/internal/models"
)

type ErrorResponse struct {
	Error string `json:"error"`
}

type SuccessResponse struct {
	Message string      `json:"message"`
	User    models.User `json:"user"`
}

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	errorResponse := ErrorResponse{
		Error: message,
	}
	WriteJSON(w, status, errorResponse)
}

func WriteSuccess(w http.ResponseWriter, status int, message string, user models.User) {
	successResponse := SuccessResponse{
		Message: message,
		User:    user,
	}
	WriteJSON(w, status, successResponse)
}
