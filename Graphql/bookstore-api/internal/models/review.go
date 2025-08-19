package models

import "github.com/google/uuid"

type Review struct {
    ID     uuid.UUID
    BookID uuid.UUID
    UserID int
    Text   string
    Rating int // 1-5
}

