package router

import (
	"github.com/gin-gonic/gin"
	v1 "ipfs-library-server/api/v1"
)

func InitAuthRouter(Router *gin.RouterGroup) {
	AuthRouter := Router.Group("auth")
	{
		AuthRouter.POST("/login", v1.Login)
	}
}
