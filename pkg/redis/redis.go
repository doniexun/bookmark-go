package redis

import (
	"log"
	"os"
	"github.com/mediocregopher/radix.v2/redis"
	"github.com/GallenHu/bookmarkgo/pkg/setting"
	// "github.com/GallenHu/bookmarkgo/model"
	"github.com/GallenHu/bookmarkgo/pkg/utils"
)

var RedisHost = setting.RedisHost
var RedisPwd = setting.RedisPwd
var RedisDb = setting.RedisDb

var client *redis.Client

func init() {
	client, err := redis.Dial("tcp", RedisHost)
	if err != nil {
		log.Println("Error on connect to redis")
		log.Println(err)
		os.Exit(1)
	}

	client.PipeClear()
}

func GetStrVal(key string) string {
	val, err := client.Cmd("GET", key).Str()
	if err != nil {
		log.Println(err)
		return ""
	}
	return val
}

func SetVal(key string, val interface {}, exphours int) error {
	log.Println("11112222")
	err := client.Cmd("SET", key, val, "EX", exphours * 3600).Err
	if err != nil {
		return err
	}
	return nil
}

func DelVal(key string) bool {
	client.Cmd("DEL", key)
	return true
}

func StoreUserToken(userid int, token string) error {
	return SetVal("user:" + utils.Int2str(userid), token, setting.AppTokenExpire)
}

func GetUserToken(userid int) string {
	return GetStrVal("user:" + utils.Int2str(userid))
}

func DelUserToken(userid int) bool {
	return DelVal("user:" + utils.Int2str(userid))
}

func StoreUserPrivate(userid int, showprivate uint) error {
	return SetVal("userprivate:" + utils.Int2str(userid), showprivate, setting.AppTokenExpire)
}

func GetUserPrivate(userid int) uint {
	val, err := client.Cmd("GET", "userprivate:" + utils.Int2str(userid)).Int()
	if err != nil {
		return 0
	}

	return uint(val)
}
