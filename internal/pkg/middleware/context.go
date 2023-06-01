package middleware

import (
	"github.com/gin-gonic/gin"
	"go-ecm/pkg/log"
)

const UsernameKey = "username"

func Context() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set(log.KeyRequestId, c.GetString(XRequestIDKey))
	}
}
