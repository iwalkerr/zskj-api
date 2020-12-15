package model

import (
	"time"
	"xframe/backend/common/constant"
)

type Entity struct {
	InfoId        int       `db:"info_id" json:"info_id"`               // 访问ID
	LoginName     string    `db:"login_name" json:"login_name"`         // 登录账号
	Ipaddr        string    `db:"ipaddr" json:"ipaddr"`                 // 登录IP地址
	LoginLocation string    `db:"login_location" json:"login_location"` // 登录地点
	Browser       string    `db:"browser" json:"browser"`               // 浏览器类型
	Os            string    `db:"os" json:"os"`                         // 操作系统
	Status        string    `db:"status" json:"status"`                 // 登录状态（0成功 1失败）
	Msg           string    `db:"msg" json:"msg"`                       // 提示消息
	LoginTime     time.Time `db:"login_time" json:"login_time"`         // 访问时间
}

//查询列表请求参数
type SelectPageReq struct {
	LoginName string `form:"loginName"` //登陆名
	Status    string `form:"status"`    //状态
	Ipaddr    string `form:"ipaddr"`    //登录地址
	*constant.PageReq
}
