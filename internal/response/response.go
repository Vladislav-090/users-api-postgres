package response

import (
	"encoding/json"
	"net/http"
	"users-api-memory/internal/models"
)

func WriteJSON(w http.ResponseWriter, status int, data any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

func WriteError(w http.ResponseWriter, status int, message string) {
	errorResponse := models.ErrorResponse{
		Error: message,
	}
	WriteJSON(w, status, errorResponse)
}

func WriteSuccess(w http.ResponseWriter, status int, message string, user models.User) {
	successResponse := models.SuccessResponse{
		Message: "User added successfully!",
		User:    models.User{},
	}
	WriteJSON(w, status, successResponse)
}
