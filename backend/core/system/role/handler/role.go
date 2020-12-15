package handler

import (
	"net/http"
	"strconv"
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"
	"xframe/backend/core/system/role/model"
	"xframe/backend/core/system/role/service"
	userService "xframe/backend/core/system/user/service"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	resp.BuildTpl(c, "core/system/role/list").Write()
}

func ListAjax(c *gin.Context) {
	var req model.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("角色管理", req).Write()
		return
	}
	rows := service.SelectRecordPage(&req)
	resp.BuildTable(c, req.PageReq, rows).Write()
}

func Add(c *gin.Context) {
	role := model.Entity{Status: "0"}
	resp.BuildTpl(c, "core/system/role/edit").Write(gin.H{"role": role})
}

func AddSave(c *gin.Context) {
	var req model.AddReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Add).Msg(err.Error()).Log("角色管理", req).Write()
		return
	}
	if service.CheckRoleNameUniqueAll(req.RoleName) == "1" {
		resp.Error(c).Btype(constant.Buniss_Add).Msg("角色名称已存在").Log("角色管理", req).Write()
		return
	}
	if service.CheckRoleKeyUniqueAll(req.RoleKey) == "1" {
		resp.Error(c).Btype(constant.Buniss_Add).Msg("角色权限已存在").Log("角色管理", req).Write()
		return
	}
	if rid := service.AddSave(c, &req); rid > 0 {
		resp.Success(c).Data(rid).Btype(constant.Buniss_Add).Log("角色管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Add).Log("角色管理", req).Write()
	}
}

func Edit(c *gin.Context) {
	id, err := strconv.Atoi(c.Query("id"))
	if id <= 0 || err != nil {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "参数错误"})
		return
	}
	if role := service.SelectRecordById(id); role.RoleId > 0 {
		resp.BuildTpl(c, "core/system/role/edit").Write(gin.H{"role": role, "msg": "edit"})
	} else {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "角色不存在"})
	}
}

func EditSave(c *gin.Context) {
	var req model.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg(err.Error()).Log("角色管理", req).Write()
		return
	}
	if service.CheckRoleKeyUnique(req.RoleKey, req.RoleId) == "1" {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg("角色权限已存在").Log("角色管理", req).Write()
		return
	}
	if service.CheckRoleNameUnique(req.RoleName, req.RoleId) == "1" {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg("角色名称已存在").Log("角色管理", req).Write()
		return
	}
	// 保存数据
	if rs := service.EditSave(c, &req); rs > 0 {
		resp.Success(c).Btype(constant.Buniss_Edit).Data(rs).Log("角色管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Edit).Log("角色管理", req).Write()
	}
}

func Remove(c *gin.Context) {
	var req constant.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Del).Msg(err.Error()).Log("角色管理", req).Write()
		return
	}
	if rs := service.DeleteRecordByIds(req.Ids); rs > 0 {
		resp.Success(c).Btype(constant.Buniss_Del).Data(rs).Log("角色管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Del).Log("角色管理", req).Write()
	}
}

// 检查角色key的唯一
func CheckRoleKeyUnique(c *gin.Context) {
	var req model.CheckRoleKeyReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	result := service.CheckRoleKeyUnique(req.RoleKey, req.RoleId)
	_, _ = c.Writer.WriteString(result)
}

// 检查角色名字的唯一
func CheckRoleNameUnique(c *gin.Context) {
	var req model.CheckRoleNameReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	result := service.CheckRoleNameUnique(req.RoleName, req.RoleId)
	_, _ = c.Writer.WriteString(result)
}

//检查角色是否已经存在
func CheckRoleKeyUniqueAll(c *gin.Context) {
	var req model.CheckRoleKeyAllReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	result := service.CheckRoleKeyUniqueAll(req.RoleKey)
	_, _ = c.Writer.WriteString(result)
}

// 检查角色是否已经存在
func CheckRoleNameUniqueAll(c *gin.Context) {
	var req model.CheckRoleNameAllReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	result := service.CheckRoleNameUniqueAll(req.RoleName)
	_, _ = c.Writer.WriteString(result)
}

// 改变角色状态
func ChangeStatus(c *gin.Context) {
	var req model.ChangeStatusReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("用户管理", req).Write()
		return
	}
	if err := service.ChangeStatus(c, &req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg(err.Error()).Log("用户管理", req).Write()
	} else {
		resp.Success(c).Btype(constant.Buniss_Edit).Log("修改用户状态", req).Write()
	}
}

//数据权限保存
func AuthDataScopeSave(c *gin.Context) {
	var req model.DataScopeReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("角色管理", req).Write()
		return
	}
	if !service.CheckRoleAllAllowed(req.RoleId) {
		resp.Error(c).Msg("不允许操作超级管理员角色").Log("角色管理", req).Write()
		return
	}
	if rs := service.AuthDataScope(c, &req); rs <= 0 {
		resp.Error(c).Msg("保存数据失败").Log("角色管理", req).Write()
	} else {
		resp.Success(c).Log("角色管理", req).Write()
	}
}

//数据权限
func AuthDataScope(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Query("id"))
	if role := service.SelectRecordById(roleId); role.RoleId > 0 {
		resp.BuildTpl(c, "core/system/role/dataScope").Write(gin.H{"role": role})
	} else {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "角色不存在"})
	}
}

// 用户授权
func AuthUser(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Query("id"))
	if role := service.SelectRecordById(roleId); role.RoleId > 0 {
		resp.BuildTpl(c, "core/system/role/authUser").Write(gin.H{"role": role})
	} else {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "角色不存在"})
	}
}

//查询已分配用户角色列表
func AllocatedList(c *gin.Context) {
	var req model.AllocatedReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("角色管理", req).Write()
		return
	}
	rows := userService.SelectAllocatedList(&req)
	c.JSON(http.StatusOK, constant.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}

//获取用户列表
func UnallocatedList(c *gin.Context) {
	var req model.AllocatedReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("角色管理", req).Write()
		return
	}
	rows := userService.SelectUnAllocatedList(&req)
	c.JSON(http.StatusOK, constant.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}

//分配用户添加
func SelectUser(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "参数错误"})
		return
	}
	if role := service.SelectRecordById(id); role.RoleId > 0 {
		resp.BuildTpl(c, "core/system/role/selectUser").Write(gin.H{"role": role})
	} else {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "角色不存在"})
	}
}

//保存角色选择
func SelectAll(c *gin.Context) {
	var req model.UserRoleEntity
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("参数错误", req).Write()
		return
	}
	if rs := service.InsertAuthUsers(&req); rs <= 0 {
		resp.Error(c).Btype(constant.Buniss_Add).Log("角色管理", req).Write()
	} else {
		resp.Success(c).Btype(constant.Buniss_Add).Log("角色管理", req).Write()
	}
}

//批量取消用户角色授权
func CancelAll(c *gin.Context) {
	var req model.UserRoleEntity
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("参数错误", req).Write()
		return
	}
	if req.RoleId <= 0 || req.UserIds == "" {
		resp.Error(c).Btype(constant.Buniss_Del).Msg("未找到角色").Log("角色管理", req).Write()
		return
	}
	if rs := service.DeleteUserRoleInfos(req.RoleId, req.UserIds); rs <= 0 {
		resp.Error(c).Btype(constant.Buniss_Del).Msg("取消授权失败").Log("角色管理", req).Write()
	} else {
		resp.Success(c).Btype(constant.Buniss_Del).Log("角色管理", req).Write()
	}
}

// 取消用户角色授权
func Cancel(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.PostForm("roleId"))
	userId := c.PostForm("userId")
	if roleId <= 0 || userId == "" {
		resp.Error(c).Btype(constant.Buniss_Del).Msg("未找到角色").Log("角色管理", gin.H{"roleId": roleId, "userId": userId}).Write()
		return
	}
	if rs := service.DeleteUserRoleInfos(roleId, userId); rs <= 0 {
		resp.Error(c).Btype(constant.Buniss_Del).Msg("取消授权失败").Log("角色管理", gin.H{"roleId": roleId, "userId": userId}).Write()
	} else {
		resp.Success(c).Btype(constant.Buniss_Del).Log("角色管理", gin.H{"roleId": roleId, "userId": userId}).Write()
	}
}
