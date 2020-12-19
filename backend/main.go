package main

import (
	"fmt"
	"os"
	"os/signal"
	"sync/atomic"
	"syscall"
	"time"
	"xframe/backend/core/middleware/sessions"

	"xframe/backend/common/cfg"
	"xframe/backend/common/server"
	"xframe/backend/core/middleware/sessions/memstore"
	"xframe/backend/core/middleware/sessions/redis"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"

	// 初始化路由配置
	_ "xframe/backend/app"
	_ "xframe/backend/core"
)

var g errgroup.Group

// 管理后台系统
func main() {
	gin.SetMode(gin.ReleaseMode)
	config := cfg.Instance()

	// 后台管理
	var store sessions.Store
	if config.Admin.RedisStore {
		store, _ = redis.NewStore(10, "tcp", config.Redis.Host, config.Redis.Password, []byte("secret"))
	} else {
		store = memstore.NewStore([]byte("secret"))
	}
	store.Options(sessions.Options{
		Path:     "/",
		HttpOnly: true,
		MaxAge:   3600 * 24 * 30, // 过期时间30天
	})
	admin := server.New("admin", config.Admin.Address, gin.Logger(), sessions.Sessions("sessionid", store))
	admin.Template("template").Static(config.Admin.ServerRoot)
	admin.Start(&g)

	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}

	var state int32 = 1
	sc := make(chan os.Signal, 1)
	signal.Notify(sc, syscall.SIGHUP, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT)

EXIT:
	for {
		sig := <-sc
		fmt.Printf("获取到信号[%s]", sig.String())
		switch sig {
		case syscall.SIGQUIT, syscall.SIGTERM, syscall.SIGINT:
			atomic.StoreInt32(&state, 0)
			break EXIT
		case syscall.SIGHUP:
		default:
			break EXIT
		}
	}

	fmt.Println("服务退出")
	time.Sleep(time.Second)
	os.Exit(int(atomic.LoadInt32(&state)))
}
