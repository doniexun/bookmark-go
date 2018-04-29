package model

import (
	"log"
	"time"
	"database/sql"
	"github.com/jinzhu/gorm"
)

type Bookmark struct {
	Model

	Title		string	`gorm:"not null;type:varchar(255);DEFAULT:'Untitle'" json:"title"`
	Url			string	`gorm:"not null;type:varchar(255)" json:"url"`
	Tag			string  `gorm:"not null;type:varchar(255)" json:"tag"`
	UserId		int		`gorm:"not null;" json:"user_id"`
	FolderId	int		`gorm:"not null;" json:"folder_id"`
	IsPrivate	uint	`gorm:"not null;type:tinyint;DEFAULT:0" json:"is_private"`
}

// JSON response format
type BookmarkJson struct {
	Id      	int		`json:"id"`
	Title		string	`json:"title"`
	Url     	string	`json:"url"`
	Tag			string 	`json:"tag"`
	FolderId	int		`gorm:"not null;" json:"folderId"`
	IsPrivate	uint	`json:"isPrivate"`
}

// models callbacks
func (bookmark *Bookmark) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedAt", time.Now().Unix())
    scope.SetColumn("UpdatedAt", time.Now().Unix())
    return nil
}
func (bookmark *Bookmark) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("UpdatedAt", time.Now().Unix())
    return nil
}

func (Bookmark) TableName() string { // 自定义表名
	return "bookmark"
}


func AddBookmark(title string, url string, tag string, userid int, folderid int, isprivate uint) error {
	log.Println("ppp")
	log.Println(isprivate)
	bookmark := Bookmark{Title: title, Url: url, Tag: tag, UserId: userid, FolderId: folderid, IsPrivate: isprivate}
	err := db.Create(&bookmark).Error

	return err
}

func GetBookmarkById(id int, userid int) (*Bookmark, error) {
	var bookmark Bookmark
	err := db.Model(&Bookmark{}).
		Select("id, title, url, tag, folder_id").
		Where(Bookmark{UserId: userid}).
		Where("id = ?", id).
		First(&bookmark).
		Error

	return &bookmark, err;
}

func GetBookmarksByFolderId(showprivate uint, pagenum int, userid int, folderid int) []*BookmarkJson  {
	var bookmarks []*BookmarkJson
	var rows *sql.Rows
	var err error
	var inwhere interface{}
	var props string = "id, title, url, tag, folder_id, is_private"
	var orders string = "updated_at desc, created_at desc"

	offset := (pagenum - 1) * PAGESIZE

	if folderid == 0 {
		inwhere = Bookmark{UserId: userid}
	} else {
		inwhere = Bookmark{UserId: userid, FolderId: folderid}
	}

	if showprivate == 1 {
		rows, err = db.Model(&Bookmark{}).
			Select(props).
			Where(inwhere).
			Offset(offset).
			Limit(PAGESIZE).
			Order(orders).
			Rows()
	} else {
		rows, err = db.Model(&Bookmark{}).
			Select(props).
			Where(inwhere).
			Where("is_private = ?", 0).
			Offset(offset).
			Limit(PAGESIZE).
			Order(orders).
			Rows()
	}

	if err != nil {
		panic(err)
	}
	for rows.Next() {
		// define var in each loop
		var bookmarkjson BookmarkJson
		db.ScanRows(rows, &bookmarkjson)
		// do something
		bookmarks = append(bookmarks, &bookmarkjson)
	}

	return bookmarks
}

func ModifyBookmark(bm *Bookmark, userid int, title string, url string, tag string, folderid int) {
	bm.UserId = userid
	bm.Title = title
	bm.Url = url
	bm.Tag = tag
	bm.FolderId = folderid
	db.Save(&bm)
}

func DeleteBookmarkByIds(ids []int, userid int) bool {
	db.Where(Bookmark{UserId: userid}).Where("id in (?)", ids).Delete(Bookmark{})
	return true
}

func SearchBookmarks(userid int, folderid int, keyword string) []*BookmarkJson {
	var bookmarks []*BookmarkJson
	var rows *sql.Rows
	var err error
	offset := 0
	limit := 20

	if folderid == 0 {
		rows, err = db.Model(&Bookmark{}).
			Select("id, title, url, tag").
			Where(Bookmark{UserId: userid}).
			Where("concat(title, url, tag) like ?", "%" + keyword + "%").
			Offset(offset).
			Limit(limit).
			Order("updated_at desc, created_at desc").
			Rows()
	} else {
		rows, err = db.Model(&Bookmark{}).
			Select("id, title, url, tag").
			Where(Bookmark{UserId: userid, FolderId: folderid}).
			Where("concat(title, url, tag) like ?", "%" + keyword + "%").
			Offset(offset).
			Limit(limit).
			Order("updated_at desc, created_at desc").
			Rows()
	}

	if err != nil {
		panic(err)
	}
	for rows.Next() {
		// define var in each loop
		var bookmarkjson BookmarkJson
		db.ScanRows(rows, &bookmarkjson)
		// do something
		bookmarks = append(bookmarks, &bookmarkjson)
	}

	return bookmarks
}