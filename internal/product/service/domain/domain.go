package domain

import (
	"github.com/google/uuid"
	"time"
)

type Product struct {
	ID           uuid.UUID `json:"id" binding:"required"`
	RestaurantID uuid.UUID `json:"restaurant_id" binding:"required"`
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	Price        uint      `json:"price" binding:"required"`
	CreatedAt    time.Time `json:"created_at" binding:"required"`
	UpdatedAt    time.Time `json:"updated_at" binding:"required"`
}
