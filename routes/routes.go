package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index", gin.H{
			"pagename": "main",
			"title": "hello world",
		})
	})

	router.GET("/signup", func(c *gin.Context) {
		c.HTML(http.StatusOK, "signup", gin.H{
			"pagename": "signup",
			"title": "hello world",
		})
	})
}