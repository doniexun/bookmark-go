package helper

import (
	"github.com/gin-gonic/gin"
)

func JsonSuccess(c *gin.Context, data) {
	c.JSON(200, gin.H{
		"data": data,
		"msg": "success",
		"code": 0,
	})
}

func JsonFailed(c *gin.Context, msg) {
	c.JSON(200, gin.H{
		"data": data,
		"msg": "success",
		"code": 0,
	})
}
