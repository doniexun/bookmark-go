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

// JSON response format
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

func GetFoldersByPage(pagenum int, userid int) []*FolderJson {
	var folders []*FolderJson

	offset := (pagenum - 1) * PAGESIZE

	rows, err := db.Model(&Folder{}).
		Select("id, name").
		Where(Folder{UserId: userid}).
		Offset(offset).
		Limit(PAGESIZE).
		Order("updated_at desc, created_at desc").
		Rows()
	if err != nil {
		panic(err)
	}
	for rows.Next() {
		// define var in each loop
		var folderjson FolderJson
		db.ScanRows(rows, &folderjson)
		// do something
		folders = append(folders, &folderjson)
	}

	return folders
}

// func GetFolderTotal(maps interface {}) int {
// 	var count int
// 	db.Model(&Folder{}).Where(maps).Count(&count)
// 	return count
// }
