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
	defer db.Close()

	repo := NewRepository(db)
	err = repo.DropTable()
	if err != nil {
		log.Panicf("Failed to drop table `todos`: %v", err)
	}
	err = repo.CreateTable()
	if err != nil {
		log.Panicf("Failed to create table `todos`: %v", err)
	}

	controller := NewController(repo)

	router := chi.NewRouter()
	registerMiddleware(router)
	registerRoutes(router, controller)

	http.ListenAndServe(":8080", router)
}
