package main

import (
	"log"
	"xframe/frontend/config"
	_ "xframe/frontend/core"
	_ "xframe/frontend/public/swagger"
	"xframe/pkg/server"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func main() {
	gin.SetMode(gin.ReleaseMode)

	// API服务
	api := server.New("api", config.ServerPort, gin.Recovery(), gin.Logger())
	api.Static("public")
	api.Start(&g)

	if err := g.Wait(); err != nil {
		log.Println(err.Error())
	}
}
