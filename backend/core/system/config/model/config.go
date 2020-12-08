package model

import (
	"fmt"
	"xframe/backend/common/db"
	"xframe/pkg/utils/base"
	"xframe/pkg/utils/page"
)

const (
	selectListByPage        = `select config_id,config_name,config_key,config_value,config_type,remark,create_time from sys_config where 1=1`
	selectRecordById        = `select config_id,config_name,config_key,config_type,config_value,remark,create_time from sys_config where config_id=?`
	checkConfigKeyUniqueAll = `select config_id from sys_config where config_key=?`
	update                  = `update sys_config set config_name=?,config_key=?,config_type=?,config_value=?,remark=?,update_by=?,update_time=? where config_id=?`
	insert                  = `insert sys_config(config_name,config_key,config_type,config_value,remark,create_time,create_by) values(?,?,?,?,?,?,?)`
	deleteBatch             = `delete from sys_config where config_id in (%s)`
	getValueByKey           = `select config_value from sys_config where config_key=?`
)

func GetValueByKey(key string) (value string) {
	_ = db.Conn().QueryRow(getValueByKey, key).Scan(&value)
	return
}

// 批量删除
func DeleteBatch(ids []string) error {
	params := make([]interface{}, len(ids))
	for i, id := range ids {
		params[i] = id
	}
	query := fmt.Sprintf(deleteBatch, base.Placeholders(len(params)))
	_, err := db.Conn().Exec(query, params[:]...)
	return err
}

func (e *Entity) Insert() (int, error) {
	rs, _ := db.Conn().Exec(insert, e.ConfigName, e.ConfigKey, e.ConfigType, e.ConfigValue, e.Remark, e.CreateTime, e.CreateBy)
	id, err := rs.LastInsertId()
	return int(id), err
}

// 更新数据
func (e *Entity) Update() (int, error) {
	rs, _ := db.Conn().Exec(update, e.ConfigName, e.ConfigKey, e.ConfigType, e.ConfigValue, e.Remark, e.UpdateBy, e.UpdateTime, e.ConfigId)
	count, err := rs.RowsAffected()
	return int(count), err
}

//校验参数键名是否唯一
func CheckConfigKeyUniqueAll(configKey string) (id int) {
	_ = db.Conn().QueryRow(checkConfigKeyUniqueAll, configKey).Scan(&id)
	return
}

//根据主键查询数据
func SelectRecordById(id int) (e Entity) {
	_ = db.Conn().QueryRow(selectRecordById, id).Scan(&e.ConfigId, &e.ConfigName, &e.ConfigKey, &e.ConfigType, &e.ConfigValue, &e.Remark, &e.CreateTime)
	return e
}

func SelectListByPage(param *SelectPageReq) (list []Entity) {
	var params []interface{}
	query := selectListByPage

	if param.ConfigName != "" {
		query += ` and config_name like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.ConfigName)
	}
	if param.ConfigType != "" {
		query += ` and config_type=?`
		params = append(params, param.ConfigType)
	}
	if param.ConfigKey != "" {
		query += ` and config_key like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.ConfigKey)
	}
	if param.BeginTime != "" {
		query += ` and date_format(t.create_time,'%y%m%d') >= date_format(?,'%y%m%d')`
		params = append(params, param.BeginTime)
	}
	if param.EndTime != "" {
		query += ` and date_format(t.create_time,'%y%m%d') <= date_format(?,'%y%m%d')`
		params = append(params, param.EndTime)
	}

	rows, err := page.New(param.PageReq).GetRows(query, params)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var d Entity
		_ = rows.Scan(&d.ConfigId, &d.ConfigName, &d.ConfigKey, &d.ConfigValue, &d.ConfigType, &d.Remark, &d.CreateTime)
		list = append(list, d)
	}
	return
}
