package restaurant

import (
	_ "demo/api/docs"
	"github.com/gin-gonic/gin"
)

// RegisterRoutes initializes restaurant routes under /api/v1
// @Summary     Initialize restaurant routes
// @Description Sets up restaurant endpoints
// @Tags        router restaurant
// @BasePath    /api/v1/restaurant
func RegisterRoutes(r *gin.Engine) *gin.Engine {
	api := r.Group("/api/v1/restaurant")
	{
		api.POST("/", createRestaurant)
		api.GET("/", getRestaurants)
		api.GET("/:id", getRestaurant)
		api.PUT("/:id", updateRestaurant)
		api.DELETE("/:id", deleteRestaurant)
	}

	return r
}
