package router

import (
	"github.com/gin-gonic/gin"
	"github.com/GallenHu/bookmarkgo/router/api/v1"
	"github.com/GallenHu/bookmarkgo/middleware/jwt"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/captcha", v1.GetCaptcha)
		apiv1.GET("/add", v1.AddFolder)

		apiv1.POST("/user", v1.Signup)  // 注册
		apiv1.POST("/auth", v1.Signin)  // 登录
		apiv1.POST("/auth/signout", v1.Signout) // 登出

		apiv1.GET("/info", jwt.JWT(), v1.GetUserInfo) // 使用中间件
	}

	return r
}
