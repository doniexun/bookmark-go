package main

import (
	"fmt"
	"log"
	"github.com/jinzhu/gorm"
	"github.com/GallenHu/bookmarkgo/model"
	"github.com/GallenHu/bookmarkgo/pkg/setting"
	_ "github.com/go-sql-driver/mysql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	dbType := setting.DbType
	dbHost := setting.DbHost
	dbName := setting.DbName
	dbUser := setting.DbUser
	dbPwd := setting.DbPwd

	db, err := gorm.Open(dbType, fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		dbUser,
		dbPwd,
		dbHost,
		dbName,
	))
	defer db.Close()

	if err != nil {
		log.Println(err)
	}

	db.DropTableIfExists(&model.User{}, &model.Bookmark{}, &model.Folder{})
	db.AutoMigrate(&model.User{}, &model.Bookmark{}, &model.Folder{})

	fmt.Println("success")
}
