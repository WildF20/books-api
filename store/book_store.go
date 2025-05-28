package store

import (
	"sync"

	"books-api/model"
)

type bookStore struct {
	books  map[int]model.Book
	lastID int
	mu     sync.Mutex
}

var instance *bookStore
var once sync.Once

func GetBookStore() *bookStore {
	once.Do(func() {
		instance = &bookStore{
			books: make(map[int]model.Book),
		}
	})
	return instance
}

// AddBook adds a new book and auto-increments the ID
func (bs *bookStore) AddBook(book model.Book) model.Book {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	bs.lastID++
	book.ID = bs.lastID
	bs.books[book.ID] = book
	return book
}

// GetBook retrieves a book by ID
func (bs *bookStore) GetBook(id int) (model.Book, bool) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	book, exists := bs.books[id]
	return book, exists
}

// ListBooks returns all books
func (bs *bookStore) ListBooks() []model.Book {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	books := make([]model.Book, 0, len(bs.books))
	for _, b := range bs.books {
		books = append(books, b)
	}
	return books
}

// checkBookExists checks if a book exists by ID
func (bs *bookStore) CheckBookExists(id int) bool {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	_, exists := bs.books[id]
	return exists
}

// UpdateBook updates an existing book by ID
func (bs *bookStore) UpdateBook(id int, book model.Book) (model.Book, bool) {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	if _, exists := bs.books[id]; !exists {
		return model.Book{}, false
	}

	book.ID = id
	bs.books[id] = book
	return book, true
}

// DeleteBook removes a book by ID
func (bs *bookStore) DeleteBook(id int) bool {
	bs.mu.Lock()
	defer bs.mu.Unlock()

	if _, exists := bs.books[id]; !exists {
		return false
	}

	delete(bs.books, id)
	return true
}