package helper

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func GetID(c *gin.Context) (uuid.UUID, error) {
	return uuid.Parse(c.Param("id"))
}
