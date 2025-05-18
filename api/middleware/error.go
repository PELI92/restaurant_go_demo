package middleware

import (
	"errors"
	"fmt"
	"net/http"

	apiErr "demo/api/error"
	"github.com/gin-gonic/gin"
)

// ErrorMiddleware intercepts errors added to the context (c.Error(err))
// and converts them into a uniform JSON response with the appropriate HTTP status code.
func ErrorMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()

		fmt.Printf("ErrorMiddleware len: %d\n", len(c.Errors))
		fmt.Printf("ErrorMiddleware len: %+v\n", c.Errors)
		if len(c.Errors) > 0 {
			err := c.Errors[0].Err
			var appErr *apiErr.APIError
			if errors.As(err, &appErr) {
				c.AbortWithStatusJSON(appErr.StatusCode(), appErr)
				return
			}
			c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		}
	}
}
