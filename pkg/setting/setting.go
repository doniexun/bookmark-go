package setting

import (
	"fmt"
	"log"
	"os"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	AppMode string
	AppPort string
	AppSecret string

	DbType string
	DbHost string
	DbName string
	DbUser string
	DbPwd string

	RedisHost string
	RedisPwd string
	RedisDb int
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
	AppMode = Cfg.Section("app").Key("RUN_MODE").MustString("debug")
	AppPort = Cfg.Section("app").Key("PORT").MustString("3001")
	AppSecret = Cfg.Section("app").Key("SECRET").MustString("")

	DbType = Cfg.Section("database").Key("TYPE").MustString("mysql")
	DbHost = Cfg.Section("database").Key("HOST").MustString("127.0.0.1:3306")
	DbName = Cfg.Section("database").Key("NAME").MustString("bookmark")
	DbUser = Cfg.Section("database").Key("USER").MustString("root")
	DbPwd = Cfg.Section("database").Key("PASSWORD").MustString("123456")

	RedisHost = Cfg.Section("redis").Key("HOST").MustString("127.0.0.1:6379")
	RedisPwd = Cfg.Section("redis").Key("PASSWORD").MustString("")
	RedisDb = Cfg.Section("redis").Key("DB").MustInt(0)
}
