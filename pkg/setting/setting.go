package setting

import (
	"github.com/timest/env"
)

var (
	// Cfg *ini.File

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

type config struct {
	AppMode		string	`default:"debug"`
	AppPort		string	`default:"3001"`
	AppSecret	string	`default:"123456"`
	AppTokenExpire int	`default:"72"`
	AppCors		string	`default:"http://localhost:8080"`
	AllowSearch	int		`default:"1"`

	DbType		string	`default:"mysql"`
	DbHost		string	`default:"localhost"`
	DbName		string	`default:"db_bookmark"`
	DbUser		string	`default:"root"`
	DbPwd		string	`default:"123456"`

	RedisHost	string	`default:"127.0.0.1"`
	RedisPwd	string	`default:""`
	RedisDb		int		`default:"0"`
}

func init() {
	cfg := new(config)
	err := env.Fill(cfg)
	if err != nil {
		panic(err)
	}
	// fmt.Println("AppMode:", cfg.AppMode)
	// fmt.Println("AppPort:", cfg.AppPort)

	AppMode = cfg.AppMode
	AppPort = cfg.AppPort
	AppSecret = cfg.AppSecret
	AppTokenExpire = cfg.AppTokenExpire
	AppCors = cfg.AppCors
	AllowSearch = cfg.AllowSearch

	DbType = cfg.DbType
	DbHost = cfg.DbHost + ":3306"
	DbName = cfg.DbName
	DbUser = cfg.DbUser
	DbPwd = cfg.DbPwd

	RedisHost = cfg.RedisHost + ":6379"
	RedisPwd = cfg.RedisPwd
	RedisDb = cfg.RedisDb
}
