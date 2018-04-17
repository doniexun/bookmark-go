package main

import (
	// "github.com/gin-gonic/gin"
	"github.com/GallenHu/bookmarkgo/router"
	"github.com/GallenHu/bookmarkgo/pkg/redis"
)

func main() {
	redis.SetVal()
	r := router.InitRouter()

	r.Run(":3000") // :8080
}
