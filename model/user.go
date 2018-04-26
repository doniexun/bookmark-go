package model

import (
	"time"
	"github.com/jinzhu/gorm"
)

type User struct {
	Model // 自定义的model

	Mail		string	`gorm:"unique_index;not null;type:varchar(50)" json:"mail"`
	Password	string	`gorm:"not null;type:varchar(50)" json:"password"`
}

// models callbacks
func (user *User) BeforeCreate(scope *gorm.Scope) error {
    scope.SetColumn("CreatedAt", time.Now().Unix())
    return nil
}
func (user *User) BeforeUpdate(scope *gorm.Scope) error {
    scope.SetColumn("UpdatedAt", time.Now().Unix())
    return nil
}

func (User) TableName() string { // 自定义表名
	return "user"
}

func ExistUserByMail(mail string) bool {
	var user User
	db.Select("id").Where("mail = ?", mail).First(&user)

	if (user.ID > 0) {
		return true
	}

	return false
}

func AddUser(mail string, pwd string) bool {
	users := User{Mail: mail, Password: pwd}
	db.Create(&users)

	return true
}

func CheckUserMd5Pwd(mail string, md5pwd string) int {
	var user User
	db.Select("id, mail").Where(User{Mail: mail, Password: md5pwd}).First(&user)

	return user.ID // if not exist return 0
}
