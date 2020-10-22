package gredis

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/jameshwc/Million-Singer/conf"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/log"
	"github.com/jameshwc/Million-Singer/repo"
)

type redisRepository struct {
	rdb *redis.Client
}

func Setup() {
	repo.Cache = NewRedisRepository()
}

func NewRedisRepository() repo.CacheRepo {
	rdb := redis.NewClient(&redis.Options{
		Addr:         conf.RedisConfig.Host + ":" + strconv.Itoa(conf.RedisConfig.Port),
		Password:     conf.RedisConfig.Password,
		IdleTimeout:  conf.RedisConfig.IdleTimeout,
		MinIdleConns: conf.RedisConfig.MinIdleConn,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		log.Fatal("redis: error when sending request", err)
	}
	return &redisRepository{rdb: rdb}
}

var setKeyScript = redis.NewScript(`
	redis.call('set', KEYS[1], ARGV[1], 'ex', ARGV[2])
`)
var getKeyScript = redis.NewScript(`
	return redis.call('get', KEYS[1])
`)

func (r *redisRepository) Set(key string, data interface{}, timeout int) error {
	return set(r.rdb, key, data, timeout)
}

func (r *redisRepository) Get(key string) ([]byte, error) {
	return get(r.rdb, key)
}

func (r *redisRepository) Del(key string) error {
	return r.rdb.Del(key).Err()
}

func setWithLua(rdb *redis.Client, key string, data interface{}, timeout int) error {
	value, err := json.Marshal(data)
	if err != nil {
		return C.ErrRedisSetKeyJsonMarshal
	}
	_, err = setKeyScript.Run(rdb, []string{key}, value, timeout).Result()
	return err
}

func getWithLua(rdb *redis.Client, key string) ([]byte, error) {
	dat, err := getKeyScript.Run(rdb, []string{key}, nil).String()
	return []byte(dat), err
}

func set(rdb *redis.Client, key string, data interface{}, timeout int) error {
	value, err := json.Marshal(data)
	if err != nil {
		return C.ErrRedisSetKeyJsonMarshal
	}
	return rdb.Set(key, value, time.Duration(timeout)*time.Second).Err()
}

func get(rdb *redis.Client, key string) ([]byte, error) {
	return rdb.Get(key).Bytes()
}
