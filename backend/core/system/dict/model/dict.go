package model

import (
	"xframe/backend/common/db"
	"xframe/pkg/utils/page"
)

const (
	listByDictType   = `select dict_label,dict_value,list_class,dict_id from sys_dict where status = '0' and parent_id=(SELECT dict_id FROM sys_dict WHERE dict_type=?) order by dict_sort asc`
	labelRecord      = `select dict_label from sys_dict where dict_value=? and dict_type=?`
	selectDictList   = `select dict_id,parent_id,dict_label from sys_dict where status='0'`
	selectPageList   = `select * from sys_dict where 1=1`
	insert           = `insert into sys_dict(dict_type,status,dict_label,dict_sort,dict_value,is_default,list_class,remark,create_time,create_by,parent_id,update_time) values(?,?,?,?,?,?,?,?,?,?,?,?)`
	selectRecordById = `select dict_id,dict_label,dict_sort,dict_value,dict_type,status,remark,is_default,create_time from sys_dict where dict_id=?`
	update           = `update sys_dict set dict_label=?,dict_sort=?,dict_value=?,dict_type=?,status=?,remark=?,update_time=?,is_default=?,list_class=? where dict_id=?`
	childDictCount   = `select count(dict_id) from sys_dict where parent_id=?`
	deleteRecordById = `delete from sys_dict where dict_id=?`
)

func DeleteRecordById(id string) error {
	_, err := db.Conn().Exec(deleteRecordById, id)
	return err
}

func ChildDictCount(id string) (count int) {
	_ = db.Conn().QueryRow(childDictCount, id).Scan(&count)
	return
}

// 更新数据
func (e *Entity) Update() error {
	_, err := db.Conn().Exec(update, e.DictLabel, e.DictSort, e.DictValue, e.DictType, e.Status, e.Remark, e.UpdateTime, e.IsDefault, e.ListClass, e.DictId)
	return err
}

//根据主键查询数据
func SelectRecordById(id int) (e Entity) {
	_ = db.Conn().QueryRow(selectRecordById, id).Scan(&e.DictId, &e.DictLabel, &e.DictSort, &e.DictValue, &e.DictType, &e.Status, &e.Remark, &e.IsDefault, &e.CreateTime)
	return
}

// 插入数据
func (e *Entity) Insert() (int, error) {
	rs, _ := db.Conn().Exec(insert, e.DictType, e.Status, e.DictLabel, e.DictSort, e.DictValue, e.IsDefault, e.ListClass, e.Remark, e.CreateTime, e.CreateBy, e.ParentId, e.UpdateTime)
	id, err := rs.LastInsertId()
	return int(id), err
}

func SelectPageList(param *SelectPageReq) (list []Entity) {
	var params []interface{}
	query := selectPageList
	if param.DictId != "" {
		query += ` and parent_id=?`
		params = append(params, param.DictId)
	}
	if param.DictLabel != "" {
		query += ` and dict_label like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.DictLabel)
	}
	if param.Status != "" {
		query += ` and status=?`
		params = append(params, param.Status)
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
		_ = rows.StructScan(&d)
		list = append(list, d)
	}
	return
}

func SelectDictList(parentId int) (list []Entity) {
	var params []interface{}

	query := selectDictList
	if parentId > 0 {
		query += ` and parent_id=?`
		params = append(params, parentId)
	}
	rows, err := db.Conn().Queryx(query, params[:]...)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var d Entity
		_ = rows.StructScan(&d)
		list = append(list, d)
	}
	return
}

// // 根据条件查询
func LabelRecord(dictType, dictValue string) (label string) {
	_ = db.Conn().QueryRow(labelRecord, dictValue, dictType).Scan(&label)
	return
}

// 更加字典类型获取集合
func ListByDictType(dictType string) (list []DictData) {
	_ = db.Conn().Select(&list, listByDictType, dictType)
	return
}
