package main

import (
	"fmt"
	"log"
	"net/http"
	"users-api-postgres/internal/database"
	"users-api-postgres/internal/handlers"
)

func main() {
	db, err := database.Connect()
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()
	fmt.Println("Connected to database successfully!")

	userHandler := handlers.NewUserHandler(db)

	http.HandleFunc("/addUser", userHandler.AddUserHandle)
	http.HandleFunc("/getUsers", userHandler.GetUsersHandler)
	http.HandleFunc("/getUser", userHandler.GetUserHandler)
	http.HandleFunc("/getCount", userHandler.GetCountHandler)
	http.HandleFunc("/deleteUser", userHandler.DeleteUserHandler)
	http.HandleFunc("/clearUsers", userHandler.ClearUsersHandler)
	http.HandleFunc("/updateUser", userHandler.UpdateUser)

	http.ListenAndServe(":8080", nil)

}
