package ws

import (
	"net/http"
	"sync"

	"github.com/gorilla/websocket"
)

// 约定定义
const (
	OnSysUserGrpId = "backend-110" // 后台系统在线用户组
)

// ws连接协议
var upGrader = &websocket.Upgrader{
	WriteBufferSize: 10240,
	ReadBufferSize:  10240,
	CheckOrigin: func(r *http.Request) bool { // 允许跨域
		return true
	},
}

// 对外维护socket信息
var connPool = clientManager{
	broadcast:  make(chan []byte, 1000),
	register:   make(chan *client, 500),
	unregister: make(chan *client, 500),
}

// 接收消息
type receiveMsg struct {
	FromUser    string `json:"from_user"`    // 发送者
	ToUser      string `json:"to_user"`      // 接收者
	GrpId       string `json:"grp_id"`       // 组或群ID: 110默认是后台在线用户
	Content     string `json:"content"`      // 传输内容
	ContentType int    `json:"content_type"` // 1文字、2图片、3文件、4语音
	Data        []byte `json:"-"`            // 用于数据传输
}

// 广播消息
type broadcastMsg struct {
	To      string
	Msg     []byte
	MsgType int // 消息类型，默认：websocket.TextMessage
}

// 连接信息管理
type clientManager struct {
	clients    sync.Map // 防止map并发访问
	broadcast  chan []byte
	register   chan *client
	unregister chan *client
}

// 定义客户端数据结构
type client struct {
	uid    string // 连接用户Id
	grpId  string // 组或群ID,默认为0表示一对一
	socket *websocket.Conn
	send   chan *broadcastMsg
}
