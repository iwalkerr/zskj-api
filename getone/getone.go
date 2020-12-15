package main

import (
	"fmt"
	"strconv"
	"sync"
	"xframe/getone/config"
	"xframe/pkg/router"
	"xframe/pkg/server"

	"github.com/gin-gonic/gin"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

func init() {
	g1 := router.New("getone", "/getone/v1/product")
	g1.GET("/", "", getProduct)
}

// API接口服务
func main() {
	gin.SetMode(gin.ReleaseMode)
	// API服务
	api := server.New("getone", config.ServerPort, gin.Recovery())
	api.Start(&g)

	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}

func getProduct(c *gin.Context) {
	c.String(200, strconv.FormatBool(GetOneProduct()))
}

var sum = 0

// 预存商品数量
var productNum = 1000000
var mutex sync.Mutex

// 计数
var count = 0

// 获取秒杀商品
func GetOneProduct() bool {
	mutex.Lock()
	defer mutex.Unlock()

	count++

	// 判断数据是否超限
	if count%100 == 0 {
		if sum < productNum {
			sum += 1
			fmt.Println(sum)
			return true
		}
	}
	return false
}
