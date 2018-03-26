package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Mail		string	`gorm:"unique_index;not null;type:varchar(50)"`
	Password	string	`gorm:"not null;type:varchar(50)"`
}