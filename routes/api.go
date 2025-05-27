package api

import (
	"books-api/handler"

	"github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func BookRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Use(middleware.Logger)

		r.Get("/", handler.GetBooks)
		r.Get("/{id}", handler.GetBookByID)
		r.Post("/", handler.CreateBook)
		r.Put("/{id}", handler.UpdateBook)
		r.Delete("/{id}", handler.DeleteBook)
	}
}