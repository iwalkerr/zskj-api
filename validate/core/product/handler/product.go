package handler

import (
	"encoding/json"
	"xframe/pkg/rabbitmq"
	"xframe/validate/common"
	"xframe/validate/config"

	"github.com/gin-gonic/gin"
)

func CheckRight(c *gin.Context) {
	// 获取分布式权限验证
	if ok := common.AccessControll.GetDdistributedRight(c); !ok {
		c.String(200, "false")
	} else {
		c.String(200, "true")
	}
}

func List(c *gin.Context) {
	productId := 1
	// 从cookie中获取
	userId := 1

	// 获取分布式权限验证
	if ok := common.AccessControll.GetDdistributedRight(c); !ok {
		c.String(200, "false")
		return
	}

	// 获取数量控制，防止秒杀出现超卖
	hostUrl := "http://" + common.GetOneIp + ":" + common.GetOnePort + "/getone/v1/product"
	responseValidate, body, err := common.GetCurl(hostUrl, c)
	if err != nil {
		c.String(200, "false")
		return
	}

	// 判断数量控制接口请求状态
	if responseValidate.StatusCode == 200 {
		if string(body) == "true" {
			// 整合下单
			message := common.NewMessage(userId, productId)
			byteMsg, err := json.Marshal(message)
			if err != nil {
				c.String(200, "false")
				return
			}

			if err := rabbitmq.NewRabbitMQSimple(config.MQURL, config.MQQueueName).PublishSimple(string(byteMsg)); err != nil {
				c.String(200, "false")
				return
			} else {
				c.String(200, "true")
				return
			}
		}
	}

	c.String(200, "false")
}
