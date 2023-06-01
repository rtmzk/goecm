package goecmserver

import (
	v1 "go-ecm/internal/goecmserver/controller/v1"
	"go-ecm/internal/goecmserver/controller/v1/agent"
	mydocker "go-ecm/internal/goecmserver/controller/v1/docker"
	"go-ecm/internal/goecmserver/controller/v1/menu"
	"go-ecm/internal/goecmserver/controller/v1/swarm"
	"go-ecm/internal/goecmserver/controller/v1/systemConfig"
	"go-ecm/internal/goecmserver/controller/v1/user"
	"go-ecm/internal/goecmserver/store/sqlite"
	"go-ecm/internal/goecmserver/util/docker"
	"go-ecm/internal/pkg/code"
	"go-ecm/internal/pkg/middleware"
	"go-ecm/internal/pkg/middleware/auth"
	"go-ecm/pkg/core"

	"github.com/gin-contrib/static"
	"github.com/gin-gonic/gin"
	"github.com/marmotedu/errors"
	"github.com/rakyll/statik/fs"

	_ "go-ecm/statik"
)

func initRouter(g *gin.Engine) {
	installMiddleware(g)
	installAPI(g)
}

func installMiddleware(g *gin.Engine) {}

func installAPI(g *gin.Engine) *gin.Engine {
	var statikFs static.ServeFileSystem
	statikFs = &GinFS{}
	statikFs.(*GinFS).FS, _ = fs.New()

	g.Use(static.Serve("/", statikFs))

	jwtStrategy, _ := newJWTAuth().(auth.JWTStrategy)

	g.POST("/v1/login", jwtStrategy.LoginHandler)
	g.POST("/v1/logout", jwtStrategy.LogoutHandler)
	g.POST("/v1/refresh", jwtStrategy.RefreshHandler)

	g.GET("/", v1.UI)
	g.StaticFS("/ui", statikFs)

	auto := newAutoAuth()
	g.NoRoute(auto.AuthFunc(), func(c *gin.Context) {
		core.WriteResponse(c, errors.WithCode(code.ErrPageNotFound, "Page not found"), nil)
	})

	storeIns, _ := sqlite.GetSqliteFactoryOr(nil)
	dockerCliIns, _ := docker.GetDockerClientOr()
	v1 := g.Group("/v1")
	{
		userv1 := v1.Group("/users")
		{
			userController := user.NewUserController(storeIns)
			userv1.POST("", userController.Create)
			userv1.Use(auto.AuthFunc(), middleware.Validation())
			userv1.GET("/:name", userController.Get)
			userv1.GET("/current", userController.GetCurrentUser)
		}
		menuv1 := v1.Group("/menu")
		{
			menuv1.Use(auto.AuthFunc())
			menuController := menu.NewMenuController(storeIns)
			menuv1.GET("/list", menuController.Get)
		}
		swarmv1 := v1.Group("/swarm")
		{
			swarmController := swarm.NewSwarmController(storeIns)
			// 鉴权: 开发环境暂时关闭
			swarmv1.Use(auto.AuthFunc())
			// 获取ssh配置脚本
			swarmv1.GET("/script/get", swarmController.GetScript)
			// not used: 主机注册
			swarmv1.POST("/host/register", swarmController.Register)
			// 主机创建
			swarmv1.POST("/host/create", swarmController.Create)
			// sshkey 连接检查
			swarmv1.POST("/host/check", swarmController.Check)
		}
		dockerv1 := v1.Group("/docker")
		{
			// 鉴权: 开发环境暂时关闭
			dockerv1.Use(auto.AuthFunc())
			dockerController := mydocker.NewDockerController(storeIns, dockerCliIns)
			dockerv1.GET("/images", dockerController.ImageGet)
			dockerv1.DELETE("/images/delete", dockerController.ImageDelete)
			dockerv1.POST("/images/upload", dockerController.ImageUpload)
			dockerv1.GET("/services", dockerController.ServiceGet)
			dockerv1.GET("/services/:id", dockerController.ServiceInspect)
			dockerv1.DELETE("/services/:id", dockerController.ServiceDelete)
			dockerv1.POST("/services/update", dockerController.ServiceUpdate)
			dockerv1.GET("/nodes", dockerController.NodeGet)
			dockerv1.GET("/nodes/:id", dockerController.NodeInspect)
			dockerv1.POST("/nodes/update", dockerController.NodeUpdate)
			dockerv1.DELETE("/nodes/:id", dockerController.NodeDelete)
		}
		systemConfigv1 := v1.Group("/system/config")
		{
			systemConfigv1.Use(auto.AuthFunc())
			systemController := systemConfig.NewSystemConfigController(storeIns)
			systemConfigv1.GET("/all", systemController.GetAllConfig)
			systemConfigv1.GET("/token", systemController.GetToken)
		}
		agentv1 := v1.Group("/agent")
		{
			agentController := agent.NewAgentController(storeIns)
			agentv1.POST("/register", agentController.Register)
			agentv1.POST("/heartbeat", agentController.Heartbeat)
			agentv1.POST("/report", agentController.Report)
			agentv1.Use(auto.AuthFunc())
			agentv1.GET("/all", agentController.GetAgents)
			agentv1.GET("/getReportItems", agentController.GetReportItems)
			agentv1.POST("/deploy", agentController.DeployTask)
			agentv1.GET("/deploy/getDeployProgress", agentController.GetTaskProcess)
			agentv1.DELETE("/container/delete", agentController.DeleteContainer)
			agentv1.DELETE("/image/delete", agentController.DeleteImage)
			agentv1.POST("/image/export", agentController.ExportImage)
			agentv1.POST("/image/pull", agentController.PullImage)
			agentv1.GET("/envc/rules/get", agentController.EnvCheckPrepare)
			agentv1.POST("/envc/result/get", agentController.EnvCheckAction)
			agentv1.POST("/image/upload", agentController.ImageUpload)
			agentv1.POST("/image/import", agentController.ImageImport)
			agentv1.GET("/image/importProgress", agentController.GetImageImportProgress)
		}
	}
	return g
}
