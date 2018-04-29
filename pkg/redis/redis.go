package redis

import (
	"fmt"
	"time"
	"os"
	"github.com/go-redis/redis"
	"github.com/GallenHu/bookmarkgo/pkg/setting"
	"github.com/GallenHu/bookmarkgo/model"
	"github.com/GallenHu/bookmarkgo/pkg/utils"
)

var RedisHost = setting.RedisHost
var RedisPwd = setting.RedisPwd
var RedisDb = setting.RedisDb

var client *redis.Client

func init() {
	client = redis.NewClient(&redis.Options{
		Addr:     RedisHost,
		Password: RedisPwd,
		DB:       RedisDb,  // default 0
		DialTimeout:  10 * time.Second,
		ReadTimeout:  30 * time.Second,
		WriteTimeout: 30 * time.Second,
		PoolSize:     10,
		PoolTimeout:  30 * time.Second,
	})
	// client.FlushDB() // clean
}

func TestConnect() {
	pong, err := client.Ping().Result()
	if err != nil {
		fmt.Println("Error on connect to redis")
		fmt.Println(pong, err)
		os.Exit(1)
	}
}

func GetVal(key string) string {
	val, err := client.Get(key).Result()
	if err != nil {
		fmt.Println(err)
		return ""
	}
	return val
}

func SetVal(key string, val interface {}, exphours int) error {
	exp := time.Duration(exphours) * time.Hour
	err := client.Set(key, val, exp).Err()
	return err
}

func SetExpiration(key string, exphours int) error {
	exp := time.Duration(exphours) * time.Hour
	err := client.Expire(key, exp).Err()
	return err
}

func DelVal(key string) bool {
	client.Del(key)
	return true
}

func StoreUserToken(userid int, token string) error {
	return SetVal("user:" + utils.Int2str(userid), token, setting.AppTokenExpire)
}

func GetUserToken(userid int) string {
	return GetVal("user:" + utils.Int2str(userid))
}

func ExtendUserTokenExpire(userid int) error {
	return SetExpiration("user:" + utils.Int2str(userid), setting.AppTokenExpire)
}

func DelUserToken(userid int) bool {
	return DelVal("user:" + utils.Int2str(userid))
}

func StoreUserPrivate(userid int, showprivate uint) error {
	return SetVal("userprivate:" + utils.Int2str(userid), showprivate, 0)
}

func GetUserPrivate(userid int) uint {
	val, err := client.Get("userprivate:" + utils.Int2str(userid)).Result() // 没有值时 err != nil
	if err != nil {
		user, err := model.GetUserById(userid)
		if err != nil {
			return 0
		} else {
			return user.ShowPrivate
		}
	} else {
		valofint := utils.Str2int(val, 0)
		return uint(valofint)
	}
}
