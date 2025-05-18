package address

import (
	_ "demo/api/docs"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes initializes address routes under /api/v1
// @Summary     Initialize address routes
// @Description Sets up address endpoints
// @Tags        router address
// @BasePath    /api/v1/address
func RegisterRoutes(r *gin.Engine) *gin.Engine {
	api := r.Group("/api/v1/address")
	{
		api.POST("/restaurant/:id", createRestaurantAddress)
		api.GET("/restaurant/:id", getAddressByRestaurantId)
		api.GET("/:id", getAddressById)
		api.PUT("/:id", updateAddress)
		api.DELETE("/:id", deleteAddress)
	}
	return r
}
