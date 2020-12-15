package model

import (
	"xframe/backend/common/db"
	"xframe/pkg/utils/page"
)

const (
	listPermMenuByParentId = `
	SELECT m.menu_id,m.visible,m.icon,m.name,m.url,m.parent_id,m.perms,m.menu_type
	FROM sys_role_menu rm 
	LEFT JOIN sys_menu m on rm.menu_id=m.menu_id 
	LEFT JOIN sys_user_role r on rm.role_id=r.role_id 
	LEFT JOIN sys_role rr on r.role_id=rr.role_id
	where m.visible=0 and rr.status=0 and m.parent_id=?  and r.user_id=?
	GROUP BY rm.menu_id
	ORDER BY m.menu_type,m.sort_id
	`
	listMenuByParentId     = `select menu_id,visible,icon,name,url,parent_id,perms,menu_type from sys_menu where parent_id=? order by menu_type,sort_id`
	selectPageList         = `select * from sys_menu where 1=1`
	selectMenuList         = `select menu_id,name,parent_id from sys_menu where visible='0'`
	selectRecordById       = `select m1.menu_id, m1.name, m1.url, m1.parent_id, m1.sort_id, m1.icon, m1.sys_type, m1.menu_type, m1.visible, m1.perms, IFNULL(m2.name,'') parentName from sys_menu m1 LEFT JOIN sys_menu m2 on m1.parent_id=m2.menu_id where 1=1 and m1.menu_id=?`
	editSave               = `update sys_menu set visible=?,name=?,parent_id=?,sort_id=?,url=?,menu_type=?,perms=?,icon=?,update_by=?,sys_type=? where menu_id=?`
	checkMenuNameUniqueAll = `select menu_id from sys_menu where menu_name=? and parent_id=?`
	addSave                = `insert into sys_menu(visible,name,parent_id,sort_id,url,menu_type,perms,icon,create_time,create_by,sys_type) values(?,?,?,?,?,?,?,?,?,?,?)`
	childMenuCount         = `select count(menu_id) from sys_menu where parent_id=?`
	deleteRecordById       = `delete from sys_menu where menu_id=?`
	selectMenuTree         = `SELECT concat(m.menu_id, ifnull(m.perms,'')) as perms FROM sys_menu m LEFT JOIN sys_role_menu rm on m.menu_id=rm.menu_id WHERE rm.role_id=? ORDER BY m.parent_id,m.sort_id`
)

//根据角色ID查询菜单
func SelectMenuTree(roleId int) (list []string) {
	rows, err := db.Conn().Query(selectMenuTree, roleId)
	if err != nil {
		return
	}
	defer rows.Close()

	for rows.Next() {
		var d string
		_ = rows.Scan(&d)
		list = append(list, d)
	}
	return
}

//根据主键删除数据
func DeleteRecordById(id string) error {
	_, err := db.Conn().Exec(deleteRecordById, id)
	return err
}

// 子菜单数量
func ChildMenuCount(id string) int {
	var count int
	_ = db.Conn().QueryRow(childMenuCount, id).Scan(&count)
	return count
}

func (e *Entity) AddSave() (int, error) {
	rs, err := db.Conn().Exec(addSave, e.Visible, e.Name, e.ParentId, e.SortId, e.Url, e.MenuType, e.Perms, e.Icon, e.CreateTime, e.CreateBy, e.SysType)
	if err != nil {
		return 0, err
	}
	id, err := rs.LastInsertId()
	if err != nil {
		return 0, err
	}

	return int(id), nil
}

//校验菜单名称是否唯一
func CheckMenuNameUniqueAll(menuName string, parentId int) int {
	var menuId int
	_ = db.Conn().QueryRow(checkMenuNameUniqueAll, menuName, parentId).Scan(&menuId)
	return menuId
}

func (e *Entity) EditSave() error {
	_, err := db.Conn().Exec(editSave, &e.Visible, &e.Name, &e.ParentId, &e.SortId, &e.Url, &e.MenuType, &e.Perms, &e.Icon, &e.UpdateBy, &e.SysType, &e.MenuId)
	return err
}

func SelectRecordById(id int) (e Entity) {
	_ = db.Conn().QueryRowx(selectRecordById, id).Scan(&e.MenuId, &e.Name, &e.Url, &e.ParentId, &e.SortId, &e.Icon, &e.SysType, &e.MenuType, &e.Visible, &e.Perms, &e.ParentName)
	return
}

func SelectMenuList(parentId int) (list []Entity) {
	var params []interface{}

	query := selectMenuList
	if parentId > 0 {
		query += ` and parent_id=?`
		params = append(params, parentId)
	}
	query += `order by sys_type,sort_id`

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

func SelectPageList(param *SelectPageReq) (list []Entity) {
	var params []interface{}
	query := selectPageList

	if param.MenuId != "" {
		query += ` and parent_id=?`
		params = append(params, param.MenuId)
	}
	if param.MenuName != "" {
		query += ` and name like CONCAT(CONCAT('%',?),'%')`
		params = append(params, param.MenuName)
	}
	if param.SysType != "" {
		query += ` and sys_type=?`
		params = append(params, param.SysType)
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

// 根据父ID查询子列表
func ListMenuByParentId(parentId int) (list []Entity) {
	rows, err := db.Conn().Queryx(listMenuByParentId, parentId)
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

func ListPermMenuByParentId(parentId, userId int) (list []Entity) {
	rows, err := db.Conn().Queryx(listPermMenuByParentId, parentId, userId)
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
