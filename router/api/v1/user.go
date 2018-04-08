package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/GallenHu/bookmarkgo/model"
)

func AddUser(c *gin.Context) {
	model.AddUser("acd@mail.com", "123321")

	fmt.Println("add success")

	c.JSON(200, gin.H{
        "code" : 200,
        "msg" : "success",
        "data" : make(map[string]string),
    })
}
