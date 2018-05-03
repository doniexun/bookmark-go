package redis

import (
	"fmt"
	"time"
	"os"
	"os/signal"
	"syscall"
	"github.com/garyburd/redigo/redis"
	"github.com/GallenHu/bookmarkgo/pkg/setting"
	"github.com/GallenHu/bookmarkgo/model"
	"github.com/GallenHu/bookmarkgo/pkg/utils"
)

var RedisHost = setting.RedisHost
var RedisPwd = setting.RedisPwd
var RedisDb = setting.RedisDb

var Pool *redis.Pool

func init() {
	Pool = newPool(RedisHost)
	cleanupHook()
}

func newPool(host string) *redis.Pool {
	return &redis.Pool{
		MaxIdle:     3,
		IdleTimeout: 60 * time.Second,

		Dial: func() (redis.Conn, error) {
			connect, err := redis.Dial("tcp", host)
			if err != nil {
				return nil, err
			}
			return connect, err
		},

		TestOnBorrow: func(c redis.Conn, t time.Time) error {
            _, err := c.Do("PING")
            return err
        },
	}
}

func cleanupHook() {
	c := make(chan os.Signal, 1)
    signal.Notify(c, os.Interrupt)
    signal.Notify(c, syscall.SIGTERM)
    signal.Notify(c, syscall.SIGKILL)
    go func() {
        <-c
        Pool.Close()
        os.Exit(0)
    }()
}

func TestConnect() {
	conn := Pool.Get()
	defer conn.Close()

	_, err := redis.String(conn.Do("PING"))
    if err != nil {
		fmt.Println("Error on connect to redis")
		fmt.Println("%v", err)
        os.Exit(1)
    }
}

func GetStrVal(key string) string {
	conn := Pool.Get()
	defer conn.Close()

	data, err := redis.String(conn.Do("GET", key))
    if err != nil {
        fmt.Errorf("error getting key %s: %v", key, err)
    }
    return data
}

func SetVal(key string, value interface {}) error {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("SET", key, value)
	if err != nil {
        fmt.Errorf("error setting key %s to %s: %v", key, value.(string), err)
    }
    return err
}

func DelVal(key string) error {
	conn := Pool.Get()
    defer conn.Close()

    _, err := conn.Do("DEL", key)
    return err
}

func SetExpire(key string, hours int) error {
	conn := Pool.Get()
	defer conn.Close()

	_, err := conn.Do("EXPIRE", key, hours * 3600)
    return err
}

func GetUserToken(userid int) string {
	return GetStrVal("user:" + utils.Int2str(userid))
}

func StoreUserToken(userid int, token string) error {
	return SetVal("user:" + utils.Int2str(userid), token)
}

func DelUserToken(userid int) error {
	return DelVal("user:" + utils.Int2str(userid))
}

func SetUserTokenExpire(userid int) error {
	return SetExpire("user:" + utils.Int2str(userid), setting.AppTokenExpire)
}

func StoreUserPrivate(userid int, showprivate uint) error {
	return SetVal("userprivate:" + utils.Int2str(userid), showprivate)
}

func GetUserPrivate(userid int) uint {
	valstr := GetStrVal("userprivate:" + utils.Int2str(userid))
	if valstr == "" {
		user, err := model.GetUserById(userid)
		if err != nil {
			return 0
		} else {
			return user.ShowPrivate
		}
	}
	valint := utils.Str2int(valstr, 0)
	return uint(valint)
}

