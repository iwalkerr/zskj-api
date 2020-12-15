package handler

import (
	"net/http"
	"strconv"
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"
	"xframe/backend/core/system/menu/model"
	"xframe/backend/core/system/menu/service"
	userService "xframe/backend/core/system/user/service"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	resp.BuildTpl(c, "core/system/menu/list").Write()
}

func TreeData(c *gin.Context) {
	ztreeList := service.SelectDictTree(0)
	c.JSON(http.StatusOK, ztreeList)
}

func ListAjax(c *gin.Context) {
	var req model.SelectPageReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("菜单管理", req).Write()
		return
	}
	rows := service.SelectRecordList(&req)
	resp.BuildTable(c, req.PageReq, rows).Write()
}

func Add(c *gin.Context) {
	pids := c.Query("pid")
	pid, _ := strconv.Atoi(pids)

	var pmenu model.Entity
	pmenu.MenuId = 0
	pmenu.ParentName = "主目录"
	menu := service.SelectRecordById(pid)
	if menu.MenuId > 0 {
		pmenu.MenuId = menu.MenuId
		pmenu.ParentId = menu.MenuId
		pmenu.ParentName = menu.Name
		pmenu.SysType = menu.SysType
	}

	resp.BuildTpl(c, "core/system/menu/edit").Write(gin.H{"menu": pmenu})
}

func AddSave(c *gin.Context) {
	var req model.AddReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Add).Msg(err.Error()).Log("菜单管理", req).Write()
		return
	}
	if service.CheckMenuNameUniqueAll(req.MenuName, req.ParentId) == "1" {
		resp.Error(c).Btype(constant.Buniss_Add).Msg("菜单名称已存在").Log("菜单管理", req).Write()
		return
	}

	if id, err := service.AddSave(c, &req); err == nil && id > 0 {
		resp.Success(c).Btype(constant.Buniss_Add).Data(id).Log("菜单管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Add).Msg("菜单添加错误").Log("菜单管理", req).Write()
	}
}

func Edit(c *gin.Context) {
	idi, err := strconv.Atoi(c.Query("id"))
	if err != nil || idi <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "请求参数错误！"})
		return
	}
	menu := service.SelectRecordById(idi)
	resp.BuildTpl(c, "core/system/menu/edit").Write(gin.H{"menu": menu, "msg": "edit"})
}

func EditSave(c *gin.Context) {
	var req model.EditReq
	// 获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg(err.Error()).Log("菜单管理", req).Write()
		return
	}

	// 验证数据的合法性

	if err := service.EditSave(c, &req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Log("菜单管理", req).Write()
	} else {
		resp.Success(c).Btype(constant.Buniss_Edit).Log("菜单管理", req).Write()
	}
}

func Remove(c *gin.Context) {
	var req constant.RemoveReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Del).Msg(err.Error()).Log("菜单管理", req).Write()
		return
	}

	if ismore := service.DeleteRecordByIds(req.Ids); ismore {
		resp.Error(c).Btype(constant.Buniss_Del).Msg("存在子菜单,删除失败").Write()
	} else {
		resp.Success(c).Btype(constant.Buniss_Del).Log("菜单管理", req).Write()
	}
}

//检查菜单名是否已经存在不包括自身
func CheckMenuNameUnique(c *gin.Context) {
	var req model.CheckMenuNameReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	rs := service.CheckMenuNameUnique(req.MenuName, req.MenuId, req.ParentId)
	_, _ = c.Writer.WriteString(rs)
}

//检查菜单名是否已经存在
func CheckMenuNameUniqueAll(c *gin.Context) {
	var req model.CheckMenuNameAllReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	result := service.CheckMenuNameUniqueAll(req.MenuName, req.ParentId)
	_, _ = c.Writer.WriteString(result)
}

func SelectMenuTree(c *gin.Context) {
	menuId, err := strconv.Atoi(c.Query("menuId"))
	if err != nil {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "菜单ID错误"})
		return
	}

	if menu := service.SelectRecordById(menuId); menu.MenuId > 0 {
		resp.BuildTpl(c, "core/system/menu/tree").Write(gin.H{"menu": menu})
	} else {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "菜单不存在"})
	}
}

// 加载所有菜单列表树
func MenuTreeData(c *gin.Context) {
	user := userService.GetProfile(c)
	if user == nil {
		resp.Error(c).Msg("登陆超时").Log("菜单管理", gin.H{"userId": user.UserId}).Write()
		return
	}
	ztrees := service.MenuTreeData(user.UserId)
	if ztrees == nil {
		resp.Error(c).Msg("菜单查询错误").Log("菜单管理", gin.H{"userId": user.UserId}).Write()
		return
	}
	c.JSON(http.StatusOK, ztrees)
}

func RoleMenuTreeData(c *gin.Context) {
	roleId, err := strconv.Atoi(c.Query("roleId"))
	if err != nil {
		resp.Error(c).Msg("ID转换异常").Log("菜单管理", gin.H{"roleId": roleId}).Write()
		return
	}
	user := userService.GetProfile(c)
	if user == nil || user.UserId <= 0 {
		resp.Error(c).Msg("登陆超时").Log("菜单管理", gin.H{"roleId": roleId}).Write()
		return
	}
	ztree := service.RoleMenuTreeData(roleId, user.UserId)
	if ztree == nil {
		resp.Error(c).Msg("获取菜单树失败").Log("菜单管理", gin.H{"roleId": roleId}).Write()
		return
	}
	c.JSON(http.StatusOK, ztree)
}
