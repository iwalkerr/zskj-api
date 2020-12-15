package ws

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 初始化连接
func init() {
	go connPool.monitor()
}

// 对外路由API调用
func RunWsServer(c *gin.Context) {
	conn, err := upGrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		http.NotFound(c.Writer, c.Request)
		return
	}
	grpId := c.Param("grpId")
	uid := c.Param("uid")

	// 数据库查询该用户是否存在，且是否在这个群组中

	newConn(conn, grpId, uid)
}

// 初始化连接信息
func newConn(conn *websocket.Conn, grpId, uid string) {
	c := &client{grpId: grpId, uid: uid, socket: conn, send: make(chan *broadcastMsg)}
	connPool.register <- c

	go c.readInMessage()
	go c.processMsg()
}

// 读取数据监听
func (c *client) readInMessage() {
	for {
		_, message, err := c.socket.ReadMessage()
		if err != nil {
			connPool.unregister <- c
			break
		}
		connPool.broadcast <- message
	}
}

// 发送数据
func (c *client) processMsg() {
	for {
		if message, ok := <-c.send; !ok {
			break
		} else {
			_ = c.socket.WriteMessage(message.MsgType, message.Msg)
		}
	}
}

// 启动全局监听
func (m *clientManager) monitor() {
	for {
		select {
		case message := <-m.broadcast:
			m.check(message)
		case conn := <-m.register:
			m.clients.Store(conn, true)
			m.dynamic()
		case conn := <-m.unregister:
			m.clients.Delete(conn)
			close(conn.send)
			_ = conn.socket.Close()
			m.dynamic()
		}
	}
}
