package product

import (
	_ "demo/api/docs"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes initializes product routes under /api/v1
// @Summary     Initialize product routes
// @Description Sets up product endpoints
// @Tags        router product
// @BasePath    /api/v1/product
func RegisterRoutes(r *gin.Engine) *gin.Engine {
	api := r.Group("/api/v1/product")
	{
		api.POST("/", createProduct)
		api.GET("/", getProducts)
		api.GET("/:id", getProduct)
		api.PUT("/:id", updateProduct)
		api.DELETE("/:id", deleteProduct)
	}
	return r
}
