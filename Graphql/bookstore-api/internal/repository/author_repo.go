package repository

import (
	"context"
	"fmt"

	"bookstore-api/internal/db"
	"bookstore-api/internal/models"
)

// CreateAuthor inserts a new author
func CreateAuthor(author *models.Author) error {
	sql := `INSERT INTO authors (name, bio) VALUES ($1, $2) RETURNING id`
	return db.Pool.QueryRow(context.Background(), sql, author.Name, author.Bio).Scan(&author.ID)
}

// GetAuthorByID fetches an author by ID
func GetAuthorByID(id int) (*models.Author, error) {
	sql := `SELECT id, name, bio FROM authors WHERE id=$1`
	row := db.Pool.QueryRow(context.Background(), sql, id)

	author := &models.Author{}
	err := row.Scan(&author.ID, &author.Name, &author.Bio)
	if err != nil {
		return nil, fmt.Errorf("failed to get author: %w", err)
	}
	return author, nil
}

// ListAuthors returns all authors
func ListAuthors() ([]models.Author, error) {
	sql := `SELECT id, name, bio FROM authors`
	rows, err := db.Pool.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("failed to list authors: %w", err)
	}
	defer rows.Close()

	var authors []models.Author
	for rows.Next() {
		var a models.Author
		if err := rows.Scan(&a.ID, &a.Name, &a.Bio); err != nil {
			return nil, err
		}
		authors = append(authors, a)
	}
	return authors, nil
}
