package repository

import (
	"context"
	"fmt"

	"bookstore-api/internal/db"
	"bookstore-api/internal/models"

	"github.com/google/uuid"
)

// CreateBook inserts a new book
func CreateBook(book *models.Book) error {
	sql := `INSERT INTO books (title, author_id, published_year, genre, page) VALUES ($1, $2, $3, $4, $5) RETURNING id`
	return db.Pool.QueryRow(context.Background(), sql, book.Title, book.AuthorID, book.PublishedYear, book.Genre, book.Page).Scan(&book.ID)
}

// GetBookByID fetches a book by UUID
func GetBookByID(id uuid.UUID) (*models.Book, error) {
	sql := `SELECT id, title, author_id, published_year, genre, page FROM books WHERE id=$1`
	row := db.Pool.QueryRow(context.Background(), sql, id)

	book := &models.Book{}
	err := row.Scan(&book.ID, &book.Title, &book.AuthorID, &book.PublishedYear, &book.Genre, &book.Page)
	if err != nil {
		return nil, fmt.Errorf("failed to get book: %w", err)
	}
	return book, nil
}

// ListBooks returns all books
func ListBooks() ([]models.Book, error) {
	sql := `SELECT id, title, author_id, published_year, genre, page FROM books`
	rows, err := db.Pool.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("failed to list books: %w", err)
	}
	defer rows.Close()

	var books []models.Book
	for rows.Next() {
		var b models.Book
		if err := rows.Scan(&b.ID, &b.Title, &b.AuthorID, &b.PublishedYear, &b.Genre, &b.Page); err != nil {
			return nil, err
		}
		books = append(books, b)
	}
	return books, nil
}
