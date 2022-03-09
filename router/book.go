package router

import (
	"github.com/gin-gonic/gin"
	v1 "ipfs-library-server/api/v1"
	"ipfs-library-server/middleware"
)

func InitBookRouter(Router *gin.RouterGroup) {
	BookRouter := Router.Group("book")
	BookRouter.Use(middleware.JWTAuth())
	{
		BookRouter.POST("/add", v1.AddBook)
		BookRouter.GET("/list", v1.QueryBooksList)
		BookRouter.DELETE("/:id", v1.DeleteBook)
	}
	BookRouter.GET("/search", v1.SearchBooks)
}
