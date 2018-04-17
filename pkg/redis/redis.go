package redis

import (
	"fmt"
	"github.com/go-redis/redis"
	"github.com/GallenHu/bookmarkgo/pkg/setting"
)

var RedisHost = setting.RedisHost
var RedisPwd = setting.RedisPwd
var RedisDb = setting.RedisDb

func SetVal() {
	fmt.Println(RedisHost)
	client := redis.NewClient(&redis.Options{
		Addr:     RedisHost,
		Password: RedisPwd,
		DB:       RedisDb,  // default 0
	})
	val, err := client.Get("test2").Result()
	fmt.Println(val)
	fmt.Println(err)
}
