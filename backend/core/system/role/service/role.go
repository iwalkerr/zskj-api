package service

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"xframe/backend/common/db"
	"xframe/backend/core/system/role/model"
	user "xframe/backend/core/system/user/model"
	userService "xframe/backend/core/system/user/service"

	"github.com/gin-gonic/gin"
)

//根据用户ID查询角色
func SelectRoleContactVo(userId int) []model.Entity {
	var paramRole model.SelectPageReq
	roleAll := model.SelectListAll(&paramRole)
	if len(roleAll) == 0 {
		return nil
	}

	userRole := model.SelectRoleContactVo(userId)
	if userRole == nil {
		return roleAll
	}

	for i := range roleAll {
		for j := range userRole {
			if userRole[j].RoleId == roleAll[i].RoleId {
				roleAll[i].Flag = true
				break
			}
		}
	}

	return roleAll
}

//根据条件查询数据
func SelectRecordAll(params *model.SelectPageReq) []model.Entity {
	return model.SelectListAll(params)
}

//批量取消授权用户角色
func DeleteUserRoleInfos(roleId int, userIds string) int {
	idarr := strings.Split(userIds, ",")

	ids := make([]interface{}, len(idarr))
	for i, id := range idarr {
		ids[i] = id
	}

	return user.DeleteUserRoleInfos(roleId, ids)
}

//批量选择用户授权
func InsertAuthUsers(param *model.UserRoleEntity) int {
	idarr := strings.Split(param.UserIds, ",")
	var roleUserList []user.UserRoleEntity
	for _, id := range idarr {
		userId, _ := strconv.Atoi(id)
		var tmp user.UserRoleEntity
		tmp.UserId = userId
		tmp.RoleId = param.RoleId
		roleUserList = append(roleUserList, tmp)
	}

	// 插入授权用户
	tx, err := db.Conn().Begin()
	if err != nil {
		return 0
	}
	if err := user.AddUserRole(tx, roleUserList); err != nil {
		return 0
	}
	if err := tx.Commit(); err != nil {
		return 0
	}
	return param.RoleId
}

//保存数据权限
func AuthDataScope(c *gin.Context, req *model.DataScopeReq) int {
	r := model.SelectRecordById(req.RoleId)
	if r.RoleId <= 0 {
		return 0
	}
	if req.DataScope != "" {
		r.DataScope = req.DataScope
	}

	u := userService.GetProfile(c)
	if u != nil {
		r.UpdateBy = u.LoginName
	}
	r.UpdateTime = time.Now()

	tx, err := db.Conn().Begin()
	if err != nil {
		return 0
	}

	if err := r.EditSave(tx); err != nil {
		return 0
	}

	if req.DeptIds != "" {
		deptIds := strings.Split(req.DeptIds, ",")
		if len(deptIds) > 0 {
			roleDepts := make([]model.EntityRoleDept, 0)
			for _, deptId := range deptIds {
				id, _ := strconv.Atoi(deptId)
				if id > 0 {
					var tmp model.EntityRoleDept
					tmp.RoleId = r.RoleId
					tmp.DeptId = id
					roleDepts = append(roleDepts, tmp)
				}
			}
			if len(roleDepts) > 0 {
				// 删除角色部门
				if err := model.DelRoleDeptById(tx, r.RoleId); err != nil {
					return 0
				}
				// 添加角色部门
				if err := r.AddRoleDepts(tx, roleDepts); err != nil {
					return 0
				}
			}
		}
	} else {
		// 删除角色部门
		if err := model.DelRoleDeptById(tx, r.RoleId); err != nil {
			return 0
		}
	}

	if err := tx.Commit(); err != nil {
		return 0
	}
	return 1
}

//判断是否是管理员
func IsAdmin(id int) bool {
	return id == 1
}

//校验角色是否允许操作
func CheckRoleAllAllowed(id int) bool {
	return !IsAdmin(id)
}

//批量删除数据记录
func DeleteRecordByIds(ids string) int {
	idArr := strings.Split(ids, ",")
	if err := model.DeleteBatch(idArr); err != nil {
		return 0
	}
	return 1
}

// 改变角色状态
func ChangeStatus(c *gin.Context, req *model.ChangeStatusReq) error {
	var r model.Entity
	r.RoleId = req.RoleId
	r.Status = req.Status

	if err := r.ChangeStatus(); err != nil {
		return errors.New("角色状态更新失败")
	}
	return nil
}

// 保存编辑数据
func EditSave(c *gin.Context, req *model.EditReq) int {
	// 查询角色是否存在
	r := model.SelectRecordById(req.RoleId)
	if r.RoleId <= 0 {
		return 0
	}

	roleSort, _ := strconv.Atoi(req.RoleSort)

	r.RoleName = req.RoleName
	r.RoleKey = req.RoleKey
	r.Status = req.Status
	r.Remark = req.Remark
	r.RoleSort = roleSort
	r.UpdateTime = time.Now()

	u := userService.GetProfile(c)
	if u != nil {
		r.CreateBy = u.LoginName
	}

	tx, err := db.Conn().Begin()
	if err != nil {
		return 0
	}

	// 保存数据
	if err := r.EditSave(tx); err != nil {
		return 0
	}

	if req.MenuIds != "" {
		menus := strings.Split(req.MenuIds, ",")
		if len(menus) > 0 {
			roleMenus := make([]model.EntityRoleMenu, 0)
			for _, menuId := range menus {
				id, _ := strconv.Atoi(menuId)
				if id > 0 {
					var tmp model.EntityRoleMenu
					tmp.RoleId = r.RoleId
					tmp.MenuId = id
					roleMenus = append(roleMenus, tmp)
				}
			}
			if len(roleMenus) > 0 {
				if err := model.DelRoleMenuById(tx, r.RoleId); err != nil {
					return 0
				}
				var rm model.Entity
				if err := rm.AddRoleMenus(tx, roleMenus); err != nil {
					return 0
				}
			}

		}
	} else {
		if err := model.DelRoleMenuById(tx, r.RoleId); err != nil {
			return 0
		}
	}

	if err := tx.Commit(); err != nil {
		return 0
	}
	return r.RoleId
}

//检查角色名是否唯一
func CheckRoleNameUnique(roleName string, roleId int) string {
	if flag := model.CheckRoleNameUnique(roleName, roleId); flag {
		return "0"
	}
	return "1"
}

//检查角色键是否唯一
func CheckRoleKeyUnique(roleKey string, roleId int) string {
	if flag := model.CheckRoleKeyUnique(roleKey, roleId); flag {
		return "0"
	}
	return "1"
}

//根据主键查询数据
func SelectRecordById(id int) model.Entity {
	return model.SelectRecordById(id)
}

func SelectRecordPage(req *model.SelectPageReq) []model.Entity {
	return model.SelectListPage(req)
}

//检查角色名是否唯一
func CheckRoleNameUniqueAll(roleName string) string {
	if model.CheckRoleNameUniqueAll(roleName) {
		return "1"
	}
	return "0"
}

//检查角色键是否唯一
func CheckRoleKeyUniqueAll(roleKey string) string {
	if model.CheckRoleKeyUniqueAll(roleKey) {
		return "1"
	}
	return "0"
}

//添加数据
func AddSave(c *gin.Context, req *model.AddReq) int {
	r := model.Entity{}
	r.RoleName = req.RoleName
	r.RoleKey = req.RoleKey
	r.Status = req.Status
	r.RoleSort = req.RoleSort
	r.Remark = req.Remark
	r.CreateTime = time.Now()
	r.DelFlag = "0"
	r.DataScope = "5"

	user := userService.GetProfile(c)
	if user != nil {
		r.CreateBy = user.LoginName
	}

	tx, err := db.Conn().Begin()
	if err != nil {
		return 0
	}

	id, err := r.Add(tx)
	if err != nil {
		return 0
	}
	r.RoleId = id

	if req.MenuIds != "" {
		menus := strings.Split(req.MenuIds, ",")
		if len(menus) > 0 {
			roleMenus := make([]model.EntityRoleMenu, 0)
			for _, menuId := range menus {
				id, _ := strconv.Atoi(menuId)
				if id > 0 {
					var tmp model.EntityRoleMenu
					tmp.RoleId = r.RoleId
					tmp.MenuId = id
					roleMenus = append(roleMenus, tmp)
				}
			}
			if len(roleMenus) > 0 {
				var rm model.Entity
				if err := rm.AddRoleMenus(tx, roleMenus); err != nil {
					return 0
				}
			}
		}
	}

	if err := tx.Commit(); err != nil {
		return 0
	}

	return id
}
