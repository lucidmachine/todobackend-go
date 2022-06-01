package main

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()

	registerMiddleware(r)
	registerRoutes(r)

	http.ListenAndServe(":8080", r)
}
