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
	testRdb = NewRedisRepository()
}
func BenchmarkRedisNoLua(b *testing.B) {
	testRedisSetup()
	for i := 0; i < b.N; i++ {
		testRdb.set("hello", conf.RedisConfig, 100)
		testRdb.get("hello")
	}
}

func BenchmarkRedisLua(b *testing.B) {
	testRedisSetup()
	for i := 0; i < b.N; i++ {
		testRdb.setWithLua("hello", conf.RedisConfig, 100)
		testRdb.getWithLua("hello")
	}
}
