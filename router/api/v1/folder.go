package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/GallenHu/bookmarkgo/model"
)

func AddFolder(c *gin.Context) {
	model.AddFolder("folderx", 123)

	fmt.Println("add success")

	c.JSON(200, gin.H{
        "code" : 200,
        "msg" : "success",
        "data" : make(map[string]string),
    })
}

func GetFolders(c *gin.Context) {
	maps := make(map[string]interface{})
	maps["id"] = 1

	c.JSON(200, gin.H{
		"code": 200,
		"msg": "success",
		"data": model.GetFolderTotal(maps),
	})
}
