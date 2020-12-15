package cache

import (
	"sync"

	"github.com/go-redis/redis"
)

var (
	instanceRedis *redis.Client
	onceRedis     sync.Once
)

func Redis(host, pwd string) *redis.Client {
	onceRedis.Do(func() {
		instanceRedis = redis.NewClient(&redis.Options{
			Addr:     host,
			Password: pwd,
		})
	})
	return instanceRedis
}
