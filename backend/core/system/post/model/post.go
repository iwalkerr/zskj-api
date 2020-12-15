package model

import (
	"fmt"
	"xframe/backend/common/db"
	"xframe/pkg/utils/base"
	"xframe/pkg/utils/page"
)

const (
	selectListByPage       = `select post_id,post_code,post_name,post_sort,status,create_time from sys_post where 1=1`
	selectRecordById       = `select post_id,post_code,post_name,post_sort,status,create_time,remark from sys_post where post_id=?`
	checkPostNameUniqueAll = `select post_id from sys_post where post_name=?`
	checkPostCodeUniqueAll = `select post_id from sys_post where post_code=?`
	update                 = `update sys_post set post_name=?,post_code=?,status=?,remark=?,post_sort=?,update_time=?,update_by=? where post_id=?`
	insert                 = `insert into sys_post(post_code,post_name,post_sort,status,create_by,create_time,remark) values(?,?,?,?,?,?,?)`
	deleteBatch            = `delete from sys_post where post_id in(%s)`
	selectListAll          = `select post_id,post_code,post_name,post_sort,status,create_by,create_time,update_by,update_time,remark from sys_post where 1=1`
	selectPostsByUserId    = `SELECT p.post_id, p.post_name, p.post_code FROM sys_post p  LEFT JOIN sys_user_post up on p.post_id=up.post_id LEFT JOIN sys_user u on up.user_id=u.user_id WHERE up.user_id=?`
)

func SelectPostsByUserId(userId int) (list []Entity) {
	rows, err := db.Conn().Query(selectPostsByUserId, userId)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var d Entity
		_ = rows.Scan(&d.PostId, &d.PostName, &d.PostCode)
		list = append(list, d)
	}
	return
}

//获取所有数据
func SelectListAll(param *SelectPageReq) (list []Entity) {
	var params []interface{}

	query := selectListAll
	if param.PostCode != "" {
		query += ` and post_code like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.PostCode)
	}
	if param.Status != "" {
		query += ` and status=?`
		params = append(params, param.Status)
	}
	if param.PostName != "" {
		query += ` and post_name like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.PostName)
	}
	if param.PageReq != nil {
		if param.BeginTime != "" {
			query += ` and date_format(p.create_time,'%y%m%d') >= date_format(?,'%y%m%d')`
			params = append(params, param.BeginTime)
		}
		if param.EndTime != "" {
			query += ` and date_format(p.create_time,'%y%m%d') <= date_format(?,'%y%m%d')`
			params = append(params, param.EndTime)
		}
	}

	rows, err := db.Conn().Query(query, params[:]...)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var d Entity
		_ = rows.Scan(&d.PostId, &d.PostCode, &d.PostName, &d.PostSort, &d.Status, &d.CreateBy, &d.CreateTime, &d.UpdateBy, &d.UpdateTime, &d.Remark)
		list = append(list, d)
	}
	return
}

//批量删除数据记录
func DeleteBatch(params []interface{}) error {
	query := fmt.Sprintf(deleteBatch, base.Placeholders(len(params)))
	_, err := db.Conn().Exec(query, params[:]...)
	return err
}

func (e *Entity) Insert() (int, error) {
	rs, _ := db.Conn().Exec(insert, e.PostCode, e.PostName, e.PostSort, e.Status, e.CreateBy, e.CreateTime, e.Remark)
	id, err := rs.LastInsertId()
	return int(id), err
}

func (e *Entity) Update() error {
	_, err := db.Conn().Exec(update, e.PostName, e.PostCode, e.Status, e.Remark, e.PostSort, e.UpdateTime, e.UpdateBy, e.PostId)
	return err
}

//校验岗位名称是否唯一
func CheckPostCodeUniqueAll(postCode string) int {
	var id int
	_ = db.Conn().QueryRow(checkPostCodeUniqueAll, postCode).Scan(&id)
	return id
}

//校验岗位名称是否唯一
func CheckPostNameUniqueAll(postName string) int {
	var id int
	_ = db.Conn().QueryRow(checkPostNameUniqueAll, postName).Scan(&id)
	return id
}

//根据主键查询数据
func SelectRecordById(id int) (e Entity) {
	_ = db.Conn().QueryRow(selectRecordById, id).Scan(&e.PostId, &e.PostCode, &e.PostName, &e.PostSort, &e.Status, &e.CreateTime, &e.Remark)
	return
}

//根据条件分页查询数据
func SelectListByPage(param *SelectPageReq) (list []Entity) {
	var params []interface{}

	query := selectListByPage
	if param.PostCode != "" {
		query += ` and post_code like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.PostCode)
	}
	if param.Status != "" {
		query += ` and status=?`
		params = append(params, param.Status)
	}
	if param.PostName != "" {
		query += ` and post_name like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.PostName)
	}
	if param.BeginTime != "" {
		query += ` and date_format(p.create_time,'%y%m%d') >= date_format(?,'%y%m%d')`
		params = append(params, param.BeginTime)
	}
	if param.EndTime != "" {
		query += ` and date_format(p.create_time,'%y%m%d') <= date_format(?,'%y%m%d')`
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
		_ = rows.Scan(&d.PostId, &d.PostCode, &d.PostName, &d.PostSort, &d.Status, &d.CreateTime)
		list = append(list, d)
	}
	return
}
