package redis

import (
	"xframe/backend/common/cfg"
	"xframe/pkg/cache"

	"github.com/go-redis/redis"
)

func Conn() *redis.Client {
	r := cfg.Instance().Redis
	return cache.Redis(r.Host, r.Password)
}
