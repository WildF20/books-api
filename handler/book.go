package handler

import (
	"books-api/model"
	"encoding/json"
	"net/http"

	"github.com/go-chi/chi/v5"
)

func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Welcome to the Books API!"}`))
}

func GetBookByID(w http.ResponseWriter, r *http.Request) {
    id := chi.URLParam(r, "id")
	
    book := model.Book{
        ID:           id,
        Title:        "Sample Book",
        Author:       "Sample Author",
        PublishedYear: 2024,
    }

    w.Header().Set("Content-Type", "application/json")
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(book)
}

func UpdateBook(w http.ResponseWriter, r *http.Request)  {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Successfully updated book with ID ` + id + `"}`))
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Successfully deleted book with ID ` + id + `"}`))
}