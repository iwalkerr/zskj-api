package model

import (
	"fmt"
	"xframe/backend/common/db"
	"xframe/pkg/utils/base"
	"xframe/pkg/utils/page"
)

const (
	insert            = `insert into sys_login_log(info_id,login_name,ipaddr,login_location,browser,os,status,msg,login_time) values(?,?,?,?,?,?,?,?,?)`
	selectPageList    = `select info_id,login_name,ipaddr,login_location,browser,os,status,msg,login_time from sys_login_log where 1=1`
	deleteRecordByIds = `delete from sys_login_log where info_id in (%s)`
)

//批量删除
func DeleteRecordByIds(ids []string) int {
	params := make([]interface{}, len(ids))
	for i, id := range ids {
		params[i] = id
	}
	query := fmt.Sprintf(deleteRecordByIds, base.Placeholders(len(params)))
	rs, _ := db.Conn().Exec(query, params[:]...)
	count, _ := rs.RowsAffected()
	return int(count)
}

// 插入日志
func (e *Entity) Insert() error {
	_, err := db.Conn().Exec(insert, e.InfoId, e.LoginName, e.Ipaddr, e.LoginLocation, e.Browser, e.Os, e.Status, e.Msg, e.LoginTime)
	return err
}

// 根据条件分页查询用户列表
func SelectPageList(param *SelectPageReq) (list []Entity) {
	var params []interface{}

	query := selectPageList
	if param.LoginName != "" {
		query += ` and login_name like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.LoginName)
	}
	if param.Ipaddr != "" {
		query += ` and ipaddr like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.Ipaddr)
	}
	if param.Status != "" {
		query += ` and status=?`
		params = append(params, param.Status)
	}
	if param.BeginTime != "" {
		query += ` and date_format(login_time,'%y%m%d') >= date_format(?,'%y%m%d')`
		params = append(params, param.BeginTime)
	}
	if param.EndTime != "" {
		query += ` and date_format(login_time,'%y%m%d') <= date_format(?,'%y%m%d')`
		params = append(params, param.EndTime)
	}
	if param.SortName != "" {
		query += ` order by ` + param.SortName + ` ` + param.SortOrder
	}

	rows, err := page.New(param.PageReq).GetRows(query, params)
	if err != nil {
		return nil
	}

	defer rows.Close()

	for rows.Next() {
		var d Entity
		_ = rows.Scan(&d.InfoId, &d.LoginName, &d.Ipaddr, &d.LoginLocation, &d.Browser, &d.Os, &d.Status, &d.Msg, &d.LoginTime)
		list = append(list, d)
	}
	return
}
