package config

//连接信息
const MQURL = "amqp://guest:guest@127.0.0.1:5672/zw_product_host"

// 服务端口
const ServerPort = ":8086"

// 主数据库
// 从数据库 如果为空则表示只使用主数据
const MysqlMaster = "root:123456@tcp(127.0.0.1:3306)/zskj_api?charset=utf8mb4&parseTime=true&loc=Local"
const MysqlSlave = ""

// redis
const RedisHost = "127.0.0.1:6379"
const RedisPwd = "123456"
