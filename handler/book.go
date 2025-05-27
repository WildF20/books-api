package handler

import (
	"books-api/model"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// CreateBook handles the creation of a new book.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	
	// Decode and validate the incoming JSON request body
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// End Validate JSON

	// Store the book in the model
	model.LastID++
    book.ID = model.LastID
    model.Books[book.ID] = book

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Successfully created book ` + strconv.Itoa(book.ID) + `"}`))
}

// GetBooks retrieves all books and returns them in JSON format.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// Return all books in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(model.Books)
}

// GetBookByID retrieves a book by its ID and returns it in JSON format.
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")

	// Validate the book ID
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	// End validate
	
	// Retrieve the book by ID
	book, ok := model.Books[id]
	if !ok {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}

	// Return the book in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

// UpdateBook updates an existing book by its ID.
func UpdateBook(w http.ResponseWriter, r *http.Request)  {
	id := chi.URLParam(r, "id")
	var book model.Book

	// Decode and validate the incoming JSON request body
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// End Validate JSON

	// Validate the book data
	bookID, err := strconv.Atoi(id)
	if err != nil || bookID <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	if _, exists := model.Books[bookID]; !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	// End Validate book data

	// Update the book in the model
	book.ID = bookID
	model.Books[bookID] = book

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Successfully updated book with ID ` + id + `"}`))
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")

	// Validate the book ID and data
	bookID, err := strconv.Atoi(id)
	if err != nil || bookID <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	if _, exists := model.Books[bookID]; !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	// End Validate book ID and data

	// Delete the book from the model
	delete(model.Books, bookID)

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Successfully deleted book with ID ` + id + `"}`))
}