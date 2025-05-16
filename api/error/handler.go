package errors

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StatusCode(err *APIError) int {
	switch err.Kind {
	case ErrNotFound:
		return http.StatusNotFound
	case ErrValidation:
		return http.StatusBadRequest
	case ErrConflict:
		return http.StatusConflict
	default:
		return http.StatusInternalServerError
	}
}

// AbortWithError es un helper para los handlers
func AbortWithError(c *gin.Context, err *APIError) {
	c.AbortWithStatusJSON(StatusCode(err), gin.H{
		"error": err.Message,
	})
}
