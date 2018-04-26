package v1

import (
	"log"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"github.com/GallenHu/bookmarkgo/model"
	"github.com/microcosm-cc/bluemonday"
	"github.com/GallenHu/bookmarkgo/pkg/utils"
)

type BookmarkCommand struct {
	Title string `json:"title"`
	Url string `json:"url"`
	Tag string `json:"tag"`
	FolderId int `json:"folderid"`
}

type DeleteAction struct {
	Id []int `json: id`
}

func NewBookmark(c *gin.Context) {
	var errors []string
	var bookmarkcommand BookmarkCommand

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

	c.BindJSON(&bookmarkcommand)

	title := bookmarkcommand.Title
	url := bookmarkcommand.Url
	tag := bookmarkcommand.Tag
	folderid := bookmarkcommand.FolderId // default 0

	p := bluemonday.UGCPolicy()
	title = p.Sanitize(title)
	url = p.Sanitize(url)
	tag = p.Sanitize(tag)

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

func GetBookmarks(c *gin.Context) {
	folderid := c.Query("folderId")

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

	if folderid == "" {
		folderid = utils.Int2str(0)
	}

	bookmarks := model.GetBookmarksByFolderId(1, userid.(int), utils.Str2int(folderid))
	c.JSON(200, gin.H{
		"code": 200,
		"msg": "success",
		"data": bookmarks,
	})
}

func DelBookmarks(c *gin.Context) {
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
	}

	model.DeleteBookmarkByIds(ids, userid.(int))
	c.JSON(200, gin.H{
        "code" : 200,
        "msg" : "success",
        "data" : nil,
    })
}
