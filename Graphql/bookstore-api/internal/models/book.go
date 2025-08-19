package models

import "github.com/google/uuid"

type Book struct {
    ID           uuid.UUID
    Title        string
    AuthorID     int
    PublishedYear int
    Genre        string
    Page         int
}
