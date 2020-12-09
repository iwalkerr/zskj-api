package config

//连接信息
const MQURL = "amqp://guest:guest@127.0.0.1:5672/zw_product_host"
const MQQueueName = "zs_product"

// 设置集群地址，最好内网IP
var HostArray = []string{"192.168.0.102"}

// 服务端口
const ServerPort = ":8082"
