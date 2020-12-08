package cron

type Entity struct {
	JobId          int    `db:"job_id" json:"job_id"`                   // 任务ID
	JobName        string `db:"job_name" json:"job_name"`               // 任务名称
	JobParams      string `db:"job_params" json:"job_params"`           // 参数
	JobGroup       string `db:"job_group" json:"job_group"`             // 任务组名 1默认 2系统
	InvokeTarget   string `db:"invoke_target" json:"invoke_target"`     // 调用目标字符串
	CronExpression string `db:"cron_expression" json:"cron_expression"` // cron执行表达式
	MisfirePolicy  string `db:"misfire_policy" json:"misfire_policy"`   // 计划执行策略（1多次执行 2执行一次）
	Concurrent     string `db:"concurrent" json:"concurrent"`           // 是否并发执行（0允许 1禁止）
	Status         string `db:"status" json:"status"`                   // 状态（0正常 1暂停）
}
