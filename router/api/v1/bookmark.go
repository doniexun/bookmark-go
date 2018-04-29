package v1

import (
	"log"
	"strings"
	"github.com/gin-gonic/gin"
	"github.com/astaxie/beego/validation"
	"github.com/GallenHu/bookmarkgo/model"
	"github.com/microcosm-cc/bluemonday"
	"github.com/GallenHu/bookmarkgo/pkg/utils"
	"github.com/GallenHu/bookmarkgo/pkg/setting"
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

type ModifyBmAction struct {
	Id int `json: id`
	Title string `json:"title"`
	Url string `json:"url"`
	Tag string `json:"tag"`
	FolderId int `json:"folderid"`
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

	tags := strings.Split(strings.TrimSpace(tag), ",")
	index := utils.InArray("___", tags) // _ repeat 3
	var isprivate uint = 0
	if index > -1 {
		isprivate = 1
	}

	model.AddBookmark(title, url, tag, userid.(int), folderid, isprivate)

	c.JSON(200, gin.H{
        "code" : 200,
        "msg" : "success",
        "data" : nil,
    })
}

func GetBookmark(c *gin.Context) {
	// todo
}

func GetBookmarks(c *gin.Context) {
	folderid := c.Query("folderId")
	page := c.Query("page")

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

	showprivate, priexists := c.Get("showprivate")
	if !priexists {
		showprivate = 0
	}

	log.Println("showprivate ", showprivate)
	bookmarks := model.GetBookmarksByFolderId(showprivate.(uint), utils.Str2int(page, 1), userid.(int), utils.Str2int(folderid, 0))
	c.JSON(200, gin.H{
		"code": 200,
		"msg": "success",
		"data": bookmarks,
	})
}

func ModifyBookmark(c *gin.Context) {
	var errors []string
	var modifyaction ModifyBmAction

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

	c.BindJSON(&modifyaction)
	bookmarkid := modifyaction.Id
	bookmarktitle := modifyaction.Title
	bookmarkurl := modifyaction.Url
	bookmarktag := modifyaction.Tag
	bookmarkfolderid := modifyaction.FolderId

	// log.Println(bookmarkid)
	// bookmarkidint := utils.Str2int(bookmarkid, 0)
	if bookmarkid == 0 {
		errors = append(errors, "书签id有误")
		c.JSON(200, gin.H{
			"code" : 400,
			"msg" : "failed",
			"data" : errors,
		})

		return
	}

	bm, err := model.GetBookmarkById(bookmarkid, userid.(int))
	if err != nil {
		errors = append(errors, "书签id有误.")
		c.JSON(200, gin.H{
			"code" : 400,
			"msg" : "failed",
			"data" : errors,
		})

		return
	}

	tags := strings.Split(strings.TrimSpace(bookmarktag), ",")
	index := utils.InArray("___", tags) // _ repeat 3
	var isprivate uint = 0
	if index > -1 {
		isprivate = 1
	}

	model.ModifyBookmark(bm, userid.(int), bookmarktitle, bookmarkurl, bookmarktag, bookmarkfolderid, isprivate)

	c.JSON(200, gin.H{
        "code" : 200,
        "msg" : "success",
        "data" : modifyaction,
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

		return
	}

	model.DeleteBookmarkByIds(ids, userid.(int))
	c.JSON(200, gin.H{
        "code" : 200,
        "msg" : "success",
        "data" : nil,
    })
}

func SearchBookmarks(c *gin.Context) {
	var errors []string

	if setting.AllowSearch != 1 {
		errors = append(errors, "搜索功能已关闭")
		c.JSON(200, gin.H{
			"code" : 400,
			"msg" : "fail",
			"data" : errors,
		})
		return
	}

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

	folderid := c.Query("folderId")
	key := c.Query("keyword")
	key = strings.TrimSpace(key)

	if key == "" {
		c.JSON(200, gin.H{
			"code" : 400,
			"msg" : "success",
			"data" : nil,
		})
		return
	}

	bookmarks := model.SearchBookmarks(userid.(int), utils.Str2int(folderid, 0), key)
	c.JSON(200, gin.H{
		"code": 200,
		"msg": "success",
		"data": bookmarks,
	})
}
