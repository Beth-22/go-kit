package services

import (
	"errors"

	"JWT-GO3/internal/models"
	"JWT-GO3/internal/storage"
)

// CreateUserService adds a new user if they don't exist
func CreateUserService(user models.User) error {
	_, err := storage.GetUser(user.Username)
	if err == nil {
		return errors.New("user already exists")
	}
	return storage.CreateUser(user)
}

// GetUserService fetches a user by username
func GetUserSperervice(username string) (models.User, error) {
	return storage.GetUser(username)
}

// UpdateUserService updates user data
func UpdateUserService(user models.User) error {
	return storage.UpdateUser(user)
}
