package ws

import (
	"encoding/json"
	"fmt"
	"sort"
	"strings"

	"github.com/gorilla/websocket"
)

// 110 后台在线分组

// 监控客户端动态
func (m *clientManager) dynamic() {
	m.onlineNum() // 在线数量
	m.sysOnUser() // 后台在线用户
}

// 后台在线用户
func (m *clientManager) sysOnUser() {
	var uids []string // 获取在线用户uid

	m.forr(func(conn *client) (*broadcastMsg, bool) {
		if conn.grpId == OnSysUserGrpId {
			uids = append(uids, conn.uid)
		}
		return nil, false
	})
	// 排序
	sort.Strings(uids)

	rec := receiveMsg{GrpId: OnSysUserGrpId, Content: strings.Join(uids, ","), ContentType: 1}
	bytes, _ := json.Marshal(rec)

	rec = receiveMsg{GrpId: OnSysUserGrpId, Data: bytes}
	m.emit(&rec)
}

// 统计连接数量
func (m *clientManager) onlineNum() {
	// 后台接收数据,用于查看app平板的动态上下线
	var count int
	m.forr(func(conn *client) (*broadcastMsg, bool) {
		count++
		return nil, false
	})

	fmt.Printf("目前在线人数: %v\n", count)
}

// 检查接收数据中type的类型
func (m *clientManager) check(data []byte) {
	if rec := unmarshal(data); rec != nil {
		// 发送给组群中的所有人
		m.emit(rec)
	}
}

// 发送消息
func (m *clientManager) emit(rec *receiveMsg) {
	bm := newBroadcastMsg(rec.ToUser, rec.Data)
	switch rec.ContentType {
	case 4: // 30s 语音
		bm.MsgType = websocket.BinaryMessage
	default:
	}

	// 发送给指定类型
	switch {
	case rec.ToUser != "": // 一对一发送
		m.forr(func(conn *client) (*broadcastMsg, bool) {
			return bm, conn.uid == bm.To
		})
	case rec.GrpId != "": // 一对多广播
		m.forr(func(conn *client) (*broadcastMsg, bool) {
			return bm, conn.grpId == rec.GrpId
		})
	}
}

// 默认设置属性
func newBroadcastMsg(to string, msg []byte) *broadcastMsg {
	return &broadcastMsg{To: to, Msg: msg, MsgType: websocket.TextMessage}
}

// 遍历连接
func (m *clientManager) forr(callback func(conn *client) (*broadcastMsg, bool)) {
	m.clients.Range(func(k, v interface{}) bool {
		if conn, ok := k.(*client); ok {
			if broadcastMsg, ok := callback(conn); ok {
				conn.send <- broadcastMsg
			}
		}
		return true
	})
}

// 解析数据
func unmarshal(data []byte) (rec *receiveMsg) {
	_ = json.Unmarshal(data, &rec)
	if rec.FromUser == "" || rec.Content == "" {
		return nil
	}

	rec.Data = data
	return
}
