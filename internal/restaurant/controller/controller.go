package controller

import (
	errors "demo/api/error"
	"demo/internal/restaurant/model"
	"demo/internal/restaurant/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateRestaurant creates a new restaurant.
// @Summary     Create a new restaurant
// @Description Receives restaurant details and creates a new record.
// @Tags        restaurants
// @Accept      json
// @Produce     json
// @Param       input  body      model.Restaurant  true  "Restaurant to create"
// @Success     201    {object}  model.Restaurant
// @Failure     400    {object}  map[string]string  "Bad request"
// @Failure     500    {object}  map[string]string  "Internal server error"
// @Router      /restaurant [post]
func CreateRestaurant(c *gin.Context) {
	var input model.Restaurant
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := service.CreateRestaurant(input)
	if err != nil {
		errors.AbortWithError(c, err)
		return
	}
	c.JSON(http.StatusCreated, created)
}

// GetRestaurants returns all restaurants.
// @Summary     List all restaurants
// @Description Retrieves a list of restaurants.
// @Tags        restaurants
// @Produce     json
// @Success     200  {array}   model.Restaurant
// @Failure     500  {object}  map[string]string  "Internal server error"
// @Router      /restaurants [get]
func GetRestaurants(c *gin.Context) {
	restaurants, err := service.GetRestaurants("")
	if err != nil {
		errors.AbortWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, restaurants)
}

// GetRestaurant retrieves a restaurant by its ID.
// @Summary     Get a restaurant by ID
// @Description Retrieves the restaurant identified by the given ID.
// @Tags        restaurants
// @Produce     json
// @Param       id   path      string  true  "Restaurant ID"
// @Success     200  {object}  model.Restaurant
// @Failure     400  {object}  map[string]string  "Bad request"
// @Failure     404  {object}  map[string]string  "Restaurant not found"
// @Failure     500  {object}  map[string]string  "Internal server error"
// @Router      /restaurant/{id} [get]
func GetRestaurant(c *gin.Context) {
	id := c.Param("id")
	restaurant, err := service.GetRestaurantByID(id)
	if err != nil {
		errors.AbortWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, restaurant)
}

// UpdateRestaurant updates an existing restaurant by its ID.
// @Summary     Update a restaurant by ID
// @Description Updates the restaurant identified by the given ID.
// @Tags        restaurants
// @Accept      json
// @Produce     json
// @Param       id     path      string             true  "Restaurant ID"
// @Param       input  body      model.Restaurant   true  "Updated restaurant data"
// @Success     200    {object}  model.Restaurant
// @Failure     400    {object}  map[string]string  "Bad request"
// @Failure     404    {object}  map[string]string  "Restaurant not found"
// @Failure     500    {object}  map[string]string  "Internal server error"
// @Router      /restaurant/{id} [put]
func UpdateRestaurant(c *gin.Context) {
	id := c.Param("id")
	var input model.Restaurant
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := service.UpdateRestaurant(id, input)
	if err != nil {
		errors.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DeleteRestaurant deletes a restaurant by its ID.
// @Summary     Delete a restaurant by ID
// @Description Deletes the restaurant identified by the given ID.
// @Tags        restaurants
// @Produce     json
// @Param       id   path      string  true  "Restaurant ID"
// @Success     200  {object}  map[string]string  "Deleted successfully"
// @Failure     400  {object}  map[string]string  "Bad request"
// @Failure     404  {object}  map[string]string  "Restaurant not found"
// @Failure     500  {object}  map[string]string  "Internal server error"
// @Router      /restaurant/{id} [delete]
func DeleteRestaurant(c *gin.Context) {
	id := c.Param("id")
	err := service.DeleteRestaurant(id)
	if err != nil {
		errors.AbortWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
