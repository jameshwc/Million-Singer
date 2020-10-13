package gredis

import (
	"testing"

	"github.com/jameshwc/Million-Singer/conf"
)

func testRedisSetup() {
	conf.RedisConfig = &conf.Redis{
		Host: "127.0.0.1",
		Port: 6379,
	}
	Setup()
}
func BenchmarkRedisNoLua(b *testing.B) {
	testRedisSetup()
	for i := 0; i < b.N; i++ {
		set("hello", conf.RedisConfig, 100)
		get("hello")
	}
}

func BenchmarkRedisLua(b *testing.B) {
	testRedisSetup()
	for i := 0; i < b.N; i++ {
		setWithLua("hello", conf.RedisConfig, 100)
		getWithLua("hello")
	}
}
