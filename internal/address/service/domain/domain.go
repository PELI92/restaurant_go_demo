package domain

import (
	"github.com/google/uuid"
	"time"
)

type Address struct {
	ID         uuid.UUID `json:"id"`
	Street     string    `json:"street" `
	City       string    `json:"city" `
	State      string    `json:"state"`
	PostalCode string    `json:"postal_code"`
	Country    string    `json:"country"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
