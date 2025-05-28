package main

import (
	"books-api/routes"
	"books-api/app/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

	// Define the routes for the books API
    r.Route("/books", api.BookRoutes())

	http.ListenAndServe(":3000", r)
}
