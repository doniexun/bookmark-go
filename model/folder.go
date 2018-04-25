package model

import (
	"time"
	"github.com/jinzhu/gorm"
)

const PAGESIZE = 10

type Folder struct {
	Model

	Name		string	`gorm:"not null;type:varchar(255)";DEFAULT:'Unnamed' json:"name"`
	UserId		int		`gorm:"not null;" json:"user_id"`
}

type FolderJson struct {
	Name	string	`json:"name"`
	Id      int		`json:"id"`
}

// models callbacks
func (folder *Folder) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedAt", time.Now().Unix())
    return nil
}
func (folder *Folder) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("UpdatedAt", time.Now().Unix())
    return nil
}

func (Folder) TableName() string { // 自定义表名
  	return "folder"
}

func AddFolder(name string, userid int) bool {
	db.Create(&Folder{
		Name: name,
		UserId: userid,
	})

	return true
}

func GetFoldersByPage(pagenum int, userid int) FolderJson {
	var folderJson FolderJson

	offset := (pagenum - 1) * PAGESIZE

	rows, err := db.Model(&Folder{}).Select("id, name").Where(Folder{UserId: userid}).Where("deleted_on = ?", 0).Offset(offset).Limit(PAGESIZE).Rows()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		db.ScanRows(rows, &folderJson)
		// do something
	}

	return folderJson
}

// func GetFolderTotal(maps interface {}) int {
// 	var count int
// 	db.Model(&Folder{}).Where(maps).Count(&count)
// 	return count
// }
