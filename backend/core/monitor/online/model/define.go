package model

import (
	"time"
	"xframe/backend/common/constant"
)

type Entity struct {
	UserId        int       `json:"user_id"`        // 用户ID
	LoginName     string    `json:"login_name"`     // 登录账号
	Ipaddr        string    `json:"ipaddr"`         // 登录IP地址
	LoginLocation string    `json:"login_location"` // 登录地点
	Browser       string    `json:"browser"`        // 浏览器类型
	Os            string    `json:"os"`             // 操作系统
	StartTime     time.Time `json:"start_time"`     // session创建时间
}

// 分页请求参数
type SelectPageReq struct {
	Ids string `form:"ids"` //登录用户Id,以逗号分隔
	*constant.PageReq
	IdsArr []string
}
