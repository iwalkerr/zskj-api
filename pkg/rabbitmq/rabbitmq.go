package rabbitmq

import (
	"fmt"
	"log"
	"sync"

	"github.com/streadway/amqp"
)

//rabbitMQ结构体
type RabbitMQ struct {
	conn    *amqp.Connection
	channel *amqp.Channel
	//队列名称
	QueueName string
	//交换机名称
	Exchange string
	//bind Key 名称
	Key string
	//连接信息
	Mqurl string
	sync.Mutex
}

//创建结构体实例
func NewRabbitMQ(mqUrl, queueName, exchange, key string) *RabbitMQ {
	return &RabbitMQ{QueueName: queueName, Exchange: exchange, Key: key, Mqurl: mqUrl}
}

//断开channel 和 connection
func (r *RabbitMQ) Destory() {
	r.channel.Close()
	r.conn.Close()
}

//错误处理函数
func (r *RabbitMQ) failOnErr(err error, message string) {
	if err != nil {
		log.Fatalf("%s:%s", message, err)
		panic(fmt.Sprintf("%s:%s", message, err))
	}
}

// 连接mq
func conn(mqurl, queueName, exchange, key string) *RabbitMQ {
	//创建RabbitMQ实例
	rabbitmq := NewRabbitMQ(mqurl, queueName, exchange, key)
	var err error
	//获取connection
	rabbitmq.conn, err = amqp.Dial(rabbitmq.Mqurl)
	rabbitmq.failOnErr(err, "failed to connect rabbitmq!")
	//获取channel
	rabbitmq.channel, err = rabbitmq.conn.Channel()
	rabbitmq.failOnErr(err, "failed to open a channel")
	return rabbitmq
}

//队列生产
func (r *RabbitMQ) publish(queueName, message string) error {
	//调用channel 发送消息到队列中
	err := r.channel.Publish(
		r.Exchange,
		r.QueueName,
		//如果为true，根据自身exchange类型和routekey规则无法找到符合条件的队列会把消息返还给发送者
		false,
		//如果为true，当exchange发送消息到队列后发现队列上没有消费者，则会把消息返还给发送者
		false,
		amqp.Publishing{
			ContentType: "text/plain",
			Body:        []byte(message),
		})
	return err
}

//消费者
func (r *RabbitMQ) consume(queueName string, callback func(msgs <-chan amqp.Delivery)) {
	//接收消息
	msgs, err := r.channel.Consume(
		queueName, // queue
		//用来区分多个消费者
		"", // consumer
		//是否自动应答
		true, // auto-ack
		//是否独有
		false, // exclusive
		//设置为true，表示 不能将同一个Conenction中生产者发送的消息传递给这个Connection中 的消费者
		false, // no-local
		//列是否阻塞
		false, // no-wait
		nil,   // args
	)
	if err != nil {
		fmt.Println(err)
	}

	forever := make(chan bool)
	//启用协程处理消息
	go callback(msgs)

	log.Printf(" [*] Waiting for messages. To exit press CTRL+C")
	<-forever

}

//////////////////////////////////////////////////////////////////////////////

// 话题模式
func NewRabbitMQTopic(mqurl, exchangeName, routingKey string) *RabbitMQ {
	return conn(mqurl, "", exchangeName, routingKey)
}

// 话题模式发送消息
func (r *RabbitMQ) PubblishTopic(message string) error {
	// 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic", // 交换机类型，广播类型
		true,
		false,
		//YES表示这个exchange不可以被client用来推送消息，仅用来exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")

	return r.publish(r.Key, message)
}

// 匹配 produc.* ,能够匹配produc.hello 。如果要匹配produc.hello.one 需要produc.*#
func (r *RabbitMQ) ReciveTopic(callback func(msgs <-chan amqp.Delivery)) {
	// 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"topic", // 交换机类型，广播类型
		true,
		false,
		//YES表示这个exchange不可以被client用来推送消息，仅用来exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")

	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an queue")

	// 绑定队列到exchange中
	_ = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下，这儿的key为空
		r.Key,
		r.Exchange,
		false,
		nil,
	)

	r.consume(q.Name, callback)
}

//////////////////////////////////////////////////////////////////////////////

// 路由模式
func NewRabbitMQRouting(mqurl, exchangeName, routingKey string) *RabbitMQ {
	return conn(mqurl, "", exchangeName, routingKey)
}

// 路由模式发送消息
func (r *RabbitMQ) PublishRouting(message string) error {
	// 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct", // 交换机类型，广播类型
		true,
		false,
		//YES表示这个exchange不可以被client用来推送消息，仅用来exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")

	return r.publish(r.Key, message)
}

func (r *RabbitMQ) ReciveRouting(callback func(msgs <-chan amqp.Delivery)) {
	// 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"direct", // 交换机类型，广播类型
		true,
		false,
		//YES表示这个exchange不可以被client用来推送消息，仅用来exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")

	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an queue")

	// 绑定队列到exchange中
	_ = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下，这儿的key为空
		r.Key,
		r.Exchange,
		false,
		nil,
	)

	r.consume(q.Name, callback)
}

//////////////////////////////////////////////////////////////////////////////

// 订阅模式
func NewRabbitMQSub(mqurl, exchangeName string) *RabbitMQ {
	return conn(mqurl, "", exchangeName, "")
}

// 订阅模式下生产
func (r *RabbitMQ) PublishSub(message string) error {
	// 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout", // 交换机类型，广播类型
		true,
		false,
		//YES表示这个exchange不可以被client用来推送消息，仅用来exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")

	return r.publish("", message)
}

// 订阅模式消费端
func (r *RabbitMQ) ReciveSub(callback func(msgs <-chan amqp.Delivery)) {
	// 尝试创建交换机
	err := r.channel.ExchangeDeclare(
		r.Exchange,
		"fanout", // 交换机类型，广播类型
		true,
		false,
		//YES表示这个exchange不可以被client用来推送消息，仅用来exchange和exchange之间的绑定
		false,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an exchange")

	q, err := r.channel.QueueDeclare(
		"", //随机生产队列名
		false,
		false,
		true,
		false,
		nil,
	)
	r.failOnErr(err, "Failed to declare an queue")

	// 绑定队列到exchange中
	_ = r.channel.QueueBind(
		q.Name,
		//在pub/sub模式下，这儿的key为空
		"",
		r.Exchange,
		false,
		nil,
	)

	r.consume(q.Name, callback)
}

//////////////////////////////////////////////////////////////////////////////

// 简单模式
func NewRabbitMQSimple(mqurl, queueName string) *RabbitMQ {
	//创建RabbitMQ实例
	return conn(mqurl, queueName, "", "")
}

func (r *RabbitMQ) PublishSimple(message string) error {
	r.Lock()
	defer r.Unlock()
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	_, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		return err
	}

	return r.publish(r.QueueName, message)
}

//simple 模式下消费者
func (r *RabbitMQ) ConsumeSimple(callback func(msgs <-chan amqp.Delivery)) {
	//1.申请队列，如果队列不存在会自动创建，存在则跳过创建
	q, err := r.channel.QueueDeclare(
		r.QueueName,
		//是否持久化
		false,
		//是否自动删除
		false,
		//是否具有排他性
		false,
		//是否阻塞处理
		false,
		//额外的属性
		nil,
	)
	if err != nil {
		fmt.Println(err)
	}

	//消费者流控
	_ = r.channel.Qos(
		1,     //当前消费者一次能接受的最大消息数量
		0,     //服务器传递的最大容量（以八位字节为单位）
		false, //如果设置为true 对channel可用
	)

	r.consume(q.Name, callback)
}

//////////////////////////////////////////////////////////////////////////////
