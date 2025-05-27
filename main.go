package main

import (
	"books-api/routes"
	"net/http"

	"github.com/go-chi/chi/v5"
    "github.com/go-chi/chi/v5/middleware"
)

func main() {
	r := chi.NewRouter()
	r.Use(middleware.Logger)

    r.Route("/books", api.BookRoutes())

	http.ListenAndServe(":3000", r)
}
