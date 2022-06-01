package main

import (
	"encoding/json"
	"net/http"
)

type Controller struct {
	repo *Repository
}

func NewController(repo *Repository) Controller {
	return Controller{repo: repo}
}

func (c Controller) getTodos(w http.ResponseWriter, r *http.Request) {
	// Retrieve
	todos, err := c.repo.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Respond
	json.NewEncoder(w).Encode(todos)
}

func (c Controller) createTodo(w http.ResponseWriter, r *http.Request) {
	// Deserialize
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Persist
	id, err := c.repo.CreateTodo(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
	todo.Id = id

	// Respond
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(todo)
}

func (c Controller) deleteTodos(w http.ResponseWriter, r *http.Request) {
	// Delete
	_, err := c.repo.DeleteTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Respond
	w.WriteHeader(204)
	json.NewEncoder(w).Encode([]Todo{})
}
