package model

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"
	"xframe/backend/common/db"
	online "xframe/backend/core/monitor/online/model"
	role "xframe/backend/core/system/role/model"
	"xframe/pkg/utils/base"
	"xframe/pkg/utils/page"
)

const (
	isUserLock            = `select status from sys_user where login_name=?`
	findOneSql            = `select user_id,phonenumber,login_name,user_name,email,sex,create_time,avatar,password,salt,dept_id,status from sys_user where login_name=?`
	selectPageList        = `SELECT u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber, u.password,u.sex, u.salt, u.status, u.del_flag, u.login_ip, u.create_by, u.create_time, u.remark, d.dept_name, d.leader FROM sys_user u LEFT JOIN sys_dept d on u.dept_id=d.dept_id WHERE u.del_flag='0'`
	selectAllocatedList   = `SELECT u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber,u.status, u.create_time from sys_user u LEFT JOIN sys_user_role ur on u.user_id=ur.user_id LEFT JOIN sys_role r on ur.role_id=r.role_id WHERE u.del_flag='0' and r.role_id=?`
	selectUnAllocatedList = `SELECT u.user_id, u.dept_id, u.login_name, u.user_name, u.email, u.avatar, u.phonenumber,u.status, u.create_time from sys_user u LEFT JOIN sys_user_role ur on u.user_id=ur.user_id LEFT JOIN sys_role r on ur.role_id=r.role_id WHERE u.del_flag='0' and u.user_id not in (select u.user_id from sys_user u inner join sys_user_role ur on u.user_id = ur.user_id and ur.role_id=?)`
	addUserRole           = `insert into sys_user_role(user_id,role_id) values(?,?)`
	deleteUserRoleInfos   = `delete from sys_user_role where role_id=? and user_id in (%s)`
	selectRecordById      = `select user_id,phonenumber,login_name,user_name,email,sex,create_time,avatar,dept_id,status,remark from sys_user where user_id=?`
	checkPhoneUnique      = `select count(user_id) from sys_user where phonenumber=? and user_id<>?`
	checkEmailUnique      = `select count(user_id) from sys_user where email=? and user_id<>?`
	updateSave            = `update sys_user set user_name=?,email=?,phonenumber=?,status=?,sex=?,dept_id=?,remark=?,update_time=?,update_by=? where user_id=?`
	delAddUserPosts       = `delete from sys_user_post where user_id=?`
	addUserPost           = `insert into sys_user_post(user_id,post_id) value(?,?)`
	delAddUserRoles       = `delete from sys_user_role where user_id=?`
	checkPhoneUniqueAll   = `select count(user_id) from sys_user where phonenumber=?`
	checkLoginName        = `select count(user_id) from sys_user where login_name=?`
	checkEmailUniqueAll   = `select count(user_id) from sys_user where email=?`
	addUser               = `insert into sys_user(user_id,password,salt,phonenumber,login_name,user_name,email,sex,create_time,avatar,dept_id,remark,create_by) values(?,?,?,?,?,?,?,?,?,?,?,?,?)`
	deleteRecordByIds     = `delete from sys_user where user_id in (%s)`
	deleteUserRoleByIds   = `delete from sys_user_role where user_id in (%s)`
	deleteUserPostsByIds  = `delete from sys_user_post where user_id in (%s)`
	upadtePwd             = `update sys_user set password=?,salt=? where user_id=?`
	selectExportList      = `select u.login_name,u.user_name,u.email,u.phonenumber,u.sex,d.dept_name,d.leader,u.status,u.del_flag,u.create_by,u.create_time,u.remark from sys_user u left join sys_dept d on u.dept_id=d.dept_id where 1=1 and u.del_flag='0'`
	upadte                = `update sys_user set avatar=? where user_id=?`
	upadteData            = `update sys_user set user_name=?,email=?,phonenumber=?,sex=? where user_id=?`
	listUsersByIds        = `select user_id,login_name,user_name,phonenumber,email,login_ip,login_date,create_by from sys_user where user_id in (%s)`
	updateLoginTime       = `update sys_user set login_date=?,login_ip=? where user_id=?`
)

// 更新登陆时间
func UpdateLoginTime(userId int, loginTime time.Time, clientIp string) {
	_, _ = db.Conn().Exec(updateLoginTime, loginTime, clientIp, userId)
}

// 根据ids查询用户列表
func ListUsersByIds(param *online.SelectPageReq) (list []Entity) {
	query := fmt.Sprintf(listUsersByIds, strings.Join(param.IdsArr, ","))

	var params []interface{}
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

// 更新基础用户数据
func (r *ToSession) UpdateData() error {
	_, err := db.Conn().Exec(upadteData, r.UserName, r.Email, r.PhoneNumber, r.Sex, r.UserId)
	return err
}

// 更新头像
func (r *ToSession) Update() error {
	_, err := db.Conn().Exec(upadte, r.Avatar, r.UserId)
	return err
}

// 导出excel
func SelectExportList(param *SelectPageReq) (list []UserListEntity) {
	var params []interface{}

	query := selectExportList
	if param.LoginName != "" {
		query += ` and u.login_name=?`
		params = append(params, param.LoginName)
	}
	if param.Phonenumber != "" {
		query += ` and u.phonenumber=?`
		params = append(params, param.Phonenumber)
	}
	if param.Status != "" {
		query += ` and u.status=?`
		params = append(params, param.Status)
	}
	if param.BeginTime != "" {
		query += ` and date_format(u.create_time,'%y%m%d') >= date_format(?,'%y%m%d')`
		params = append(params, param.BeginTime)
	}
	if param.EndTime != "" {
		query += ` and date_format(u.create_time,'%y%m%d') <= date_format(?,'%y%m%d')`
		params = append(params, param.EndTime)
	}
	if param.DeptId != 0 {
		query += ` and (u.dept_id=? or u.dept_id in(SELECT t.dept_id FROM sys_dept t WHERE FIND_IN_SET (?,ancestors) ))`
		params = append(params, param.DeptId, param.DeptId)
	}

	rows, err := db.Conn().Query(query, params[:]...)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var d UserListEntity
		_ = rows.Scan(&d.LoginName, &d.UserName, &d.Email, &d.PhoneNumber, &d.Sex, &d.DeptName, &d.Leader, &d.Status, &d.DelFlag, &d.CreateBy, &d.CreateTime, &d.Remark)
		if d.Sex == "0" {
			d.Sex = "女"
		} else if d.Sex == "1" {
			d.Sex = "男"
		}
		list = append(list, d)
	}
	return
}

// 更新密码
func (r *Entity) UpdatePwd() error {
	_, err := db.Conn().Exec(upadtePwd, r.Password, r.Salt, r.UserId)
	return err
}

// 删除用户岗位
func DeleteUserPostsByIds(tx *sql.Tx, ids []string) error {
	params := make([]interface{}, len(ids))
	for i, id := range ids {
		params[i] = id
	}
	query := fmt.Sprintf(deleteUserPostsByIds, base.Placeholders(len(params)))
	_, err := tx.Exec(query, params[:]...)
	if err != nil {
		_ = tx.Rollback()
	}
	return err
}

// 删除用户角色
func DeleteUserRoleByIds(tx *sql.Tx, ids []string) error {
	params := make([]interface{}, len(ids))
	for i, id := range ids {
		params[i] = id
	}
	query := fmt.Sprintf(deleteUserRoleByIds, base.Placeholders(len(params)))
	_, err := tx.Exec(query, params[:]...)
	if err != nil {
		_ = tx.Rollback()
	}
	return err
}

// 批量删除
func DeleteRecordByIds(tx *sql.Tx, ids []string) error {
	params := make([]interface{}, len(ids))
	for i, id := range ids {
		params[i] = id
	}
	query := fmt.Sprintf(deleteRecordByIds, base.Placeholders(len(params)))
	_, err := tx.Exec(query, params[:]...)
	if err != nil {
		_ = tx.Rollback()
	}
	return err
}

// 添加用户
func (r *Entity) AddUser(tx *sql.Tx) error {
	res, err := tx.Exec(addUser, r.UserId, r.Password, r.Salt, r.PhoneNumber, r.LoginName, r.UserName, r.Email, r.Sex, r.CreateTime, r.Avatar, r.DeptId, r.Remark, r.CreateBy)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	id, err := res.LastInsertId()
	if err != nil || id <= 0 {
		_ = tx.Rollback()
		return err
	}
	r.UserId = int(id)
	return nil
}

// 检查邮箱是否存在,存在返回true,否则false
func CheckEmailUniqueAll(email string) bool {
	var count int
	_ = db.Conn().QueryRow(checkEmailUniqueAll, email).Scan(&count)
	return count != 0
}

// 检查登录名是否唯一
func CheckLoginName(loginName string) bool {
	var count int
	_ = db.Conn().QueryRow(checkLoginName, loginName).Scan(&count)
	return count != 0
}

func CheckPhoneUniqueAll(phoneNumber string) bool {
	var count int
	_ = db.Conn().QueryRow(checkPhoneUniqueAll, phoneNumber).Scan(&count)
	return count != 0
}

// 更新角色数据
func DelAddUserRoles(tx *sql.Tx, list []UserRoleEntity, userId int) error {
	_, err := tx.Exec(delAddUserRoles, userId)
	if err != nil {
		_ = tx.Rollback()
		return err
	}
	return AddUserRole(tx, list)
}

// 添加岗位与用户关联数据
func AddUserPost(tx *sql.Tx, list []UserPostEntity) error {
	stmt, err := tx.Prepare(addUserPost)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	var e error
	for i := 0; i < len(list); i++ {
		_, err = stmt.Exec(list[i].UserId, list[i].PostId)
		if err != nil {
			_ = tx.Rollback()
			e = err
			break
		}
	}

	return e
}

// 更新岗位数据
func DelAddUserPosts(tx *sql.Tx, list []UserPostEntity, userId int) error {
	_, err := tx.Exec(delAddUserPosts, userId)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	return AddUserPost(tx, list)
}

// 更新修改数据
func (r *Entity) UpdateSave(tx *sql.Tx) error {
	_, err := tx.Exec(updateSave, r.UserName, r.Email, r.PhoneNumber, r.Status, r.Sex, r.DeptId, r.Remark, r.UpdateTime, r.UpdateBy, r.UserId)
	if err != nil {
		_ = tx.Rollback()
	}
	return err
}

//检查邮箱是否存在,存在返回true,否则false
func (r *Entity) CheckEmailUnique() bool {
	var count int
	_ = db.Conn().QueryRow(checkEmailUnique, r.Email, r.UserId).Scan(&count)
	if count > 0 {
		return true
	} else {
		return false
	}
}

//检查手机号是否已使用,存在返回true,否则false
func (r *Entity) CheckPhoneUnique() bool {
	var count int
	_ = db.Conn().QueryRow(checkPhoneUnique, r.PhoneNumber, r.UserId).Scan(&count)
	if count > 0 {
		return true
	} else {
		return false
	}
}

// 根据主键查询用户信息
func SelectRecordById(id int) (e Entity) {
	_ = db.Conn().QueryRow(selectRecordById, id).Scan(&e.UserId, &e.PhoneNumber, &e.LoginName, &e.UserName, &e.Email, &e.Sex, &e.CreateTime, &e.Avatar, &e.DeptId, &e.Status, &e.Remark)
	return
}

// 删除用户角色信息
func DeleteUserRoleInfos(roleId int, users []interface{}) int {
	query := fmt.Sprintf(deleteUserRoleInfos, base.Placeholders(len(users)))
	users = append([]interface{}{roleId}, users...)

	rs, _ := db.Conn().Exec(query, users[:]...)
	count, err := rs.RowsAffected()
	if err != nil {
		return 0
	}
	return int(count)
}

// 添加用户角色关系表
func AddUserRole(tx *sql.Tx, list []UserRoleEntity) error {
	stmt, err := tx.Prepare(addUserRole)
	if err != nil {
		_ = tx.Rollback()
		return err
	}

	var e error
	for i := 0; i < len(list); i++ {
		_, err = stmt.Exec(list[i].UserId, list[i].RoleId)
		if err != nil {
			_ = tx.Rollback()
			e = err
			break
		}
	}

	return e
}

// 根据条件分页查询未分配用户角色列表
func SelectUnAllocatedList(param *role.AllocatedReq) (list []Entity) {
	if param == nil {
		return nil
	}
	var params []interface{}
	query := selectUnAllocatedList
	params = append(params, param.RoleId)

	if param.LoginName != "" {
		query += ` and u.login_name like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.LoginName)
	}
	if param.PhoneNumber != "" {
		query += ` and u.phonenumber like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.PhoneNumber)
	}
	if param.OrderByColumn != "" {
		query += ` order by ` + param.OrderByColumn + ` ` + param.IsAsc
	}

	rows, err := db.Conn().Query(query, params[:]...)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var e Entity
		_ = rows.Scan(&e.UserId, &e.DeptId, &e.LoginName, &e.UserName, &e.Email, &e.Avatar, &e.PhoneNumber, &e.Status, &e.CreateTime)
		list = append(list, e)
	}
	return
}

// 根据条件分页查询已分配用户角色列表
func SelectAllocatedList(param *role.AllocatedReq) (list []Entity) {
	if param == nil {
		return nil
	}

	var params []interface{}
	query := selectAllocatedList
	params = append(params, param.RoleId)

	if param.LoginName != "" {
		query += ` and u.login_name like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.LoginName)
	}
	if param.PhoneNumber != "" {
		query += ` and u.phonenumber like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.PhoneNumber)
	}
	if param.OrderByColumn != "" {
		query += ` order by ` + param.OrderByColumn + ` ` + param.IsAsc
	}

	rows, err := db.Conn().Query(query, params[:]...)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var e Entity
		_ = rows.Scan(&e.UserId, &e.DeptId, &e.LoginName, &e.UserName, &e.Email, &e.Avatar, &e.PhoneNumber, &e.Status, &e.CreateTime)
		list = append(list, e)
	}
	return
}

func SelectPageList(param *SelectPageReq) (list []Entity) {
	var params []interface{}

	query := selectPageList
	if param.LoginName != "" {
		query += ` and  u.login_name like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.LoginName)
	}
	if param.Phonenumber != "" {
		query += ` and u.phonenumber like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.Phonenumber)
	}
	if param.Status != "" {
		query += ` and u.status = ?`
		params = append(params, param.Status)
	}
	if param.BeginTime != "" {
		query += ` and date_format(u.create_time,'%y%m%d') >= date_format(?,'%y%m%d')`
		params = append(params, param.BeginTime)
	}
	if param.EndTime != "" {
		query += ` and date_format(u.create_time,'%y%m%d') <= date_format(?,'%y%m%d')`
		params = append(params, param.EndTime)
	}
	if param.DeptId != 0 {
		query += ` and (u.dept_id = ? OR u.dept_id IN ( SELECT t.dept_id FROM sys_dept t WHERE FIND_IN_SET (?,ancestors) ))`
		params = append(params, param.DeptId, param.DeptId)
	}
	if param.SortName != "" {
		query += ` order by u.status,` + param.SortName + ` ` + param.SortOrder
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

// 用户是否被锁
func IsUserLock(loginName string) error {
	var status string
	_ = db.Conn().QueryRow(isUserLock, loginName).Scan(&status)
	if status == "" {
		return errors.New("用户名不存在")
	} else if status == "1" {
		return errors.New("用户被锁。请联系管理员")
	}
	return nil
}

// 查询用户
func (r *Entity) FindUser() error {
	err := db.Conn().QueryRow(findOneSql, r.LoginName).Scan(&r.UserId, &r.PhoneNumber, &r.LoginName, &r.UserName, &r.Email, &r.Sex, &r.CreateTime, &r.Avatar, &r.Password, &r.Salt, &r.DeptId, &r.Status)
	switch err {
	case sql.ErrNoRows:
		return errors.New("此用户没有注册")
	case nil:
		return nil
	default:
		return errors.New("系统内部问题")
	}
}
