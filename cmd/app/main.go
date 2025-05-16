// @title       Restaurant API
// @version     1.0
// @description This is the REST API for managing restaurants & products.
// @host        localhost:8080
// @BasePath    /api/v1
package main

import (
	"demo/api"
	"demo/config"
)

//go:generate swag init -g main.go -o ../../api/docs -d ../../cmd/app,../../internal/ping/controller,../../internal/restaurant/controller,../../internal/product/controller,../../internal/restaurant/model,../../internal/product/model --parseInternal

func main() {
	config.ConnectDB()
	defer config.DisconnectDB()

	r := api.SetupRouter()
	r.Run(":8080")
}
