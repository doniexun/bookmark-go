package main

import "fmt"
// import "net/http"
import "github.com/gin-gonic/gin"
import "github.com/GallenHu/bookmarkgo/common"
import "github.com/GallenHu/bookmarkgo/routes"

func main() {
	config := common.LoadConfig("/path/to/conf.json")
	fmt.Println(config.Name)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	routes.Register(router)

	router.Run(":8080")
}