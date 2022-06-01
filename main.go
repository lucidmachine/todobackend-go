package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/go-chi/chi/v5"
)

const dbFilename string = "todo.db"

func main() {
	db, err := sql.Open("sqlite3", dbFilename)
	if err != nil {
		log.Panicf("Failed to open database file %s: %v", dbFilename, err)
	}

	repo := NewRepository(db)
	controller := NewController(repo)

	router := chi.NewRouter()
	registerMiddleware(router)
	registerRoutes(router, controller)

	http.ListenAndServe(":8080", router)
}
