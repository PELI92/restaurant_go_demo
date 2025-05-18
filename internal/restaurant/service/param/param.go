package param

type CreateRestaurantParam struct {
	Name    string  `json:"name" binding:"required"`
	Phone   string  `json:"phone" binding:"required"`
	Email   string  `json:"email" binding:"required"`
	Website *string `json:"website,omitempty"`
}

type UpdateRestaurantParam struct {
	Name    *string `json:"name"`
	Phone   *string `json:"phone"`
	Email   *string `json:"email"`
	Website *string `json:"website"`
	Status  *string `json:"status"`
}
