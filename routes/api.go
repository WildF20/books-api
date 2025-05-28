package api

import (
	"books-api/app/handler"
	"books-api/app/middleware"

	"github.com/go-chi/chi/v5"
)

// define the routes for the books API
func BookRoutes() func(r chi.Router) {
	return func(r chi.Router) {
		r.Get("/", handler.GetBooks)
		r.Get("/{id}", handler.GetBookByID)
		r.With(middleware.RequestValidationMiddleware).Post("/", handler.CreateBook)
		r.With(middleware.RequestValidationMiddleware).Put("/{id}", handler.UpdateBook)
		r.Delete("/{id}", handler.DeleteBook)
	}
}