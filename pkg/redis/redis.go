package redis

import (
	"github.com/go-redis/redis"
	"github.com/jameshwc/Million-Singer/conf"
)

var rdb *redis.Client

func Setup() {
	rdb = redis.NewClient(&redis.Options{
		Addr:         conf.RedisConfig.Host,
		Password:     conf.RedisConfig.Password,
		IdleTimeout:  conf.RedisConfig.IdleTimeout,
		MinIdleConns: conf.RedisConfig.MinIdleConn,
	})
}
