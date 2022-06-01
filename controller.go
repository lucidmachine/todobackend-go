package main

import (
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
)

type Controller struct {
	repo    *Repository
	baseUrl string
}

func NewController(repo *Repository, baseUrl string) Controller {

	return Controller{repo: repo, baseUrl: baseUrl}
}

func (c Controller) CreateTodo(w http.ResponseWriter, r *http.Request) {
	// Deserialize
	var todo Todo
	err := json.NewDecoder(r.Body).Decode(&todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Assign ID
	id := uuid.New()
	todo.Id = id
	todo.Url = c.baseUrl + id.String()

	// Persist
	err = c.repo.CreateTodo(todo)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Respond
	w.WriteHeader(201)
	json.NewEncoder(w).Encode(todo)
}

func (c Controller) GetTodos(w http.ResponseWriter, r *http.Request) {
	// Retrieve
	todos, err := c.repo.GetTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Respond
	json.NewEncoder(w).Encode(todos)
}

func (c Controller) GetTodoById(w http.ResponseWriter, r *http.Request) {
	// Read
	idStr := chi.URLParam(r, "id")
	id, err := uuid.Parse(idStr)
	if err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
	}

	// Retrieve
	todo, err := c.repo.GetTodo(id)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Respond
	json.NewEncoder(w).Encode(todo)
}

func (c Controller) DeleteTodos(w http.ResponseWriter, r *http.Request) {
	// Delete
	_, err := c.repo.DeleteTodos()
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}

	// Respond
	w.WriteHeader(204)
	json.NewEncoder(w).Encode([]Todo{})
}
