package main

import (
	"fmt"
	"log"
	"net/http"
	"users-api-memory/internal/database"
	"users-api-memory/internal/handlers"
)

func main() {
	db, err := database.Connect()
		if err != nil{
			log.Fatal(err)
		}
	defer db.Close()
	fmt.Println("Connected to database successfully!")
	
	userHandler := handlers.NewUserHandler(db)



	http.HandleFunc("/addUser", userHandler.AddUserHandle )
	http.HandleFunc("/getUsers", userHandler.GetUsersHandler)
	http.HandleFunc("/getUser", userHandler.GetUserHandler)
	http.HandleFunc("/getCount", handlers.GetCountHandler)
	http.HandleFunc("/deleteUser", handlers.DeleteUserHandler)
	http.HandleFunc("/clearUsers", handlers.ClearUsersHandler)

	http.ListenAndServe(":8080", nil)

}
