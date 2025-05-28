package api

import (
	"books-api/app/handler"

	"github.com/go-chi/chi/v5"
)

// define the routes for the books API
func BookRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", handler.GetBooks)
		r.Get("/{id}", handler.GetBookByID)
		r.Post("/", handler.CreateBook)
		r.Put("/{id}", handler.UpdateBook)
		r.Delete("/{id}", handler.DeleteBook)
	}
}