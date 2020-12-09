package main

import (
	"fmt"
	"xframe/pkg/server"
	"xframe/validate/common"
	"xframe/validate/config"
	_ "xframe/validate/core"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var ag errgroup.Group

// API接口服务
func main() {
	// 存储服务器信息，分布到hash环上
	common.AddHostInfo(config.HostArray)

	gin.SetMode(gin.ReleaseMode)
	// API服务
	api := server.New("validate", config.ServerPort, gin.Recovery())
	api.Start(&ag)

	if err := ag.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}
