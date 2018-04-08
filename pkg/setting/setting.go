package setting

import (
	"fmt"
	"log"
	"os"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	RunMode string

	DbType string
	DbHost string
	DbName string
	DbUser string
	DbPwd string
)

func init() {
	var err error
	Cfg, err = ini.Load("conf/app.ini") // 相对根目录的路径
	if err != nil {
		fmt.Println("failed load ini")
		log.Fatalf("load ini error %v", err)
		os.Exit(1)
	}

	LoadSetting()
}

func LoadSetting() {
	RunMode = Cfg.Section("").Key("RUN_MODE").MustString("debug")
	DbType = Cfg.Section("database").Key("TYPE").MustString("mysql")
	DbHost = Cfg.Section("database").Key("HOST").MustString("127.0.0.1:3306")
	DbName = Cfg.Section("database").Key("NAME").MustString("bookmark")
	DbUser = Cfg.Section("database").Key("USER").MustString("root")
	DbPwd = Cfg.Section("database").Key("PASSWORD").MustString("123456")
}
