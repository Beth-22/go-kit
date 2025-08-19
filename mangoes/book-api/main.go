package main

import (
	"fmt"
	"net/http"

	"BOOK-API/books"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// Register books routes
	books.RegisterRoutes(r)

	fmt.Println("Server running on http://localhost:8080")
	http.ListenAndServe(":8080", r)
}
