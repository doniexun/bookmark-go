package v1

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"github.com/GallenHu/bookmarkgo/model"
	"github.com/microcosm-cc/bluemonday"
)

type FolderCommand struct {
	Name string `json:"name"`
}

func NewFolder(c *gin.Context) {
	var errors []string
	var folderCommand FolderCommand

	userid, exists := c.Get("userid")
	if !exists {
		errors = append(errors, "读取用户信息失败")
		c.JSON(200, gin.H{
			"code" : 500,
			"msg" : "failed",
			"data" : errors,
		})

		return
	}

	c.BindJSON(&folderCommand)

	name := folderCommand.Name

	p := bluemonday.UGCPolicy()
	name = p.Sanitize(name)

	valid := validation.Validation{}
	valid.Required(name, "name").Message("名称不能为空")

	if valid.HasErrors() {
        for _, err := range valid.Errors {
			log.Println(err.Key, err.Message)
			errors = append(errors, err.Message)
		}

		c.JSON(200, gin.H{
			"code" : 400,
			"msg" : "failed",
			"data" : errors,
		})

		return
	}

	model.AddFolder(name, userid.(int))

	c.JSON(200, gin.H{
        "code" : 200,
        "msg" : "success",
        "data" : make(map[string]string),
    })
}

func GetFolders(c *gin.Context) {
	var errors []string
	userid, exists := c.Get("userid")
	if !exists {
		errors = append(errors, "读取用户信息失败")
		c.JSON(200, gin.H{
			"code" : 500,
			"msg" : "failed",
			"data" : errors,
		})

		return
	}

	folders := model.GetFoldersByPage(1, userid.(int))

	c.JSON(200, gin.H{
		"code": 200,
		"msg": "success",
		"data": folders,
	})
}

func DelFolders(c *gin.Context) {
	var errors []string
	var deleteaction DeleteAction

	userid, exists := c.Get("userid")
	if !exists {
		errors = append(errors, "读取用户信息失败")
		c.JSON(200, gin.H{
			"code" : 500,
			"msg" : "failed",
			"data" : errors,
		})

		return
	}

	c.BindJSON(&deleteaction)
	ids := deleteaction.Id

	if len(ids) == 0 {
		errors = append(errors, "id 为空")
		c.JSON(200, gin.H{
			"code" : 400,
			"msg" : "fail",
			"data" : errors,
		})

		return
	}

	model.DeleteFolderByIds(ids, userid.(int))
	c.JSON(200, gin.H{
        "code" : 200,
        "msg" : "success",
        "data" : nil,
    })
}
