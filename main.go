package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/GallenHu/bookmarkgo/pkg/setting"
	"github.com/GallenHu/bookmarkgo/router"
	"github.com/GallenHu/bookmarkgo/pkg/redis"
)

func main() {
	fmt.Println("DbHost:", setting.DbHost)
	fmt.Println("RedisHost:", setting.RedisHost)
	fmt.Println("AppCors:", setting.AppCors)

	gin.SetMode(setting.AppMode)

	redis.TestConnect()

	r := router.InitRouter()

	r.Run(":" + setting.AppPort) // :8080
}
