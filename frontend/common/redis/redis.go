package redis

import (
	"xframe/frontend/config"
	"xframe/pkg/cache"

	"github.com/go-redis/redis"
)

func Conn() *redis.Client {
	return cache.Redis(config.RedisHost, config.RedisPwd)
}
