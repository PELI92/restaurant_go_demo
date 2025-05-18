package dto

import (
	address_dto "demo/internal/address/controller/dto"
	"github.com/google/uuid"
	"time"
)

// --Requests--

type CreateRestaurantRequest struct {
	Name    string                            `json:"name" binding:"required"`
	Phone   string                            `json:"phone" binding:"required"`
	Email   string                            `json:"email" binding:"required"`
	Website string                            `json:"website,omitempty"`
	Address *address_dto.CreateAddressRequest `json:"address,omitempty"`
}

type UpdateRestaurantRequest struct {
	Name    *string `json:"name"`
	Phone   *string `json:"phone"`
	Email   *string `json:"email"`
	Website *string `json:"website"`
	Status  *string `json:"status"`
}

// --Responses--

type RestaurantResponse struct {
	ID        uuid.UUID                    `json:"id"`
	Name      string                       `json:"name" binding:"required"`
	Phone     *string                      `json:"phone" binding:"required"`
	Email     *string                      `json:"email" binding:"required"`
	Website   *string                      `json:"website,omitempty"`
	Status    int8                         `json:"status" binding:"required"`
	Address   *address_dto.AddressResponse `json:"address,omitempty"`
	CreatedAt time.Time                    `json:"created_at" binding:"required"`
	UpdatedAt time.Time                    `json:"updated_at" binding:"required"`
}
