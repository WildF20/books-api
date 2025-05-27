package main

import (
	"books-api/routes"
	"books-api/middleware"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

    r.Route("/books", api.BookRoutes())

	http.ListenAndServe(":3000", r)
}
