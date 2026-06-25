package main

import (
	"net/http"
	"users-api-memory/internal/handlers"
)

func main() {

	http.HandleFunc("/addUser", handlers.AddUserHandle)
	http.HandleFunc("/getUsers", handlers.GetUsersHandler)
	http.HandleFunc("/getUser", handlers.GetUserHandler)
	http.HandleFunc("/getCount", handlers.GetCountHandler)
	http.HandleFunc("/deleteUser", handlers.DeleteUserHandler)
	http.HandleFunc("/clearUsers", handlers.ClearUsersHandler)

	http.ListenAndServe(":8080", nil)

}
