package handler

import (
	"books-api/data/model"
	"books-api/data/store"

	"encoding/json"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

// CreateBook handles the creation of a new book.
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book model.Book
	bs := store.GetBookStore()
	
	// Decode and validate the incoming JSON request body
	if err := json.NewDecoder(r.Body).Decode(&book); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return
	}
	// End Validate JSON

	// Store the book in the model
	bs.AddBook(book)

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Successfully created book"}`))
}

// GetBooks retrieves all books and returns them in JSON format.
func GetBooks(w http.ResponseWriter, r *http.Request) {
	bs := store.GetBookStore()

	// Return all books in JSON format
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bs.ListBooks())
}

// GetBookByID retrieves a book by its ID and returns it in JSON format.
func GetBookByID(w http.ResponseWriter, r *http.Request) {
	idStr := chi.URLParam(r, "id")
	bs := store.GetBookStore()

	// Validate the book ID
	id, err := strconv.Atoi(idStr)
	if err != nil {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	// End validate
	
	// Retrieve the book by ID
	book, ok := bs.GetBook(id)
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
	bs := store.GetBookStore()

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
	if exists := bs.CheckBookExists(bookID); !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	// End Validate book data

	// Update the book in the model
	bs.UpdateBook(bookID, book)

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Successfully updated book with ID ` + id + `"}`))
}

func DeleteBook(w http.ResponseWriter, r *http.Request) {
	id := chi.URLParam(r, "id")
	bs := store.GetBookStore()

	// Validate the book ID and data
	bookID, err := strconv.Atoi(id)
	if err != nil || bookID <= 0 {
		http.Error(w, "Invalid book ID", http.StatusBadRequest)
		return
	}
	if exists := bs.CheckBookExists(bookID); !exists {
		http.Error(w, "Book not found", http.StatusNotFound)
		return
	}
	// End Validate book ID and data

	// Delete the book from the model
	bs.DeleteBook(bookID)

	// Respond with a success message
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte(`{"message": "Successfully deleted book with ID ` + id + `"}`))
}