package router

import (
	"time"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/gin-contrib/cors"
	"github.com/GallenHu/bookmarkgo/router/api/v1"
	"github.com/GallenHu/bookmarkgo/middleware/jwt"
	"github.com/GallenHu/bookmarkgo/pkg/setting"
)

func InitRouter() *gin.Engine {
	r := gin.New()
	r.Use(gin.Logger())

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{setting.AppCors},
		AllowMethods:     []string{"GET", "POST", "PUT", "DETELE"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowOriginFunc: func(origin string) bool {
			// return origin == "https://github.com"
			return strings.HasPrefix(origin, "chrome-extension://") // 允许chrome-extension访问
		},
		MaxAge: 12 * time.Hour,
	}))

	apiv1 := r.Group("/api/v1")
	{
		apiv1.GET("/captcha", v1.GetCaptcha)
		apiv1.GET("/add", v1.AddFolder)

		apiv1.POST("/user", v1.Signup)  // 注册
		apiv1.POST("/auth", v1.Signin)  // 登录
		apiv1.POST("/auth/signout", v1.Signout) // 登出

		apiv1.GET("/userinfo", jwt.JWT(), v1.GetUserInfo) // 使用中间件

		apiv1.POST("/bookmark", jwt.JWT(), v1.NewBookmark)
	}

	return r
}
