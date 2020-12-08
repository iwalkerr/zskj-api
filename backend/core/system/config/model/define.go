package model

import (
	"time"
	"xframe/backend/common/constant"
)

// 存放实体和sql文件
type Entity struct {
	ConfigId    int64     `json:"config_id"`    // 参数主键
	ConfigName  string    `json:"config_name"`  // 参数名称
	ConfigKey   string    `json:"config_key"`   // 参数键名
	ConfigValue string    `json:"config_value"` // 参数键值
	ConfigType  string    `json:"config_type"`  // 系统内置（Y是 N否）
	CreateBy    string    `json:"create_by"`    // 创建者
	CreateTime  time.Time `json:"create_time"`  // 创建时间
	UpdateBy    string    `json:"update_by"`    // 更新者
	UpdateTime  time.Time `json:"update_time"`  // 更新时间
	Remark      string    `json:"remark"`       //  备注
}

//分页请求参数
type SelectPageReq struct {
	ConfigName string `form:"configName"` //参数名称
	ConfigKey  string `form:"configKey"`  //参数键名
	ConfigType string `form:"configType"` //状态
	*constant.PageReq
}

//检查参数键名请求参数
type CheckConfigKeyAllReq struct {
	ConfigKey string `form:"configKey"  binding:"required"`
}

//新增页面请求参数
type AddReq struct {
	ConfigName  string `form:"configName"  binding:"required"`
	ConfigKey   string `form:"configKey"  binding:"required"`
	ConfigValue string `form:"configValue"  binding:"required"`
	ConfigType  string `form:"configType"  binding:"required"`
	Remark      string `form:"remark"`
}

//检查参数键名请求参数
type CheckConfigKeyReq struct {
	ConfigId  int    `form:"configId"  binding:"required"`
	ConfigKey string `form:"configKey"  binding:"required"`
}

//修改页面请求参数
type EditReq struct {
	ConfigId int `form:"configId" binding:"required"`
	AddReq
}
