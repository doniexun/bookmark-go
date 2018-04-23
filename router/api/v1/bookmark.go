package v1

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"github.com/GallenHu/bookmarkgo/model"
)

type BookmarkCommand struct {
	Title string `json:"title"`
	Url string `json:"url"`
	Tag string `json:"tag"`
	FolderId int `json:"folderid"`
}

func NewBookmark(c *gin.Context) {
	var errors []string
	var bookmarkCommand BookmarkCommand

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

	c.BindJSON(&bookmarkCommand)

	title := bookmarkCommand.Title
	url := bookmarkCommand.Url
	tag := bookmarkCommand.Tag
	folderid := bookmarkCommand.FolderId // default 0

	valid := validation.Validation{}

	valid.Required(url, "url").Message("url不能为空")

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

	if title == "" {
		title = "Untitled"
	}

	model.AddBookmark(title, url, tag, userid.(int), folderid)

	c.JSON(200, gin.H{
        "code" : 200,
        "msg" : "success",
        "data" : nil,
    })
}
