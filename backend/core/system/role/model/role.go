package model

import (
	"database/sql"
	"fmt"
	"xframe/backend/common/db"
	"xframe/pkg/utils/base"
	"xframe/pkg/utils/page"
)

const (
	selectListPage         = `SELECT r.role_id, r.role_name, r.role_key, r.role_sort, r.data_scope,r.status, r.create_time, r.remark from sys_role r where 1=1 and r.del_flag='0'`
	checkRoleNameUniqueAll = `select count(role_id) from sys_role WHERE role_name=?`
	checkRoleKeyUniqueAll  = `select count(role_id) from sys_role WHERE role_key=?`
	add                    = `insert into sys_role(role_id,role_name,role_key,role_sort,data_scope,status,create_by,create_time,remark) values(?,?,?,?,?,?,?,?,?)`
	addRoleMenus           = `insert into sys_role_menu(role_id,menu_id) values(?,?)`
	selectRecordById       = `select role_id,role_name,role_key,role_sort,data_scope,status,remark from sys_role where role_id=?`
	checkRoleKeyUnique     = `select count(role_id) from sys_role where role_key=? and role_id<>?`
	checkRoleNameUnique    = `select count(role_id) from sys_role where role_name=? and role_id<>?`
	editSave               = `update sys_role set role_name=?,role_key=?,status=?,remark=?,update_time=?,create_by=?,role_sort=? where role_id=?`
	delRoleMenuById        = `delete from sys_role_menu where role_id=?`
	changeStatus           = `update sys_role set status=? where role_id=?`
	deleteBatch            = `delete from sys_role where role_id in (%s)`
	delRoleDeptById        = `delete from sys_role_dept where role_id=?`
	addRoleDepts           = `insert into sys_role_dept(role_id,dept_id) values(?,?)`
	selectListAll          = `SELECT role_id,role_name,role_key,role_sort,data_scope,status,create_by,create_time,update_by,update_time,remark FROM sys_role WHERE del_flag='0'`
	selectRoleContactVo    = `SELECT r.role_id, r.role_name, r.role_key, r.role_sort, r.data_scope,r.status, r.create_time, r.remark FROM sys_role r LEFT JOIN sys_user_role ur on r.role_id=ur.role_id LEFT JOIN sys_user u on ur.user_id=u.user_id LEFT JOIN sys_dept d on u.dept_id=d.dept_id WHERE r.del_flag='0' and ur.user_id=?`
)

//根据用户ID查询角色
func SelectRoleContactVo(userId int) (list []Entity) {
	rows, err := db.Conn().Query(selectRoleContactVo, userId)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var e Entity
		_ = rows.Scan(&e.RoleId, &e.RoleName, &e.RoleKey, &e.RoleSort, &e.DataScope, &e.Status, &e.CreateTime, &e.Remark)
		list = append(list, e)
	}
	return
}

//获取所有角色数据
func SelectListAll(param *SelectPageReq) (list []Entity) {
	var params []interface{}

	query := selectListAll
	if param.RoleName != "" {
		query += ` and role_name like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.RoleName)
	}
	if param.RoleKey != "" {
		query += ` and role_key like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.RoleKey)
	}
	if param.Status != "" {
		query += ` and status=?`
		params = append(params, param.Status)
	}
	if param.DataScope != "" {
		query += ` and data_scope=?`
		params = append(params, param.DataScope)
	}
	if param.PageReq != nil {
		if param.BeginTime != "" {
			query += ` and date_format(r.create_time,'%y%m%d') >= date_format(?,'%y%m%d')`
			params = append(params, param.BeginTime)
		}
		if param.EndTime != "" {
			query += ` and date_format(r.create_time,'%y%m%d') <= date_format(?,'%y%m%d')`
			params = append(params, param.EndTime)
		}
	}

	rows, err := db.Conn().Query(query, params[:]...)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var d Entity
		_ = rows.Scan(&d.RoleId, &d.RoleName, &d.RoleKey, &d.RoleSort, &d.DataScope, &d.Status, &d.CreateBy, &d.CreateTime, &d.UpdateBy, &d.UpdateTime, &d.Remark)
		list = append(list, d)
	}
	return
}

// 添加角色部门
func (e *Entity) AddRoleDepts(tx *sql.Tx, list []EntityRoleDept) error {
	stmt, err := tx.Prepare(addRoleDepts)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	var errs error
	for _, roleDept := range list {
		_, err := stmt.Exec(roleDept.RoleId, roleDept.DeptId)
		if err != nil {
			errs = err
			break
		}
	}
	if errs != nil {
		_ = tx.Rollback()
	}
	return errs
}

// 根据角色ID删除关联数据
func DelRoleDeptById(tx *sql.Tx, roleId int) error {
	_, err := tx.Exec(delRoleDeptById, roleId)
	if err != nil {
		_ = tx.Rollback()
	}
	return err
}

// 批量删除角色
func DeleteBatch(ids []string) error {
	params := make([]interface{}, len(ids))
	for i, id := range ids {
		params[i] = id
	}
	query := fmt.Sprintf(deleteBatch, base.Placeholders(len(ids)))
	_, err := db.Conn().Exec(query, params[:]...)
	return err
}

func (e *Entity) ChangeStatus() error {
	_, err := db.Conn().Exec(changeStatus, e.Status, e.RoleId)
	return err
}

// 根据角色ID删除关联数据
func DelRoleMenuById(tx *sql.Tx, roleId int) error {
	_, err := tx.Exec(delRoleMenuById, roleId)
	if err != nil {
		_ = tx.Rollback()
	}
	return err
}

func (e *Entity) EditSave(tx *sql.Tx) error {
	_, err := tx.Exec(editSave, e.RoleName, e.RoleKey, e.Status, e.Remark, e.UpdateTime, e.CreateBy, e.RoleSort, e.RoleId)
	if err != nil {
		_ = tx.Rollback()
	}
	return err
}

//检查角色键是否唯一
func CheckRoleNameUnique(roleName string, roleId int) bool {
	var count int
	_ = db.Conn().QueryRow(checkRoleNameUnique, roleName, roleId).Scan(&count)
	return count == 0
}

//检查角色键是否唯一
func CheckRoleKeyUnique(roleKey string, roleId int) bool {
	var count int
	_ = db.Conn().QueryRow(checkRoleKeyUnique, roleKey, roleId).Scan(&count)
	return count == 0
}

//根据主键查询数据
func SelectRecordById(id int) (e Entity) {
	_ = db.Conn().QueryRow(selectRecordById, id).Scan(&e.RoleId, &e.RoleName, &e.RoleKey, &e.RoleSort, &e.DataScope, &e.Status, &e.Remark)
	return
}

// 添加角色菜单
func (e *Entity) AddRoleMenus(tx *sql.Tx, list []EntityRoleMenu) error {
	stmt, err := tx.Prepare(addRoleMenus)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	var errs error
	for _, roleMenu := range list {
		_, err := stmt.Exec(roleMenu.RoleId, roleMenu.MenuId)
		if err != nil {
			errs = err
			break
		}
	}
	if errs != nil {
		_ = tx.Rollback()
	}
	return errs
}

func (e *Entity) Add(tx *sql.Tx) (int, error) {
	res, _ := tx.Exec(add, e.RoleId, e.RoleName, e.RoleKey, e.RoleSort, e.DataScope, e.Status, e.CreateBy, e.CreateTime, e.Remark)
	id, err := res.LastInsertId()
	if err != nil || id <= 0 {
		_ = tx.Rollback()
		return 0, err
	}
	return int(id), nil
}

//检查角色键是否唯一
func CheckRoleKeyUniqueAll(roleKey string) bool {
	var count int
	_ = db.Conn().QueryRow(checkRoleKeyUniqueAll, roleKey).Scan(&count)
	return count != 0
}

// 检查角色键是否唯一
func CheckRoleNameUniqueAll(roleName string) bool {
	var count int
	_ = db.Conn().QueryRow(checkRoleNameUniqueAll, roleName).Scan(&count)
	return count != 0
}

// 根据条件分页查询角色数据
func SelectListPage(param *SelectPageReq) (list []Entity) {
	var params []interface{}

	query := selectListPage
	if param.RoleName != "" {
		query += ` and r.role_name like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.RoleName)
	}
	if param.Status != "" {
		query += ` and r.status=?`
		params = append(params, param.Status)
	}
	if param.RoleKey != "" {
		query += ` and r.role_key like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.RoleKey)
	}
	if param.DataScope != "" {
		query += ` and r.data_scope=?`
		params = append(params, param.DataScope)
	}
	if param.BeginTime != "" {
		query += ` and date_format(r.create_time,'%y%m%d') >= date_format(?,'%y%m%d')`
		params = append(params, param.BeginTime)
	}
	if param.EndTime != "" {
		query += ` and date_format(r.create_time,'%y%m%d') <= date_format(?,'%y%m%d')`
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
		_ = rows.Scan(&d.RoleId, &d.RoleName, &d.RoleKey, &d.RoleSort, &d.DataScope, &d.Status, &d.CreateTime, &d.Remark)
		list = append(list, d)
	}
	return
}
