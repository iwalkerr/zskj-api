package config

// 服务端口
const ServerPort = ":8086"

// 主数据库
// 从数据库 如果为空则表示只使用主数据
const (
	MysqlMaster = "root:123456@tcp(127.0.0.1:3306)/zskj_api?charset=utf8mb4&parseTime=true&loc=Local"
	MysqlSlave  = ""
)

//连接信息
const MQURL = "amqp://guest:guest@127.0.0.1:5672/zw_product_host"

// redis
const (
	RedisHost = "127.0.0.1:6379"
	RedisPwd  = "123456"
)

// jwt 加密字符串
const (
	JwtEncryptKey = "djhuiflrf8832cfvldsbkbsdwqwla1wshk3999sd"
	OutTime       = 3600 * 24 * 30 // 过期时间30天
	RefreshTime   = 1              // 刷新时间
)

// app与服务器较密文件
const (
	PrivateKeyPath = "res/keys/privateKey.pem"
	PublicKeyPath  = "res/keys/publicKey.pem"
)

// 用户加盐
const UserSalt = "XXffff4423$$77%%&7jjj"
