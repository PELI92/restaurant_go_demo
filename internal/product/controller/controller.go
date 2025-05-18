package product

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	apiErr "demo/api/error"
	apiHelper "demo/api/helper"
	dto "demo/internal/product/controller/dto"
	"demo/internal/product/mapper"
	"demo/internal/product/service"
)

const maxPaginationSize = 100

// CreateProduct godoc
// @Summary     Create a new product
// @Description Receives product details and creates a new record.
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       input  body      dto.CreateProductRequest  true  "Product to create"
// @Success     201    {object}  dto.ProductResponse
// @Failure     400    {object}  map[string]string  "Bad request"
// @Failure     500    {object}  map[string]string  "Internal server error"
// @Router      /product [post]
func createProduct(c *gin.Context) {
	var input dto.CreateProductRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(apiErr.NewValidation(err.Error()))
		return
	}

	param := mapper.NewCreateProductParam(input)
	product, svcErr := service.CreateProduct(c, param)
	if svcErr != nil {
		c.Error(svcErr)
		return
	}

	response := mapper.NewProductResponse(product)
	c.JSON(http.StatusCreated, response)
}

// GetProducts godoc
// @Summary     List all products
// @Description Retrieves a paginated list of products.
// @Tags        products
// @Produce     json
// @Param       page  query     int  false  "Page number"
// @Param       size  query     int  false  "Page size"
// @Success     200   {object}  map[string]interface{}
// @Failure     500   {object}  map[string]string  "Internal server error"
// @Router      /products [get]
func getProducts(c *gin.Context) {
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

	products, svcErr := service.GetProductsPaginated(c, size, offset)
	if svcErr != nil {
		c.Error(svcErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data":  products,
		"page":  page,
		"size":  size,
		"count": len(products),
	})
}

// GetProduct godoc
// @Summary     Get a product by ID
// @Description Retrieves the product identified by the given ID.
// @Tags        products
// @Produce     json
// @Param       id   path      string  true  "Product ID"
// @Success     200  {object}  dto.ProductResponse
// @Failure     400  {object}  map[string]string  "Bad request"
// @Failure     404  {object}  map[string]string  "Product not found"
// @Failure     500  {object}  map[string]string  "Internal server error"
// @Router      /product/{id} [get]
func getProduct(c *gin.Context) {
	id, err := apiHelper.GetID(c)
	if err != nil {
		c.Error(apiErr.NewValidation("invalid product ID"))
		return
	}

	product, svcErr := service.GetProductByID(c, id)
	if svcErr != nil {
		c.Error(svcErr)
		return
	}

	response := mapper.NewProductResponse(*product)
	c.JSON(http.StatusOK, response)
}

// UpdateProduct godoc
// @Summary     Update a product by ID
// @Description Updates the product identified by the given ID.
// @Tags        products
// @Accept      json
// @Produce     json
// @Param       id     path      string                  true  "Product ID"
// @Param       input  body      dto.UpdateProductRequest true  "Updated product data"
// @Success     200    {object}  dto.ProductResponse
// @Failure     400    {object}  map[string]string  "Bad request"
// @Failure     404    {object}  map[string]string  "Product not found"
// @Failure     500    {object}  map[string]string  "Internal server error"
// @Router      /product/{id} [put]
func updateProduct(c *gin.Context) {
	id, err := apiHelper.GetID(c)
	if err != nil {
		c.Error(apiErr.NewValidation("invalid product ID"))
		return
	}

	var input dto.UpdateProductRequest
	if err := c.ShouldBindJSON(&input); err != nil {
		c.Error(apiErr.NewValidation(err.Error()))
		return
	}

	param := mapper.NewUpdateProductParam(input)
	product, svcErr := service.UpdateProduct(c, id, param)
	if svcErr != nil {
		c.Error(svcErr)
		return
	}

	response := mapper.NewProductResponse(product)
	c.JSON(http.StatusOK, response)
}

// DeleteProduct godoc
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
func deleteProduct(c *gin.Context) {
	id, err := apiHelper.GetID(c)
	if err != nil {
		c.Error(apiErr.NewValidation("invalid product ID"))
		return
	}

	svcErr := service.DeleteProduct(c, id)
	if svcErr != nil {
		c.Error(svcErr)
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "product deleted"})
}
