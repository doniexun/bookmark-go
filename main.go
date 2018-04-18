package main

import (
	"github.com/gin-gonic/gin"
	"github.com/GallenHu/bookmarkgo/pkg/setting"
	"github.com/GallenHu/bookmarkgo/router"
	"github.com/GallenHu/bookmarkgo/pkg/redis"
)

func main() {
	gin.SetMode(setting.AppMode)

	redis.TestConnect()

	r := router.InitRouter()

	r.Run(":" + setting.AppPort) // :8080
}
