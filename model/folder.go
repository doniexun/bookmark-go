package model

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Folder struct {
	Model

	Name		string	`gorm:"not null;type:varchar(255)";DEFAULT:'Unnamed' json:"name"`
	UserId		int		`gorm:"not null;" json:"user_id"`
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

func AddFolder(name string, userId int) bool {
	db.Create(&Folder{
		Name: name,
		UserId: userId,
	})

	return true
}

// func GetFolder(pageNum int, pageSize int, maps interface {}) (Folder []Folder) {
// 	db.Where(maps).Offset(pageNum).Limit(pageSize).Find(&Folder)

// 	return
// }

func GetFolderTotal(maps interface {}) int {
	var count int
	db.Model(&Folder{}).Where(maps).Count(&count)
	return count
}
