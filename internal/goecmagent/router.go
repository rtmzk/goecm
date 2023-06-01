package goecmagent

import (
	"github.com/gin-gonic/gin"
	apiv1 "go-ecm/internal/goecmagent/api/v1"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installApis(g)
}

func installMiddleware(g *gin.Engine) {}

func installApis(g *gin.Engine) *gin.Engine {
	v1 := g.Group("/v1")
	{
		v1.POST("/multiPartUpload", apiv1.MultiPartUpload)
		v1.POST("/DeployCore/:action", apiv1.DeployCore)
		v1.GET("/terminal", apiv1.Terminal)
		v1.POST("/terminal/resize", apiv1.Resize)
		v1.DELETE("/container/delete", apiv1.DeleteContainer)
		v1.DELETE("/image/delete", apiv1.DeleteImage)
		v1.POST("/image/export", apiv1.ExportImage)
		v1.POST("/image/pull", apiv1.ImagePull)
		v1.GET("/envc", apiv1.Envc)
		v1.POST("/image/import", apiv1.ImageImport)
	}

	return g
}
