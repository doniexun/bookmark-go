package main

import (
	"flag"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/GallenHu/bookmarkgo/common"
	"github.com/GallenHu/bookmarkgo/routes"
)

func main() {
	configPath := flag.String("c", "./conf.json", "config file path")
	flag.Parse() // 执行命令行解析

	config := common.LoadConfig(*configPath)
	fmt.Println(config.Name)

	router := gin.Default()
	router.LoadHTMLGlob("templates/*")

	router.Static("/static", "./static")
	routes.Register(router)

	router.Run(":8080")
}