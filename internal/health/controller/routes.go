package health

import (
	_ "demo/api/docs"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// RegisterRoutes initializes health routes
// @Summary     Initialize health routes
// @Description Sets up ping and swagger endpoints
// @Tags        router health
// @BasePath    /
func RegisterRoutes(r *gin.Engine) *gin.Engine {
	r.GET("/ping", ping)

	url := ginSwagger.URL("/swagger/doc.json")
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler, url))

	return r
}
