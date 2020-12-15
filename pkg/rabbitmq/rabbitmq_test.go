package rabbitmq

import (
	"fmt"
	"log"
	"strconv"
	"testing"
	"time"

	"github.com/streadway/amqp"
)

const MQURL = "amqp://guest:guest@127.0.0.1:5672/zw_product_host"

// 简单模式
// 工作模式 --- 就是开多个消费者
// 订阅模式
// 话题模式

// 订阅模式
func TestRabitMQSub(t *testing.T) {
	rabbitmq := NewRabbitMQSub(MQURL, "newProduct")

	for i := 0; i < 100; i++ {
		rabbitmq.PublishSub("订阅模式生产第" + strconv.Itoa(i) + "条数据")
		time.Sleep(time.Second)
	}
}

func TestRecieveSub(t *testing.T) {
	rabbitmq := NewRabbitMQSub(MQURL, "newProduct")
	rabbitmq.ReciveSub(func(msgs <-chan amqp.Delivery) {
		for d := range msgs {
			log.Printf("recevied a message: %s", d.Body)
		}
	})
}

// 路由模式
func TestRabbitMQRouting(t *testing.T) {
	one := NewRabbitMQRouting(MQURL, "eexIproduct", "one")
	two := NewRabbitMQRouting(MQURL, "eexIproduct", "two")

	for i := 0; i < 10; i++ {
		one.PublishRouting("hello one" + strconv.Itoa(i))
		two.PublishRouting("hello two" + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func TestRabbitMQReviceRouting(t *testing.T) {
	one := NewRabbitMQRouting(MQURL, "eexIproduct", "one")
	one.ReciveRouting(func(msgs <-chan amqp.Delivery) {
		for d := range msgs {
			log.Printf("recevied a message: %s", d.Body)
		}
	})
}

// 话题
func TestRabbitMQTopic(t *testing.T) {
	one := NewRabbitMQTopic(MQURL, "eexIproducttopic", "one")
	two := NewRabbitMQTopic(MQURL, "eexIproducttopic", "two")

	for i := 0; i < 10; i++ {
		one.PubblishTopic("hello one" + strconv.Itoa(i))
		two.PubblishTopic("hello two" + strconv.Itoa(i))
		time.Sleep(time.Second)
		fmt.Println(i)
	}
}

func TestRabbitMQReviceTopic(t *testing.T) {
	// # 匹配所有
	one := NewRabbitMQTopic(MQURL, "eexIproducttopic", "#")
	one.ReciveTopic(func(msgs <-chan amqp.Delivery) {
		for d := range msgs {
			log.Printf("recevied a message: %s", d.Body)
		}
	})
}
