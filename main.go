package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func getTodos(w http.ResponseWriter, r *http.Request) {
	todos := []Todo{}
	json.NewEncoder(w).Encode(todos)
}

func createTodo(w http.ResponseWriter, r *http.Request) {
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	w.WriteHeader(201)
	json.NewEncoder(w).Encode(todo)
}

func deleteTodo(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(204)
}

func main() {
	r := chi.NewRouter()

	registerMiddleware(r)
	registerRoutes(r)

	http.ListenAndServe(":8080", r)
}
