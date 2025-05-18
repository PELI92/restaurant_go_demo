package entity

import (
	"github.com/google/uuid"
	"time"
)

type Restaurant struct {
	ID        uuid.UUID  `json:"id"`
	Name      string     `json:"name"`
	Phone     *string    `json:"phone"`
	Email     *string    `json:"email"`
	Website   *string    `json:"website"`
	AddressID *uuid.UUID `json:"address_id"`
	Status    int8       `json:"status"`
	CreatedAt time.Time  `json:"created_at"`
	UpdatedAt time.Time  `json:"updated_at"`
}
