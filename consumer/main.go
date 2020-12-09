package main

import (
	"fmt"
	"log"
	"xframe/frontend/config"
	"xframe/pkg/rabbitmq"
	"xframe/pkg/server"

	"github.com/gin-gonic/gin"
	"github.com/streadway/amqp"
	"golang.org/x/sync/errgroup"
)

var g errgroup.Group

// rabbitmq消费信息
func main() {
	// 监听消息队列
	rabbitmq.NewRabbitMQSimple(config.MQURL, "zs_product").ConsumeSimple(consumeSimple)

	gin.SetMode(gin.ReleaseMode)
	// API服务
	api := server.New("consumer", ":8085", gin.Recovery())
	api.Start(&g)

	if err := g.Wait(); err != nil {
		fmt.Println(err.Error())
	}
}

// 消息队列处理业务逻辑
func consumeSimple(msgs <-chan amqp.Delivery) {
	for d := range msgs {
		//消息逻辑处理，可以自行设计逻辑
		log.Printf("Received a message: %s", d.Body)
	}
}
