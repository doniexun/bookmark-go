package setting

import (
	"fmt"
	"log"
	"os"
	"flag"
	"github.com/go-ini/ini"
)

var (
	Cfg *ini.File

	AppMode string
	AppPort string
	AppSecret string
	AppTokenExpire int
	AppCors string
	AllowSearch int

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
	configFilePath := flag.String("c", "conf/app.ini", "config file path")
	flag.Parse()
	log.Println("[config file path]: ", *configFilePath)
	Cfg, err = ini.Load(*configFilePath) // 相对根目录的路径
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
	AppTokenExpire = Cfg.Section("app").Key("TOKEN_EXPIRE").MustInt(72)
	AppCors = Cfg.Section("app").Key("CORS").MustString("http://localhost:8080")
	AllowSearch = Cfg.Section("app").Key("ALLOW_SEARCH").MustInt(1)

	DbType = Cfg.Section("database").Key("TYPE").MustString("mysql")
	DbHost = Cfg.Section("database").Key("HOST").MustString("127.0.0.1:3306")
	DbName = Cfg.Section("database").Key("NAME").MustString("bookmark")
	DbUser = Cfg.Section("database").Key("USER").MustString("root")
	DbPwd = Cfg.Section("database").Key("PASSWORD").MustString("123456")

	RedisHost = Cfg.Section("redis").Key("HOST").MustString("127.0.0.1:6379")
	RedisPwd = Cfg.Section("redis").Key("PASSWORD").MustString("")
	RedisDb = Cfg.Section("redis").Key("DB").MustInt(0)
}
