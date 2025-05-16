package controller

import (
	"demo/internal/ping/service"
	"github.com/gin-gonic/gin"
	"net/http"
)

type PingResponse struct {
	Message string `json:"message" example:"pong"`
}

// Ping
// @Tags	Health
// @Router	/ping [get]
func Ping(c *gin.Context) {
	response := service.GetPing()
	c.JSON(http.StatusOK, PingResponse{Message: response})
}
