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

	r.GET("/", v1.Hello)

	apiv1 := r.Group("/bookmarkapi/v1")
	{
		apiv1.GET("/captcha", v1.GetCaptcha)

		apiv1.POST("/user", v1.Signup)  // 注册
		apiv1.POST("/auth", v1.Signin)  // 登录
		apiv1.POST("/auth/signout", jwt.JWT(), v1.Signout) // 登出

		apiv1.GET("/user", jwt.JWT(), v1.GetUserInfo) // 使用中间件
		apiv1.PUT("/user", jwt.JWT(), v1.ModifyUserInfo)

		apiv1.POST("/bookmark", jwt.JWT(), v1.NewBookmark)
		apiv1.PUT("/bookmark", jwt.JWT(), v1.ModifyBookmark)
		apiv1.PUT("/folder", jwt.JWT(), v1.ModifyFolder)
		apiv1.POST("/folder", jwt.JWT(), v1.NewFolder)
		apiv1.GET("/folders", jwt.JWT(), v1.GetFolders)
		apiv1.GET("/bookmarks", jwt.JWT(), v1.GetBookmarks)

		apiv1.GET("/search", jwt.JWT(), v1.SearchBookmarks)
	}

	// delete
	apiv1del := r.Group("/bookmarkapi/v1/del")
	apiv1del.Use(jwt.JWT())
	{
		apiv1del.POST("/bookmarks", v1.DelBookmarks)
		apiv1del.POST("/folders", v1.DelFolders)
	}

	return r
}
