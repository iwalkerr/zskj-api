package model

import (
	"fmt"
	"strconv"
	"strings"
	"xframe/backend/common/db"
	"xframe/pkg/utils/base"
	"xframe/pkg/utils/page"
)

const (
	deptList               = `select dept_id,dept_name,status,parent_id,create_time,order_num from sys_dept where del_flag='0'`
	selectPageList         = `select * from sys_dept where 1=1 and del_flag='0'`
	selectDeptList         = `select dept_id,dept_name,status,parent_id,create_time,order_num from sys_dept d where d.del_flag='0'`
	selectRoleDeptTree     = `SELECT CONCAT(d.dept_id,d.dept_name) dept_name from sys_dept d LEFT JOIN sys_role_dept rd on d.dept_id=rd.dept_id WHERE d.del_flag='0' and rd.role_id=? ORDER BY d.parent_id,order_num`
	selectDepById          = `select d.dept_id, d.parent_id, d.ancestors, d.dept_name, d.order_num, d.leader, d.phone, d.email, d.status,(select dept_name from sys_dept where dept_id = d.parent_id) parent_name from sys_dept d where d.dept_id=?`
	checkDeptNameUniqueAll = `select dept_id from sys_dept where dept_name=? and parent_id=?`
	selectChildrenDeptById = `SELECT dept_id,ancestors FROM sys_dept WHERE find_in_set(?, ancestors)`
	update                 = `update sys_dept set dept_name=?,status=?,parent_id=?,del_flag=?,email=?,leader=?,phone=?,order_num=?,update_by=?,update_time=? where dept_id=?`
	insert                 = `insert into sys_dept(ancestors,dept_name,status,parent_id,del_flag,email,leader,phone,order_num,create_by,create_time) values(?,?,?,?,?,?,?,?,?,?,?)`
	deleteBatch            = `delete from sys_dept where dept_id in (%s)`
	deptNameById           = `select dept_name from sys_dept where dept_id=?`
)

// 根据ID获取部门名称
func DeptNameById(deptId int) (deptName string) {
	_ = db.Conn().QueryRow(deptNameById, deptId).Scan(&deptName)
	return
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

// 插入数据
func (e *Entity) Insert() (int, error) {
	rs, _ := db.Conn().Exec(insert, e.Ancestors, e.DeptName, e.Status, e.ParentId, e.DelFlag, e.Email, e.Leader, e.Phone, e.OrderNum, e.CreateBy, e.CreateTime)
	id, err := rs.LastInsertId()
	if err != nil {
		return 0, err
	}
	return int(id), nil
}

// 更新数据
func (e *Entity) Update() error {
	_, err := db.Conn().Exec(update, e.DeptName, e.Status, e.ParentId, e.DelFlag, e.Email, e.Leader, e.Phone, e.OrderNum, e.UpdateBy, e.UpdateTime, e.DeptId)
	return err
}

//根据ID查询所有子部门
func SelectChildrenDeptById(deptId int) (list []Entity) {
	rows, err := db.Conn().Query(selectChildrenDeptById, deptId)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var d Entity
		_ = rows.Scan(&d.DeptId, &d.Ancestors)
		list = append(list, d)
	}
	return
}

//修改子元素关系
func UpdateDeptChildren(deptId int, newAncestors, oldAncestors string) {
	deptList := SelectChildrenDeptById(deptId)
	if len(deptList) == 0 {
		return
	}
	for _, tmp := range deptList {
		tmp.Ancestors = strings.ReplaceAll(tmp.Ancestors, oldAncestors, newAncestors)
	}
	ancestors := " case dept_id"
	idStr := ""
	for _, dept := range deptList {
		ancestors += " when " + strconv.Itoa(dept.DeptId) + " then '" + dept.Ancestors + "'"
		if idStr == "" {
			idStr = strconv.Itoa(dept.DeptId)
		} else {
			idStr += "," + strconv.Itoa(dept.DeptId)
		}
	}
	ancestors += " end"
	updateDeptChildren := fmt.Sprintf(`update sys_dept set ancestors=%s where dept_id in (%s)`, ancestors, idStr)
	rs, err := db.Conn().Exec(updateDeptChildren)
	if err != nil {
		fmt.Printf("修改了%v行 错误信息：%v", rs, err.Error())
	}
}

//校验部门名称是否唯一
func CheckDeptNameUniqueAll(deptName string, parentId int) int {
	var id int
	_ = db.Conn().QueryRow(checkDeptNameUniqueAll, deptName, parentId).Scan(&id)
	return id
}

//根据部门ID查询信息
func SelectDepById(deptId int) (d Entity) {
	_ = db.Conn().QueryRow(selectDepById, deptId).Scan(&d.DeptId, &d.ParentId, &d.Ancestors, &d.DeptName, &d.OrderNum, &d.Leader, &d.Phone, &d.Email, &d.Status, &d.ParentName)
	return
}

//根据角色ID查询部门
func SelectRoleDeptTree(roleId int) (list []string) {
	rows, err := db.Conn().Query(selectRoleDeptTree, roleId)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var deptName string
		_ = rows.Scan(&deptName)
		list = append(list, deptName)
	}
	return
}

//查询部门管理数据
func SelectDeptList(parentId int, deptName, status string) (list []Entity, err error) {
	var params []interface{}
	query := selectDeptList
	if parentId > 0 {
		query += ` and d.parent_id = ? `
		params = append(params, parentId)
	}
	if deptName != "" {
		query += ` and d.dept_name like CONCAT(CONCAT('%',?),'%')`
		params = append(params, deptName)
	}
	if status != "" {
		query += ` and d.status = ?`
		params = append(params, status)
	}
	query += ` order by d.parent_id, d.order_num`

	rows, err := db.Conn().Query(query, params[:]...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		var d Entity
		_ = rows.Scan(&d.DeptId, &d.DeptName, &d.Status, &d.ParentId, &d.CreateTime, &d.OrderNum)
		list = append(list, d)
	}
	return
}

func SelectPageList(param *SelectPageReq) (list []Entity) {
	var params []interface{}
	query := selectPageList

	if param.DeptId != "" {
		query += ` and parent_id=?`
		params = append(params, param.DeptId)
	}
	if param.DeptName != "" {
		query += ` and dept_name like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.DeptName)
	}
	if param.Visible != "" {
		query += ` and visible=?`
		params = append(params, param.Visible)
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

//查询部门管理数据
func DeptList(parentId int, deptName, status string) (list []Entity) {
	var params []interface{}

	query := deptList
	if parentId > 0 {
		query += ` and parent_id=?`
		params = append(params, parentId)
	}
	if deptName != "" {
		query += ` and dept_name like CONCAT(CONCAT('%',?),'%')`
		params = append(params, deptName)
	}
	if status != "" {
		query += ` and status=?`
		params = append(params, status)
	}
	query += ` order by parent_id,order_num`

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
