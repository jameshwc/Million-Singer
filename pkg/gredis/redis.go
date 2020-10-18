package gredis

import (
	"encoding/json"
	"strconv"
	"time"

	"github.com/go-redis/redis"
	"github.com/jameshwc/Million-Singer/conf"
	C "github.com/jameshwc/Million-Singer/pkg/constant"
	"github.com/jameshwc/Million-Singer/pkg/log"
)

type redisRepository struct {
	rdb *redis.Client
}

func NewRedisRepository() *redisRepository {
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
	return r.set(key, data, timeout)
}

func (r *redisRepository) Get(key string) ([]byte, error) {
	return r.get(key)
}

func (r *redisRepository) Del(key string) error {
	return r.rdb.Del(key).Err()
}

func (r *redisRepository) setWithLua(key string, data interface{}, timeout int) error {
	value, err := json.Marshal(data)
	if err != nil {
		return C.ErrRedisSetKeyJsonMarshal
	}
	_, err = setKeyScript.Run(r.rdb, []string{key}, value, timeout).Result()
	return err
}

func (r *redisRepository) getWithLua(key string) ([]byte, error) {
	dat, err := getKeyScript.Run(r.rdb, []string{key}, nil).String()
	return []byte(dat), err
}

func (r *redisRepository) set(key string, data interface{}, timeout int) error {
	value, err := json.Marshal(data)
	if err != nil {
		return C.ErrRedisSetKeyJsonMarshal
	}
	return r.rdb.Set(key, value, time.Duration(timeout)*time.Second).Err()
}

func (r *redisRepository) get(key string) ([]byte, error) {
	return r.rdb.Get(key).Bytes()
}
