package main

import (
	"github.com/google/uuid"
)

type Todo struct {
	Id        uuid.UUID `json:"id"`
	Title     string    `json:"title"`
	Completed bool      `json:"completed"`
	Order     int       `json:"order"`
	Url       string    `json:"url"`
}

type UpdateTodoRequest struct {
	Title     *string `json:"title,omitempty"`
	Completed *bool   `json:"completed,omitempty"`
	Order     *int    `json:"order,omitempty"`
}
