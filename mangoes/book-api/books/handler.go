package books

import (
	"encoding/json"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
)

// Book struct and in-memory storage
type Book struct {
	ID     string `json:"id"`
	Title  string `json:"title"`
	Genre   string `json:"genre"`
	Status string `json:"status"` // "read", "not read", "to be read"
}

var books = make(map[string]Book)

// CreateBookHandler handles POST /books
func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	var b Book
	if err := json.NewDecoder(r.Body).Decode(&b); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}
	b.ID = uuid.New().String()
	books[b.ID] = b

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(b)
}

// ListBooksHandler handles GET /books
func ListBooksHandler(w http.ResponseWriter, r *http.Request) {
	bookList := make([]Book, 0, len(books))
	for _, b := range books {
		bookList = append(bookList, b)
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(bookList)
}

// Optional: Add router registration helper
func RegisterRoutes(r *mux.Router) {
	r.HandleFunc("/books", CreateBookHandler).Methods("POST")
	r.HandleFunc("/books", ListBooksHandler).Methods("GET")
}
