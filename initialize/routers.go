package initialize

import (
	"github.com/gin-gonic/gin"
	"ipfs-library-server/middleware"
	"ipfs-library-server/router"
)

func Routers() *gin.Engine {
	gin.SetMode(gin.DebugMode)
	r := gin.Default()
	r.Use(middleware.Cors())
	v1 := r.Group("api/v1")
	{
		router.InitAuthRouter(v1)
		router.InitUserRouter(v1)
		router.InitBookRouter(v1)
	}
	return r
}
