package main

import (
	"fmt"
	_ "xframe/frontend/core"
	_ "xframe/frontend/public/swagger"
	"xframe/pkg/server"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	gin.SetMode(gin.DebugMode)

	// API服务
	api := server.New("api", ":8086", gin.Recovery(), gin.Logger())
	api.Static("public")
	api.Start(&g)

	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}
