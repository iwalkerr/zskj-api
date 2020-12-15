package model

import (
	"fmt"
	"xframe/backend/common/db"
	"xframe/pkg/utils/base"
	"xframe/pkg/utils/page"
)

const (
	selectPageList   = `select oper_id,title,business_type,oper_name,dept_name,oper_location,status,oper_time,oper_ip,operator_type from sys_oper_log where 1=1`
	deleteBatch      = `delete from sys_oper_log where oper_id in (%s)`
	deleteRecordAll  = `delete from sys_oper_log where 1=1`
	selectRecordById = `select oper_id,title,business_type,oper_name,dept_name,oper_location,status,oper_time,oper_ip,operator_type,oper_url,request_method,method,json_result,oper_param,error_msg from sys_oper_log where oper_id=?`
	insert           = `insert into sys_oper_log(title,business_type,method,request_method,operator_type,oper_name,dept_name,oper_url,oper_ip,oper_location,oper_param,json_result,status,error_msg,oper_time) values(?,?,?,?,?,?,?,?,?,?,?,?,?,?,?)`
)

// 插入数据
func (e *Entity) Insert() error {
	_, err := db.Conn().Exec(insert, e.Title, e.BusinessType, e.Method, e.RequestMethod, e.OperatorType, e.OperName, e.DeptName, e.OperUrl, e.OperIp, e.OperLocation, e.OperParam, e.JsonResult, e.Status, e.ErrorMsg, e.OperTime)
	return err
}

// 查询一条记录
func SelectRecordById(id int) (e Entity) {
	_ = db.Conn().QueryRow(selectRecordById, id).Scan(&e.OperId, &e.Title, &e.BusinessType, &e.OperName, &e.DeptName, &e.OperLocation, &e.Status, &e.OperTime, &e.OperIp, &e.OperatorType, &e.OperUrl, &e.RequestMethod, &e.Method, &e.JsonResult, &e.OperParam, &e.ErrorMsg)
	return
}

//清空记录
func DeleteRecordAll() error {
	_, err := db.Conn().Exec(deleteRecordAll)
	return err
}

// 批量删除
func DeleteBatch(ids []string) int {
	params := make([]interface{}, len(ids))
	for i, id := range ids {
		params[i] = id
	}
	query := fmt.Sprintf(deleteBatch, base.Placeholders(len(params)))
	rs, _ := db.Conn().Exec(query, params[:]...)
	count, err := rs.RowsAffected()
	if err != nil {
		return 0
	}
	return int(count)
}

// 根据条件分页查询用户列表
func SelectPageList(param *SelectPageReq) (list []Entity) {
	var params []interface{}

	query := selectPageList
	if param.Title != "" {
		query += ` and title like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.Title)
	}
	if param.OperName != "" {
		query += ` and oper_name like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.OperName)
	}
	if param.Status != "" {
		query += ` and status=?`
		params = append(params, param.Status)
	}
	if param.BusinessTypes != "" {
		query += ` and business_type=?`
		params = append(params, param.BusinessTypes)
	}
	if param.BeginTime != "" {
		query += ` and date_format(oper_time,'%y%m%d') >= date_format(?,'%y%m%d')`
		params = append(params, param.BeginTime)
	}
	if param.EndTime != "" {
		query += ` and date_format(oper_time,'%y%m%d') <= date_format(?,'%y%m%d')`
		params = append(params, param.EndTime)
	}
	if param.SortName != "" {
		query += ` order by ` + param.SortName + ` ` + param.SortOrder
	}

	rows, err := page.New(param.PageReq).GetRows(query, params)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var d Entity
		_ = rows.Scan(&d.OperId, &d.Title, &d.BusinessType, &d.OperName, &d.DeptName, &d.OperLocation, &d.Status, &d.OperTime, &d.OperIp, &d.OperatorType)
		list = append(list, d)
	}
	return
}
