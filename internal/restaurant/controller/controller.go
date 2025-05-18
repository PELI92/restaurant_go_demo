package restaurant

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	apiErr "demo/api/error"
	apiHelper "demo/api/helper"
	dto "demo/internal/restaurant/controller/dto"
	"demo/internal/restaurant/mapper"
	"demo/internal/restaurant/service"
)

const maxPaginationSize = 100

// createRestaurant godoc
// @Summary     Create a new restaurant
// @Description Receives restaurant details and creates a new record.
// @Tags        restaurants
// @Accept      json
// @Produce     json
// @Param       input  body      dto.CreateRestaurantRequest  true  "Restaurant to create"
// @Success     201    {object}  dto.RestaurantResponse
// @Failure     400    {object}  map[string]string  "Bad request"
// @Failure     500    {object}  map[string]string  "Internal server error"
// @Router      /restaurant [post]
func createRestaurant(c *gin.Context) {
	var input dto.CreateRestaurantRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.Error(apiErr.NewValidation(err.Error()))
		return
	}

	param := mapper.NewCreateRestaurantParam(input)
	restaurant, svcErr := service.CreateRestaurant(c.Request.Context(), param)
	if svcErr != nil {
		_ = c.Error(svcErr)
		return
	}

	response := mapper.NewRestaurantResponse(*restaurant)
	c.JSON(http.StatusCreated, response)
}

// GetRestaurants godoc
// @Summary     List all restaurants
// @Description Retrieves a paginated list of restaurants.
// @Tags        restaurants
// @Produce     json
// @Param       page  query     int  false  "Page number"
// @Param       size  query     int  false  "Page size"
// @Success     200   {object}  map[string]interface{}
// @Failure     500   {object}  map[string]string  "Internal server error"
// @Router      /restaurants [get]
func getRestaurants(c *gin.Context) {
	page, err := strconv.Atoi(c.DefaultQuery("page", "1"))
	if err != nil || page < 1 {
		page = 1
	}
	size, err := strconv.Atoi(c.DefaultQuery("size", "10"))
	if err != nil || size < 1 {
		size = 10
	}
	if size > maxPaginationSize {
		size = maxPaginationSize
	}
	offset := (page - 1) * size

	restaurants, svcErr := service.GetRestaurantsPaginated(c.Request.Context(), size, offset)
	if svcErr != nil {
		_ = c.Error(svcErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  restaurants,
		"page":  page,
		"size":  size,
		"count": len(restaurants),
	})
}

// GetRestaurant godoc
// @Summary     Get a restaurant by ID
// @Description Retrieves the restaurant identified by the given ID.
// @Tags        restaurants
// @Produce     json
// @Param       id    path      string  true  "Restaurant ID"
// @Success     200   {object}  dto.RestaurantResponse
// @Failure     400   {object}  map[string]string  "Bad request"
// @Failure     404   {object}  map[string]string  "Restaurant not found"
// @Failure     500   {object}  map[string]string  "Internal server error"
// @Router      /restaurant/{id} [get]
func getRestaurant(c *gin.Context) {
	id, err := apiHelper.GetID(c)
	if err != nil {
		_ = c.Error(apiErr.NewValidation("invalid restaurant ID"))
		return
	}

	restaurant, svcErr := service.GetRestaurantByID(c.Request.Context(), id)
	if svcErr != nil {
		_ = c.Error(svcErr)
		return
	}
	response := mapper.NewRestaurantResponse(*restaurant)
	c.JSON(http.StatusOK, response)
}

// UpdateRestaurant godoc
// @Summary     Update a restaurant by ID
// @Description Updates the restaurant identified by the given ID.
// @Tags        restaurants
// @Accept      json
// @Produce     json
// @Param       id    path      string                     true  "Restaurant ID"
// @Param       input body      dto.UpdateRestaurantRequest true  "Updated restaurant data"
// @Success     200   {object}  dto.RestaurantResponse
// @Failure     400   {object}  map[string]string  "Bad request"
// @Failure     404   {object}  map[string]string  "Restaurant not found"
// @Failure     500   {object}  map[string]string  "Internal server error"
// @Router      /restaurant/{id} [put]
func updateRestaurant(c *gin.Context) {
	id, err := apiHelper.GetID(c)
	if err != nil {
		_ = c.Error(apiErr.NewValidation("invalid restaurant ID"))
		return
	}

	var input dto.UpdateRestaurantRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		_ = c.Error(apiErr.NewValidation(err.Error()))
		return
	}

	param := mapper.NewUpdateRestaurantParam(input)
	restaurant, svcErr := service.UpdateRestaurant(c.Request.Context(), id, param)
	if svcErr != nil {
		_ = c.Error(svcErr)
		return
	}

	response := mapper.NewRestaurantResponse(*restaurant)
	c.JSON(http.StatusOK, response)
}

// DeleteRestaurant godoc
// @Summary     Delete a restaurant by ID
// @Description Deletes the restaurant identified by the given ID.
// @Tags        restaurants
// @Produce     json
// @Param       id    path      string  true  "Restaurant ID"
// @Success     200   {object}  map[string]string  "Deleted successfully"
// @Failure     400   {object}  map[string]string  "Bad request"
// @Failure     404   {object}  map[string]string  "Restaurant not found"
// @Failure     500   {object}  map[string]string  "Internal server error"
// @Router      /restaurant/{id} [delete]
func deleteRestaurant(c *gin.Context) {
	id, err := apiHelper.GetID(c)
	if err != nil {
		_ = c.Error(apiErr.NewValidation("invalid restaurant ID"))
		return
	}

	svcErr := service.DeleteRestaurant(c.Request.Context(), id)
	if svcErr != nil {
		_ = c.Error(svcErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "restaurant deleted"})
}
