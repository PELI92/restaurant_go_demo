package param

import "github.com/google/uuid"

type CreateProductParam struct {
	RestaurantID uuid.UUID `json:"restaurant_id" binding:"required"`
	Name         string    `json:"name" binding:"required"`
	Description  string    `json:"description" binding:"required"`
	Price        uint      `json:"price" binding:"required"`
}

type UpdateProductParam struct {
	RestaurantID *uuid.UUID `json:"restaurant_id" binding:"required"`
	Name         *string    `json:"name" binding:"required"`
	Description  *string    `json:"description" binding:"required"`
	Price        *uint      `json:"price" binding:"required"`
}
