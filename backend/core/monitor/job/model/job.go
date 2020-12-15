package model

import (
	"xframe/backend/common/db"
	"xframe/pkg/utils/page"
)

const (
	selectListByPage = `select job_id,job_name,job_group,invoke_target,cron_expression,misfire_policy,status,create_time from sys_job where 1=1`
	selectRecordById = `select * from sys_job where job_id=?`
	update           = `update sys_job set job_name=?,job_params=?,job_group=?,invoke_target=?,cron_expression=?,misfire_policy=?,remark=?,update_by=?,status=? where job_id=?`
	selectListAll    = `select * from sys_job where 1=1`
	insert           = `insert into sys_job(job_name,job_params,job_group,invoke_target,cron_expression,misfire_policy,concurrent,status,remark,create_time,update_time,create_by) values(?,?,?,?,?,?,?,?,?,?,?,?)`
	updateStatus     = `update sys_job set status=? where job_id=?`
)

func DeleteRecord(id int) error {
	_, err := db.Conn().Exec(`delete from sys_job where job_id=?`, id)
	return err
}

func (e *Entity) Insert() (int, error) {
	rs, err := db.Conn().Exec(insert, &e.JobName, &e.JobParams, &e.JobGroup, &e.InvokeTarget, &e.CronExpression, &e.MisfirePolicy, &e.Concurrent, &e.Status, &e.Remark, &e.CreateTime, &e.UpdateTime, &e.CreateBy)
	if err != nil {
		return 0, err
	}
	id, err := rs.LastInsertId()
	return int(id), err
}

// 获取所有数据
func SelectListAll() (list []Entity) {
	_ = db.Conn().Select(&list, selectListAll)
	return
}

func (e *Entity) Update() error {
	_, err := db.Conn().Exec(update, e.JobName, e.JobParams, e.JobGroup, e.InvokeTarget, e.CronExpression, e.MisfirePolicy, e.Remark, e.UpdateBy, e.Status, e.JobId)
	return err
}

func UpdateStatus(id int, status string) error {
	_, err := db.Conn().Exec(updateStatus, status, id)
	return err
}

func SelectRecordById(id int) *Entity {
	var e Entity
	_ = db.Conn().QueryRowx(selectRecordById, id).StructScan(&e)
	return &e
}

//根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) (list []Entity) {
	var params []interface{}

	query := selectListByPage
	if param.JobName != "" {
		query += ` and job_name like  CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.JobName)
	}
	if param.JobGroup != "" {
		query += ` and job_group=?`
		params = append(params, param.JobGroup)
	}
	if param.InvokeTarget != "" {
		query += ` and invoke_target=?`
		params = append(params, param.InvokeTarget)
	}
	if param.CronExpression != "" {
		query += ` and cron_expression=?`
		params = append(params, param.CronExpression)
	}
	if param.MisfirePolicy != "" {
		query += ` and misfire_policy=?`
		params = append(params, param.MisfirePolicy)
	}
	if param.Concurrent != "" {
		query += ` and concurrent=?`
		params = append(params, param.Concurrent)
	}
	if param.Status != "" {
		query += ` and status=?`
		params = append(params, param.Status)
	}
	if param.BeginTime != "" {
		query += ` and date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d')`
		params = append(params, param.BeginTime)
	}
	if param.EndTime != "" {
		query += ` and date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d')`
		params = append(params, param.EndTime)
	}
	if param.SortName != "" && param.SortOrder != "" {
		query += ` order by ` + param.SortName + ` ` + param.SortOrder
	}

	rows, err := page.New(param.PageReq).GetRows(query, params)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var d Entity
		_ = rows.Scan(&d.JobId, &d.JobName, &d.JobGroup, &d.InvokeTarget, &d.CronExpression, &d.MisfirePolicy, &d.Status, &d.CreateTime)
		list = append(list, d)
	}
	return
}
