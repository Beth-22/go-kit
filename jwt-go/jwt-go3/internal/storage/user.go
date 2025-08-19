package storage

import (
	"JWT-GO3/internal/models"
	"context"
	"errors"
)

// CreateUser inserts a new user into the database.
func CreateUser(user models.User) error {
	sql := `INSERT INTO users (username, password, role) VALUES ($1, $2, $3)`

	_, err := Pool.Exec(context.Background(), sql, user.Username, user.Password, user.Role)
	if err != nil {
		return err
	}
	return nil    //no error occured
}

// GetUser fetches a user by username.
func GetUser(username string) (models.User, error) {
	sql := `SELECT username, password, role FROM users WHERE username = $1`

	var user models.User
	err := Pool.QueryRow(context.Background(), sql, username).Scan(&user.Username, &user.Password, &user.Role)
	if err != nil {
		return models.User{}, errors.New("user not found")
	}
	return user, nil
}

// UpdateUser updates an existing user's data.
func UpdateUser(user models.User) error {
	sql := `UPDATE users SET password=$1, role=$2 WHERE username=$3`

	commandTag, err := Pool.Exec(context.Background(), sql, user.Password, user.Role, user.Username)
	if err != nil {
		return err
	}
	if commandTag.RowsAffected() == 0 {
		return errors.New("user not found")
	}
	return nil
}
/*what is commandTag?
commandTag is a value returned by the Exec method from the pgxpool package (PostgreSQL Go driver).

It’s of type pgconn.CommandTag.

It represents metadata about the result of the executed SQL command.

What info does CommandTag contain?
How many rows were affected by the command (INSERT, UPDATE, DELETE, etc.).

The command status string, e.g., "UPDATE 1" or "INSERT 0".

It allows you to check if the operation was successful in terms of rows modified.
*/