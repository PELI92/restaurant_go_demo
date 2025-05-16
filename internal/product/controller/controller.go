package controller

import (
	errors "demo/api/error"
	"demo/internal/product/model"
	"demo/internal/product/service"
	"net/http"

	"github.com/gin-gonic/gin"
)

// CreateProduct creates a new product.
// @Summary     Create a new product
// @Description Receives product details and creates a new product record.
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       input  body      model.Product  true  "Product to create"
// @Success     201    {object}  model.Product
// @Failure     400    {object}  map[string]string  "Bad request"
// @Failure     500    {object}  map[string]string  "Internal server error"
// @Router      /product [post]
func CreateProduct(c *gin.Context) {
	var input model.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	created, err := service.CreateProduct(input)
	if err != nil {
		errors.AbortWithError(c, err)
		return
	}
	c.JSON(http.StatusCreated, created)
}

// GetProducts returns all products.
// @Summary     List all products
// @Description Retrieves a list of products.
// @Tags        products
// @Produce     json
// @Success     200  {array}   model.Product
// @Failure     500  {object}  map[string]string  "Internal server error"
// @Router      /products [get]
func GetProducts(c *gin.Context) {
	products, err := service.GetProducts("")
	if err != nil {
		errors.AbortWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, products)
}

// GetProduct retrieves a product by its ID.
// @Summary     Get a product by ID
// @Description Retrieves the product identified by the given ID.
// @Tags        products
// @Produce     json
// @Param       id   path      string  true  "Product ID"
// @Success     200  {object}  model.Product
// @Failure     400  {object}  map[string]string  "Bad request"
// @Failure     404  {object}  map[string]string  "Product not found"
// @Failure     500  {object}  map[string]string  "Internal server error"
// @Router      /product/{id} [get]
func GetProduct(c *gin.Context) {
	id := c.Param("id")
	product, err := service.GetProductByID(id)
	if err != nil {
		errors.AbortWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, product)
}

// UpdateProduct updates an existing product by its ID.
// @Summary     Update a product by ID
// @Description Updates the product identified by the given ID with new data.
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id     path      string         true  "Product ID"
// @Param       input  body      model.Product  true  "Updated product data"
// @Success     200    {object}  model.Product
// @Failure     400    {object}  map[string]string  "Bad request"
// @Failure     404    {object}  map[string]string  "Product not found"
// @Failure     500    {object}  map[string]string  "Internal server error"
// @Router      /product/{id} [put]
func UpdateProduct(c *gin.Context) {
	id := c.Param("id")
	var input model.Product
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	updated, err := service.UpdateProduct(id, input)
	if err != nil {
		errors.AbortWithError(c, err)
		return
	}

	c.JSON(http.StatusOK, updated)
}

// DeleteProduct deletes a product by its ID.
// @Summary     Delete a product by ID
// @Description Deletes the product identified by the given ID.
// @Tags        products
// @Produce     json
// @Param       id   path      string  true  "Product ID"
// @Success     200  {object}  map[string]string  "Deleted successfully"
// @Failure     400  {object}  map[string]string  "Bad request"
// @Failure     404  {object}  map[string]string  "Product not found"
// @Failure     500  {object}  map[string]string  "Internal server error"
// @Router      /product/{id} [delete]
func DeleteProduct(c *gin.Context) {
	id := c.Param("id")
	err := service.DeleteProduct(id)
	if err != nil {
		errors.AbortWithError(c, err)
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "deleted"})
}
