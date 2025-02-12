package main

import (
	"fmt"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	"github.com/yourusername/Project/database"
	"github.com/yourusername/Project/routes"
)

func main() {
	// Initialize the database
	database.InitDB()

	// Create a new router
	r := mux.NewRouter()

	// Register routes
	routes.RegisterTaskRoutes(r)

	// Start the server
	fmt.Println("Server started on :8080")
	log.Fatal(http.ListenAndServe(":8080", r))
}
