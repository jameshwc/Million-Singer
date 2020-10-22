package gredis

import (
	"testing"

	"github.com/jameshwc/Million-Singer/conf"
)

var testRdb *redisRepository

func testRedisSetup() {
	conf.RedisConfig = &conf.Redis{
		Host: "127.0.0.1",
		Port: 6379,
	}
	testRdb = &redisRepository{}
}
func BenchmarkRedisNoLua(b *testing.B) {
	testRedisSetup()
	for i := 0; i < b.N; i++ {
		set(testRdb.rdb, "hello", conf.RedisConfig, 100)
		get(testRdb.rdb, "hello")
	}
}

func BenchmarkRedisLua(b *testing.B) {
	testRedisSetup()
	for i := 0; i < b.N; i++ {
		setWithLua(testRdb.rdb, "hello", conf.RedisConfig, 100)
		getWithLua(testRdb.rdb, "hello")
	}
}
