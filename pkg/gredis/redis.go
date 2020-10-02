package gredis

import (
	"encoding/json"

	"github.com/jameshwc/Million-Singer/pkg/log"

	"strconv"

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
	redis.call('set', KEYS[1], ARGV[1], 'ex', ARGV[2])
`)
var getKeyScript = redis.NewScript(`
	return redis.call('get', KEYS[1])
`)

func Set(key string, data interface{}, timeout int) error {
	value, err := json.Marshal(data)
	if err != nil {
		return C.ErrRedisSetKeyJsonMarshal
	}
	_, err = setKeyScript.Run(rdb, []string{key}, value, timeout).Result()
	return err
}

func Get(key string) ([]byte, error) {
	dat, err := getKeyScript.Run(rdb, []string{key}, nil).String()
	return []byte(dat), err
}
