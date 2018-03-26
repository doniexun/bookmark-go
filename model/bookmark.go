package model

import "github.com/jinzhu/gorm"

type Bookmark struct {
	gorm.Model
	Title		string	`gorm:"not null;type:varchar(255)";DEFAULT:'Untitle'`
	Url			string	`gorm:"not null;type:varchar(255)"`
	UserId		int		`gorm:"not null;"`
}