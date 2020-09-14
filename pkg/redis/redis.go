package redis

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

func Set(key string, data interface{}, timeout int) error {
	value, err := json.Marshal(data)
	if err != nil {
		return C.ErrRedisSetKeyJsonMarshal
	}
	return rdb.Set(key, value, time.Duration(timeout)*time.Second).Err()
}

func Get(key string) ([]byte, error) {
	return rdb.Get(key).Bytes()
}
