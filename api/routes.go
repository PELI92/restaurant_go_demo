package api

import (
	_ "demo/api/docs"
	"demo/internal/ping/controller"
	product_controller "demo/internal/product/controller"
	restaurant_controller "demo/internal/restaurant/controller"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// SetupRouter initializes all routes under /api/v1
// @Summary     Initialize routes
// @Description Sets up ping, restaurant and product endpoints
// @Tags        router
// @BasePath    /api/v1
func SetupRouter() *gin.Engine {
	r := gin.Default()
	// Grupo con prefijo /api/v1
	api := r.Group("/api/v1")
	{
		api.GET("/ping", controller.Ping)

		api.POST("/restaurant", restaurant_controller.CreateRestaurant)
		api.GET("/restaurants", restaurant_controller.GetRestaurants)
		api.GET("/restaurant/:id", restaurant_controller.GetRestaurant)
		api.PUT("/restaurant/:id", restaurant_controller.UpdateRestaurant)
		api.DELETE("/restaurant/:id", restaurant_controller.DeleteRestaurant)

		api.POST("/product", product_controller.CreateProduct)
		api.GET("/products", product_controller.GetProducts)
		api.GET("/product/:id", product_controller.GetProduct)
		api.PUT("/product/:id", product_controller.UpdateProduct)
		api.DELETE("/product/:id", product_controller.DeleteProduct)
	}

	// Swagger UI en /api/v1/swagger/index.html
	// y el JSON de la spec en /api/v1/swagger/doc.json
	url := ginSwagger.URL("/api/v1/swagger/doc.json")
	api.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
