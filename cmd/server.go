package main

import (
	"PresentationProject/config"
	"PresentationProject/internal/handler"
	"PresentationProject/internal/infra"
	"PresentationProject/internal/repository"
	"PresentationProject/migration"
	"fmt"

	"net/http"
)

func main() {
	// Initialize the database connection
	cfg := config.ReadConfigAndArg()
	db := infra.ConnectDb(cfg.DB)
	err := migration.Migration(db)
	if err != nil {
		fmt.Println(err)
		return
	}

	// Create an instance of UserRepository with the database connection
	userRepo := repository.NewUserRepository(db)

	// Create an instance of UserHandler with the UserRepository
	userHandler := handler.NewUserHandler(userRepo)

	// Set up the HTTP routes for each of the UserHandler methods
	http.HandleFunc("/user/create", userHandler.Create)
	http.HandleFunc("/user/find", userHandler.FindByEmail)
	http.HandleFunc("/user/update", userHandler.Update)
	http.HandleFunc("/user/delete", userHandler.Delete)

	// Start the HTTP server
	http.ListenAndServe(":8080", nil)
}
