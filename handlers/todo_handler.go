package handlers

import (
	"Project/database"
	"Project/models"
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
)

func GetTodos(w http.ResponseWriter, r *http.Request) {
	var todos []models.Todo
	database.DB.Find(&todos)
	json.NewEncoder(w).Encode(todos)
}

func CreateTodo(w http.ResponseWriter, r *http.Request) {
	var todo models.Todo
	json.NewDecoder(r.Body).Decode(&todo)
	database.DB.Create(&todo)
	json.NewEncoder(w).Encode(todo)
}

func UpdateTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todo models.Todo
	database.DB.First(&todo, vars["id"])
	json.NewDecoder(r.Body).Decode(&todo)
	database.DB.Save(&todo)
	json.NewEncoder(w).Encode(todo)
}

func DeleteTodo(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	var todo models.Todo
	database.DB.Delete(&todo, vars["id"])
	w.WriteHeader(http.StatusNoContent)
}
