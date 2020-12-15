package model

import (
	"time"
	"xframe/backend/common/constant"
)

type Entity struct {
	OperId        int64     `json:"oper_id"`        // 日志主键
	Title         string    `json:"title"`          // 模块标题
	BusinessType  int       `json:"business_type"`  // 业务类型（0其它 1新增 2修改 3删除）
	Method        string    `json:"method"`         // 方法名称
	RequestMethod string    `json:"request_method"` // 请求方式
	OperatorType  int       `json:"operator_type"`  // 操作类别（0其它 1后台用户 2手机端用户）
	OperName      string    `json:"oper_name"`      // 操作人员
	DeptName      string    `json:"dept_name"`      // 部门名称
	OperUrl       string    `json:"oper_url"`       // 请求URL
	OperIp        string    `json:"oper_ip"`        // 主机地址
	OperLocation  string    `json:"oper_location"`  // 操作地点
	OperParam     string    `json:"oper_param"`     // 请求参数
	JsonResult    string    `json:"json_result"`    // 返回参数
	Status        int       `json:"status"`         // 操作状态（0正常 1异常）
	ErrorMsg      string    `json:"error_msg"`      // 错误消息
	OperTime      time.Time `json:"oper_time"`      // 操作时间
}

//查询列表请求参数
type SelectPageReq struct {
	Title         string `form:"title"`         //系统模块
	OperName      string `form:"operName"`      //操作人员
	BusinessTypes string `form:"businessTypes"` //操作类型
	Status        string `form:"status"`        //操作类型

	*constant.PageReq
}
