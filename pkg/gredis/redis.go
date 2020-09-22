package gredis

import (
	"encoding/json"
	"log"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/jameshwc/Million-Singer/conf"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
)

var rdb *redis.Client

func Setup() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         conf.RedisConfig.Host + ":" + strconv.Itoa(conf.RedisConfig.Port),
		Password:     conf.RedisConfig.Password,
		IdleTimeout:  conf.RedisConfig.IdleTimeout,
		MinIdleConns: conf.RedisConfig.MinIdleConn,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal("redis: error when sending request", err)
	}
}

var setKeyScript = redis.NewScript(`
	redis.call("SET KEY[1] ARGV[1] EX ARGV[2]")
`)

func Set(key string, data interface{}, timeout int) error {
	value, err := json.Marshal(data)
	if err != nil {
		return C.ErrRedisSetKeyJsonMarshal
	}
	// setKeyScript.Run()
	return rdb.Set(key, value, time.Duration(timeout)*time.Second).Err()
}

func Get(key string) ([]byte, error) {
	return rdb.Get(key).Bytes()
}

func Exists(key string) bool {
	if _, err := rdb.Get(key).Result(); err == redis.Nil {
		return false
	}
	return true
}
