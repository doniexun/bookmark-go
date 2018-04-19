package model

import (
	"fmt"
	"log"
	"os"
	"github.com/jinzhu/gorm"
	"github.com/GallenHu/bookmarkgo/pkg/setting"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var db *gorm.DB

type Model struct {
	ID int `gorm:"primary_key" json:"id"`
    CreatedAt int `json:"created_at"`
    UpdatedAt int `json:"updated_at"`
    DeletedOn int `json:"deleted_on"` // not deleted_at
}

func init() {
	var err error

	dbType := setting.DbType
	dbHost := setting.DbHost
	dbName := setting.DbName
	dbUser := setting.DbUser
	dbPwd := setting.DbPwd

	db, err = gorm.Open(dbType, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", // 注意!!!不要使用 :=
		dbUser,
		dbPwd,
		dbHost,
		dbName,
	))

	if err != nil {
		fmt.Println("gorm open database error")
		log.Println(err)
		os.Exit(1)
	}

	db.LogMode(true)
	db.SingularTable(true)
    db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)
}

func CloseDB() {
	defer db.Close()
}
