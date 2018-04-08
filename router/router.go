package router

import (
	"github.com/gin-gonic/gin"
	"github.com/GallenHu/bookmarkgo/router/api/v1"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/adduser", v1.AddUser)
		apiv1.GET("/add", v1.AddFolder)
		apiv1.GET("/test", v1.GetFolders)
	}

	return r
}
