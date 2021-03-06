package jwt

import (
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/GallenHu/bookmarkgo/pkg/utils"
	"github.com/GallenHu/bookmarkgo/pkg/redis"
)

func JWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		var errors []string
		var token string

		values, _ := c.Request.Header["Authorization"]; // values type: array

		if len(values) == 0 {
			errors = append(errors, "缺少token")
			c.JSON(401, gin.H{
				"code" : 401,
				"msg" : "unauthorized",
				"data" : errors,
			})

			c.Abort()
            return
		}

		tempArr := strings.Split(values[0], " ")
		if len(tempArr) > 1 && tempArr[0] == "Bearer" {
			token = tempArr[1]
		} else {
			token = ""

			errors = append(errors, "token无效")
			c.JSON(401, gin.H{
				"code" : 401,
				"msg" : "unauthorized",
				"data" : errors,
			})

			c.Abort()
            return
		}

		claims, err := utils.ParseToken(token)
		if err != nil {
			errors = append(errors, "token无效")
			c.JSON(401, gin.H{
				"code" : 401,
				"msg" : "unauthorized",
				"data" : errors,
			})

			c.Abort()
            return
		}

		if redis.GetUserToken(claims.Id) != token {
			errors = append(errors, "token过期")
			c.JSON(401, gin.H{
				"code" : 401,
				"msg" : "unauthorized",
				"data" : errors,
			})

			c.Abort()
            return
		}

		c.Set("userid", claims.Id)
		c.Next()
	}
}
