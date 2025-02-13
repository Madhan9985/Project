package main

import (
	"fmt"
	"net/http"

	"github.com/Madhan9985/Project/database"
	"github.com/Madhan9985/Project/handlers"
	"github.com/gorilla/mux"
)

func main() {
	database.InitDB()

	r := mux.NewRouter()
	r.HandleFunc("/todos", handlers.GetTodos).Methods("GET")
	r.HandleFunc("/todos", handlers.CreateTodo).Methods("POST")
	r.HandleFunc("/todos/{id}", handlers.UpdateTodo).Methods("PUT")
	r.HandleFunc("/todos/{id}", handlers.DeleteTodo).Methods("DELETE")

	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", r)
}
