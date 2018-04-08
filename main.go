package main

import (
	// "github.com/gin-gonic/gin"
	"github.com/GallenHu/bookmarkgo/router"
)

func main() {
	r := router.InitRouter()

	r.Run(":3000") // :8080
}
