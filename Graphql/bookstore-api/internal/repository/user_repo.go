package repository

import (
	"context"
	//"errors"
	"fmt"

	"bookstore-api/internal/db"
	"bookstore-api/internal/models"
)

// CreateUser inserts a new user into the profile table
func CreateUser(user models.User) error {
	sql := `INSERT INTO profile (username, email, password_hash, role)
	        VALUES ($1, $2, $3, $4)`

	_, err := db.Pool.Exec(context.Background(), sql, user.Username, user.Email, user.PasswordHash, user.Role)
	if err != nil {
		return fmt.Errorf("failed to create user: %w", err)
	}
	return nil
}

// GetUserByID fetches a user by their ID
func GetUserByID(id int) (*models.User, error) {
	sql := `SELECT id, username, email, password_hash, role FROM profile WHERE id = $1`

	row := db.Pool.QueryRow(context.Background(), sql, id)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by ID: %w", err)
	}
	return user, nil
}

// GetUserByEmail fetches a user by their email
func GetUserByEmail(email string) (*models.User, error) {
	sql := `SELECT id, username, email, password_hash, role FROM profile WHERE email = $1`

	row := db.Pool.QueryRow(context.Background(), sql, email)
	user := &models.User{}
	err := row.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role)
	if err != nil {
		return nil, fmt.Errorf("failed to get user by email: %w", err)
	}
	return user, nil
}

// ListUsers returns all users
func ListUsers() ([]models.User, error) {
	sql := `SELECT id, username, email, password_hash, role FROM profile`
	rows, err := db.Pool.Query(context.Background(), sql)
	if err != nil {
		return nil, fmt.Errorf("failed to list users: %w", err)
	}
	defer rows.Close()

	var users []models.User
	for rows.Next() {
		var user models.User
		if err := rows.Scan(&user.ID, &user.Username, &user.Email, &user.PasswordHash, &user.Role); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}
