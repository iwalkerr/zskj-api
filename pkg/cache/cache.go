package cache

import (
	"sync"
	"time"

	"github.com/patrickmn/go-cache"
)

var (
	instance *cache.Cache
	once     sync.Once
)

//获取缓存单例,默认失效时间为5分钟，10分钟清理一次
func Instance() *cache.Cache {
	once.Do(func() {
		instance = cache.New(5*time.Minute, 10*time.Minute)
	})

	return instance
}
