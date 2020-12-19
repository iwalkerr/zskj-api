package cfg

import (
	"log"
	"sync"

	"github.com/BurntSushi/toml"
)

// 配置文件路径
var filePath = "config/config.toml"

var (
	instance *config
	once     sync.Once
)

//获取配置文档实例
func Instance() *config {
	once.Do(func() {
		if _, err := toml.DecodeFile(filePath, &instance); err != nil {
			log.Fatal("配置文件参数错误", err.Error())
		}
	})
	return instance
}

type config struct {
	Admin    admin
	Database database
	Redis    redis
}

type redis struct {
	Host     string
	Password string
}

type admin struct {
	Address      string
	WebsocketIp  string
	ServerRoot   string
	RedisStore   bool
	BusinessName string
	CompanyName  string
}

type database struct {
	Master string
	Slave  string
}
