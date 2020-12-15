package model

import (
	"xframe/backend/common/constant"
	"xframe/pkg/cron"
)

type Entity struct {
	cron.Entity
	constant.ModelData
}

//分页请求参数
type SelectPageReq struct {
	JobId          int    `form:"jobId"`          //任务ID
	JobName        string `form:"jobName"`        //任务名称
	JobGroup       string `form:"jobGroup"`       //任务组名
	InvokeTarget   string `form:"invokeTarget"`   //调用目标字符串
	CronExpression string `form:"cronExpression"` //cron执行表达式
	MisfirePolicy  string `form:"misfirePolicy"`  //计划执行错误策略（1立即执行 2执行一次 3放弃执行）
	Concurrent     string `form:"concurrent"`     //是否并发执行（0允许 1禁止）
	Status         string `form:"status"`         //状态（0正常 1暂停）

	*constant.PageReq
}

//新增页面请求参数
type AddReq struct {
	JobName        string `form:"jobName"`
	JobParams      string `form:"jobParams"` // 任务参数
	JobGroup       string `form:"jobGroup"`
	InvokeTarget   string `form:"invokeTarget"`
	CronExpression string `form:"cronExpression"`
	MisfirePolicy  string `form:"misfirePolicy"`
	Concurrent     string `form:"concurrent"`
	Status         string `form:"status"`
	Remark         string `form:"remark"`
}

//修改页面请求参数
type EditReq struct {
	JobId int `form:"jobId" binding:"required"`
	AddReq
}
