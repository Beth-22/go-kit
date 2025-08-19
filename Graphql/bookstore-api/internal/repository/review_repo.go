package repository

import (
	"context"
	"fmt"

	"bookstore-api/internal/db"
	"bookstore-api/internal/models"

	"github.com/google/uuid"
)

// CreateReview inserts a new review
func CreateReview(review *models.Review) error {
	sql := `INSERT INTO reviews (book_id, user_id, text, rating) VALUES ($1, $2, $3, $4) RETURNING id`
	return db.Pool.QueryRow(context.Background(), sql, review.BookID, review.UserID, review.Text, review.Rating).Scan(&review.ID)
}

// GetReviewByID fetches a review by UUID
func GetReviewByID(id uuid.UUID) (*models.Review, error) {
	sql := `SELECT id, book_id, user_id, text, rating FROM reviews WHERE id=$1`
	row := db.Pool.QueryRow(context.Background(), sql, id)

	review := &models.Review{}
	err := row.Scan(&review.ID, &review.BookID, &review.UserID, &review.Text, &review.Rating)
	if err != nil {
		return nil, fmt.Errorf("failed to get review: %w", err)
	}
	return review, nil
}

// ListReviews returns all reviews
func ListReviews() ([]models.Review, error) {
	sql := `SELECT id, book_id, user_id, text, rating FROM reviews`
	rows, err := db.Pool.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("failed to list reviews: %w", err)
	}
	defer rows.Close()

	var reviews []models.Review
	for rows.Next() {
		var r models.Review
		if err := rows.Scan(&r.ID, &r.BookID, &r.UserID, &r.Text, &r.Rating); err != nil {
			return nil, err
		}
		reviews = append(reviews, r)
	}
	return reviews, nil
}
