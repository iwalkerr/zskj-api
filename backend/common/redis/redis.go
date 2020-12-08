package cache

import (
	"sync"
	"xframe/backend/common/cfg"

	"github.com/go-redis/redis"
)

var (
	instanceRedis *redis.Client
	onceRedis     sync.Once
)

func Redis() *redis.Client {
	onceRedis.Do(func() {
		r := cfg.Instance().Redis
		instanceRedis = redis.NewClient(&redis.Options{
			Addr:     r.Host,
			Password: r.Password,
		})
	})
	return instanceRedis
}
