package main

import (
	"github.com/gin-gonic/gin"
	"github.com/GallenHu/bookmarkgo/pkg/setting"
	"github.com/GallenHu/bookmarkgo/router"
)

func main() {
	gin.SetMode(setting.AppMode)

	r := router.InitRouter()

	r.Run(":" + setting.AppPort) // :8080
}
