package v1

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func UI(c *gin.Context) {
	c.Redirect(http.StatusMovedPermanently, "/ui/")
}
