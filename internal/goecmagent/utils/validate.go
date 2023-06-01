package utils

import (
	"github.com/gin-gonic/gin"
	"go-ecm/internal/goecmagent/constand"
)

func ValidateToken(c *gin.Context) bool {
	token := c.GetString("token")
	if token == "" {
		return false
	}

	return token == constand.Token
}
