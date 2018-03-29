package routes

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func Register(router *gin.Engine) {
	// Page
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

	// API
	v1 := router.Group("/api/v1")
	{
		v1.POST("/signup", func(c *gin.Context) {
			c.JSON(200, gin.H{
				"message": "pong",
			})
		})
	}

}