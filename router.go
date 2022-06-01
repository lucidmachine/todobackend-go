package main

import (
	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
	"github.com/go-chi/cors"
)

func registerMiddleware(r chi.Router) {
	r.Use(middleware.Logger)
	r.Use(middleware.CleanPath)
	r.Use(middleware.RedirectSlashes)
	r.Use(middleware.Timeout(5000))
	r.Use(middleware.Compress(5))
	r.Use(middleware.Recoverer)
	r.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))
}

func registerRoutes(r chi.Router, c Controller) {
	r.Get("/", c.getTodos)
	r.Post("/", c.createTodo)
	r.Delete("/", c.deleteTodos)
}
