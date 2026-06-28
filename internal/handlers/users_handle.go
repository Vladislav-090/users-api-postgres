package handlers

import (
	"database/sql"
	"encoding/json"
	"log"
	"net/http"
	"users-api-memory/internal/models"
	"users-api-memory/internal/response"
)

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

	users := make([]models.User, 0)
	
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

func (h *UserHandler) DeleteUserHandler(w http.ResponseWriter, r *http.Request) {

	if r.Method != http.MethodDelete {
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	name := r.URL.Query().Get("name")
	if name == "" {
		response.WriteError(w, http.StatusBadRequest, "Name is empty!")
		return
	}

	query := `
	DELETE FROM users
	WHERE name = $1
	`
	result, err := h.DB.Exec(query,name)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to delete user!")
		return
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "")
		return
	}
	if rowsAffected == 0 {
		response.WriteError(w, http.StatusNotFound, "User not found!")
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]string{
		"message" : "User deleted successfully",
		"name" : name,
	})
}

func (h *UserHandler) GetCountHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	var count int

	query := `
	SELECT COUNT(*) AS count_of_users
	FROM users`

	err := h.DB.QueryRow(query).Scan(&count)
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Faild to get count of users!")
		return
	}

	response.WriteJSON(w, http.StatusOK, map[string]int{
		"count_of_users:" : count,
	})
}

func (h *UserHandler) ClearUsersHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodDelete {
		response.WriteError(w, http.StatusMethodNotAllowed, "Method not allowed!")
		return
	}

	query := `
	DELETE FROM users
	`
	result, err := h.DB.Exec(query)
		if err != nil {
			response.WriteError(w, http.StatusInternalServerError, "Failed to delete all users")
			return
		} 
	
	rawAffected, err := result.RowsAffected()
	if err != nil {
		response.WriteError(w, http.StatusInternalServerError, "Failed to check deleted rows!")
		return
	}

	response.WriteJSON(w , http.StatusOK, map[string]interface{}{
		"message:" : "All users deleted successfully!",
		"deleted_count:" : rawAffected,
	} )

}
