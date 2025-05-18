package product

import (
	"time"

	"github.com/google/uuid"
)

// -- Requests --

// CreateProductRequest defines the payload for creating a new product.
type CreateProductRequest struct {
	RestaurantID uuid.UUID `json:"restaurant_id" binding:"required"`
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	Price        uint      `json:"price" binding:"required,gt=0"`
}

// UpdateProductRequest defines the payload for updating an existing product.
type UpdateProductRequest struct {
	RestaurantID *uuid.UUID `json:"restaurant_id,omitempty"`
	Name         *string    `json:"name,omitempty"`
	Description  *string    `json:"description,omitempty"`
	Price        *uint      `json:"price,omitempty"`
}

// -- Responses --

// ProductResponse defines the JSON structure returned for product resources.
type ProductResponse struct {
	ID           uuid.UUID `json:"id" binding:"required"`
	RestaurantID uuid.UUID `json:"restaurant_id" binding:"required"`
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	Price        uint      `json:"price" binding:"required"`
	CreatedAt    time.Time `json:"created_at" binding:"required"`
	UpdatedAt    time.Time `json:"updated_at" binding:"required"`
}
