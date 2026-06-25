package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"users-api-memory/internal/models"
	"users-api-memory/internal/response"
)

var users []models.User

func AddUserHandle(w http.ResponseWriter, r *http.Request) {
	var user models.User
	if r.Method != http.MethodPost {
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid JSON!")
		return
	}

	if user.Name == "" {
		response.WriteError(w, http.StatusBadRequest, "Name is empty!")
		return
	}

	if user.Age <= 0 {
		response.WriteError(w, http.StatusBadRequest, "Age must be positive!")
		return
	}

	for _, userExist := range users {
		if userExist.Name == user.Name {
			response.WriteError(w, http.StatusConflict, "User already exist!")
			return
		}
	}

	users = append(users, user)
	response.WriteSuccess(w, http.StatusCreated, "User added successfully!", user)
}

func GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}
	jsonData, err := json.MarshalIndent(users, "", "  ")
	if err != nil {
		response.WriteError(w, http.StatusBadRequest, "Invalid JSON!")
		return
	}
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, string(jsonData))
}

func GetUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		response.WriteError(w, http.StatusBadRequest, "Name is empty!")
		return
	}

	for _, user := range users {
		if user.Name == name {
			jsonData, err := json.MarshalIndent(user, "", "  ")
			if err != nil {
				response.WriteError(w, http.StatusBadRequest, "Invalid JSON!")
				return
			}
			w.Header().Set("Content-Type", "application/json")
			fmt.Fprintln(w, string(jsonData))
			return
		}
	}
	response.WriteError(w, http.StatusBadRequest, "User not found!")
}

func DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		response.WriteError(w, http.StatusBadRequest, "Name is empty!")
		return
	}
	for i, user := range users {
		if user.Name == name {
			users = append(users[:i], users[i+1:]...)
			fmt.Fprintln(w, "Student deleted!")
			return
		}
	}
	response.WriteError(w, http.StatusBadRequest, "User not found!")
	return
}

func GetCountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}
	fmt.Fprintln(w, "Count of users:", len(users))
}

func ClearUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	users = []models.User{}
	fmt.Fprintln(w, "All Users deleted!")

}
