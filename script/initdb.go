// This file help you create database;
// Run this script only on develop env!!!
package main

import "fmt"
import "github.com/GallenHu/bookmarkgo/model"
import "github.com/jinzhu/gorm"
import _ "github.com/go-sql-driver/mysql"
import _ "github.com/jinzhu/gorm/dialects/mysql"

func main() {
	db, err := gorm.Open("mysql", "root:123456@(127.0.0.1:3306)/bookmark?charset=utf8&parseTime=True&loc=Local")

	defer db.Close()
	if err == nil {
		//db.LogMode(true)
		db.DropTableIfExists(&model.User{}, &model.Bookmark{})
		db.AutoMigrate(&model.User{}, &model.Bookmark{})
		// db.Model(&PostTag{}).AddUniqueIndex("uk_post_tag", "post_id", "tag_id")
		fmt.Println("DB init success!")
		return
	}
	fmt.Println("[error]", err);
}
