package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"users-api-memory/internal/models"
	"users-api-memory/internal/response"
)

var users []models.User

type UserHandler struct {
	DB *sql.DB
}

func NewUserHandler(db *sql.DB) *UserHandler{
	return &UserHandler{
		DB: db,
	}
}

func (h *UserHandler) AddUserHandle(w http.ResponseWriter, r *http.Request) {
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

	if user.Age < 0 {
		response.WriteError(w, http.StatusBadRequest, "Age must be positive!")
		return
	}
	
	if user.Email == "" {
		response.WriteError(w, http.StatusBadRequest, "Email is empty!")
		return
	}
	
	query := `
	INSERT INTO users(name, age, email)
	VALUES ($1, $2, $3)
	RETURNING id, is_active, created_at
	`

	err = h.DB.QueryRow(query, user.Name, user.Age, user.Email).Scan(
		&user.ID,
		&user.IsActive,
		&user.CreatedAt,
	)
	if err != nil {
		log.Println("Failed to create user!",err)
		response.WriteError(w, http.StatusInternalServerError, "Failed to create user!")
		return
	}

	response.WriteSuccess(w, http.StatusCreated, "User added successfully!", user)

}

func (h *UserHandler)GetUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}
	
	query := `
	SELECT id, name, age, email, is_active, created_at
	FROM users
	ORDER BY id ASC`

	rows, err := h.DB.Query(query)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to get users!")
		return
	}
	defer rows.Close()

	var users []models.User
	
	for rows.Next() {
		var user models.User

		err := rows.Scan(
			&user.ID,
			&user.Name,
			&user.Age,
			&user.Email,
			&user.IsActive,
			&user.CreatedAt,
		)
		if err != nil {
			response.WriteError(w, http.StatusInternalServerError, "Failed to scan user!")
			return
		}
		users = append(users, user)
	}
		response.WriteJSON(w, http.StatusOK, users)
}

func (h *UserHandler)GetUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	var user models.User

	name := r.URL.Query().Get("name")
	if name == ""{
		response.WriteError(w, http.StatusBadRequest, "Name is empty!")
		return
	}

	query:= `
	SELECT  id, name, age, email, is_active, created_at
	FROM users
	WHERE name = $1
	`

	err := h.DB.QueryRow(query,name).Scan(
		&user.ID,
		&user.Name,
		&user.Age,
		&user.Email,
		&user.IsActive,
		&user.CreatedAt,
	)
	if err != nil {
		response.WriteError(w, http.StatusNotFound, "User not found!")
		return
	}
	response.WriteJSON(w, http.StatusOK, user)
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
