// @title       Restaurant API
// @version     1.0
// @description This is the REST API for managing restaurants & products.
// @host        localhost:8080
// @BasePath    /api/v1
package main

import (
	"demo/api/middleware"
	"demo/config"
	"demo/internal/address/controller"
	"demo/internal/health/controller"
	"demo/internal/product/controller"
	"demo/internal/restaurant/controller"
	"github.com/gin-gonic/gin"
)

//go:generate swag init -g main.go -o ../../api/docs -d ../../cmd/app,../../internal/health/controller,../../internal/restaurant/controller,../../internal/address/controller,../../internal/product/controller --parseInternal

func main() {
	config.ConnectDB()
	defer config.DisconnectDB()

	r := gin.New()
	r.Use(gin.Logger(), gin.Recovery(), middleware.ErrorMiddleware())
	registerRoutes(r)
	_ = r.Run(":8080")
}

func registerRoutes(r *gin.Engine) {
	health.RegisterRoutes(r)
	restaurant.RegisterRoutes(r)
	product.RegisterRoutes(r)
	address.RegisterRoutes(r)
}
