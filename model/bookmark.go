package model

import (
	"time"
	"github.com/jinzhu/gorm"
)

type Bookmark struct {
	Model

	Title		string	`gorm:"not null;type:varchar(255)";DEFAULT:'Untitle' json:"title"`
	Url			string	`gorm:"not null;type:varchar(255)" json:"url"`
	UserId		int		`gorm:"not null;" json:"user_id"`
	FolderId	int		`gorm:"not null;" json:"folder_id"`
}

// models callbacks
func (bookmark *Bookmark) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedAt", time.Now().Unix())
    return nil
}
func (bookmark *Bookmark) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("UpdatedAt", time.Now().Unix())
    return nil
}

func (Bookmark) TableName() string { // 自定义表名
	return "bookmark"
}
