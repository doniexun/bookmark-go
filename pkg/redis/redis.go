package redis

import (
	"fmt"
	"time"
	"os"
	"github.com/go-redis/redis"
	"github.com/GallenHu/bookmarkgo/pkg/setting"
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

func SetVal(key string, val string, exphours int) error {
	exp := time.Duration(exphours) * time.Hour
	err := client.Set(key, val, exp).Err()
	return err
}

func SetExpiration(key string, exphours int) error {
	exp := time.Duration(exphours) * time.Hour
	err := client.Expire(key, exp).Err()
	return err
}
