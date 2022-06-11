package main

import (
	"testing"

	"github.com/google/uuid"
)

func TestUpdate(t *testing.T) {
	id, _ := uuid.Parse("7b6202a9-c7a6-4adc-86c1-fe29e5a76049")
	todo := Todo{
		Id:        id,
		Title:     "Existing Todo",
		Completed: false,
		Order:     1,
		Url:       "http://example.com/7b6202a9-c7a6-4adc-86c1-fe29e5a76049",
	}

	updatedTitle := "Updated Todo"
	updatedCompleted := true
	updatedOrder := 3
	updateTodoRequest := UpdateTodoRequest{
		Title:     &updatedTitle,
		Completed: &updatedCompleted,
		Order:     &updatedOrder,
	}

	todo = Update(todo, updateTodoRequest)

	expectedTodo := Todo{
		Id:        id,
		Title:     updatedTitle,
		Completed: updatedCompleted,
		Order:     updatedOrder,
		Url:       "http://example.com/7b6202a9-c7a6-4adc-86c1-fe29e5a76049",
	}
	if todo != expectedTodo {
		t.Errorf("\nExpected: %v\nActual:   %v", expectedTodo, todo)
	}
}
