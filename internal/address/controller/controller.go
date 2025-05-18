package address

import (
	"net/http"

	"github.com/gin-gonic/gin"

	apiErr "demo/api/error"
	apiHelper "demo/api/helper"
	dto "demo/internal/address/controller/dto"
	"demo/internal/address/mapper"
	"demo/internal/address/service"
)

// CreateRestaurantAddress godoc
// @Summary     Create a new address
// @Description Receives address details and creates a new record for a restaurant.
// @Tags        addresses
// @Accept      json
// @Produce     json
// @Param       input  body      dto.CreateAddressRequest  true  "Address to create"
// @Success     201    {object}  dto.AddressResponse
// @Failure     400    {object}  map[string]string  "Bad request"
// @Failure     500    {object}  map[string]string  "Internal server error"
// @Router      /address [post]
func createRestaurantAddress(c *gin.Context) {
	var input dto.CreateAddressRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.Error(apiErr.NewValidation(err.Error()))
		return
	}

	param := mapper.NewCreateAddressParam(input)
	address, svcErr := service.CreateRestaurantAddress(c.Request.Context(), param)
	if svcErr != nil {
		_ = c.Error(svcErr)
		return
	}

	response := mapper.NewAddressResponse(*address)
	c.JSON(http.StatusCreated, response)
}

// GetAddressById godoc
// @Summary     Get an address by ID
// @Description Retrieves the address identified by the given address ID.
// @Tags        addresses
// @Produce     json
// @Param       id    path      string  true  "Address ID"
// @Success     200   {object}  dto.AddressResponse
// @Failure     400   {object}  map[string]string  "Bad request"
// @Failure     404   {object}  map[string]string  "Address not found"
// @Failure     500   {object}  map[string]string  "Internal server error"
// @Router      /address/{id} [get]
func getAddressById(c *gin.Context) {
	id, err := apiHelper.GetID(c)
	if err != nil {
		_ = c.Error(apiErr.NewValidation("invalid address ID"))
		return
	}

	address, svcErr := service.GetAddressById(c.Request.Context(), id)
	if svcErr != nil {
		_ = c.Error(svcErr)
		return
	}

	response := mapper.NewAddressResponse(*address)
	c.JSON(http.StatusOK, response)
}

// GetAddressByRestaurantId godoc
// @Summary     Get an address by restaurant ID
// @Description Retrieves the address associated with the given restaurant ID.
// @Tags        addresses
// @Produce     json
// @Param       restaurant_id   path      string  true  "Restaurant ID"
// @Success     200   {object}  dto.AddressResponse
// @Failure     400   {object}  map[string]string  "Bad request"
// @Failure     404   {object}  map[string]string  "Address not found"
// @Failure     500   {object}  map[string]string  "Internal server error"
// @Router      /address/restaurant/{restaurant_id} [get]
func getAddressByRestaurantId(c *gin.Context) {
	id, err := apiHelper.GetID(c)
	if err != nil {
		_ = c.Error(apiErr.NewValidation("invalid restaurant ID"))
		return
	}

	address, svcErr := service.GetAddressByRestaurantId(c.Request.Context(), id)
	if svcErr != nil {
		_ = c.Error(svcErr)
		return
	}

	response := mapper.NewAddressResponse(*address)
	c.JSON(http.StatusOK, response)
}

// UpdateAddress godoc
// @Summary     Update an address by ID
// @Description Updates the address identified by the given address ID.
// @Tags        addresses
// @Accept      json
// @Produce     json
// @Param       id     path      string                    true  "Address ID"
// @Param       input  body      dto.UpdateAddressRequest  true  "Updated address data"
// @Success     200    {object}  dto.AddressResponse
// @Failure     400    {object}  map[string]string  "Bad request"
// @Failure     404    {object}  map[string]string  "Address not found"
// @Failure     500    {object}  map[string]string  "Internal server error"
// @Router      /address/{id} [put]
func updateAddress(c *gin.Context) {
	id, err := apiHelper.GetID(c)
	if err != nil {
		_ = c.Error(apiErr.NewValidation("invalid address ID"))
		return
	}

	var input dto.UpdateAddressRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.Error(apiErr.NewValidation(err.Error()))
		return
	}

	param := mapper.NewUpdateAddressParam(input)
	address, svcErr := service.UpdateAddress(c.Request.Context(), id, param)
	if svcErr != nil {
		_ = c.Error(svcErr)
		return
	}

	response := mapper.NewAddressResponse(*address)
	c.JSON(http.StatusOK, response)
}

// DeleteAddress godoc
// @Summary     Delete an address by ID
// @Description Deletes the address identified by the given address ID.
// @Tags        addresses
// @Produce     json
// @Param       id    path      string  true  "Address ID"
// @Success     200   {object}  map[string]string  "Deleted successfully"
// @Failure     400   {object}  map[string]string  "Bad request"
// @Failure     404   {object}  map[string]string  "Address not found"
// @Failure     500   {object}  map[string]string  "Internal server error"
// @Router      /address/{id} [delete]
func deleteAddress(c *gin.Context) {
	id, err := apiHelper.GetID(c)
	if err != nil {
		_ = c.Error(apiErr.NewValidation("invalid address ID"))
		return
	}

	svcErr := service.DeleteAddress(c.Request.Context(), id)
	if svcErr != nil {
		_ = c.Error(svcErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "address deleted"})
}
