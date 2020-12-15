package handler

import (
	"net/http"
	"strconv"
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"
	"xframe/backend/core/system/dept/model"
	"xframe/backend/core/system/dept/service"

	"github.com/gin-gonic/gin"
)

//加载部门列表树结构的数据
func TreeData(c *gin.Context) {
	ztreeList := service.SelectDeptTree(0, "", "")
	c.JSON(http.StatusOK, ztreeList)
}

func List(c *gin.Context) {
	resp.BuildTpl(c, "core/system/dept/list").Write()
}

func ListAjax(c *gin.Context) {
	var req model.SelectPageReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("部门管理", req).Write()
		return
	}
	rows := service.SelectRecordList(&req)
	resp.BuildTable(c, req.PageReq, rows).Write()
}

func Add(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Query("pid"))
	if pid == 0 {
		pid = 100
	}

	dept := model.Entity{}
	tmp := service.SelectDeptById(pid)
	dept.ParentName = tmp.DeptName
	dept.ParentId = tmp.DeptId

	resp.BuildTpl(c, "core/system/dept/edit").Write(gin.H{"dept": dept})
}

func AddSave(c *gin.Context) {
	var req model.AddReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Add).Msg(err.Error()).Log("部门管理", req).Write()
		return
	}
	if service.CheckDeptNameUniqueAll(req.DeptName, req.ParentId) == "1" {
		resp.Error(c).Btype(constant.Buniss_Add).Msg("部门名称已存在").Log("部门管理", req).Write()
		return
	}
	if id, err := service.AddSave(c, &req); err != nil || id <= 0 {
		resp.Error(c).Btype(constant.Buniss_Add).Log("部门管理", req).Write()
	} else {
		resp.Success(c).Btype(constant.Buniss_Add).Log("部门管理", req).Write()
	}
}

func Edit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "参数错误"})
		return
	}
	if dept := service.SelectDeptById(id); dept.DeptId > 0 {
		resp.BuildTpl(c, "core/system/dept/edit").Write(gin.H{"dept": dept, "msg": "edit"})
	} else {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "部门不存在"})
	}
}

func EditSave(c *gin.Context) {
	var req model.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg(err.Error()).Log("部门管理", req).Write()
		return
	}
	if service.CheckDeptNameUnique(req.DeptName, req.DeptId, req.ParentId) == "1" {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg("部门名称已存在").Log("部门管理", req).Write()
		return
	}
	if err := service.EditSave(c, &req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Log("部门管理", req).Write()
		return
	}
	resp.Success(c).Data("1").Btype(constant.Buniss_Edit).Log("部门管理", req).Write()
}

func Remove(c *gin.Context) {
	var req constant.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Del).Msg(err.Error()).Log("部门管理", req).Write()
		return
	}
	if rs := service.DeleteRecordByIds(req.Ids); rs > 0 {
		resp.Success(c).Btype(constant.Buniss_Del).Data(rs).Log("部门管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Del).Log("部门管理", req).Write()
	}
}

//加载角色部门（数据权限）列表树
func RoleDeptTreeData(c *gin.Context) {
	roleId, _ := strconv.Atoi(c.Query("roleId"))
	if result, err := service.RoleDeptTreeData(roleId); err == nil {
		c.JSON(http.StatusOK, result)
	} else {
		resp.Error(c).Log("菜单树", gin.H{"roleId": roleId})
	}
}

//加载部门列表树选择页面
func SelectDeptTree(c *gin.Context) {
	deptId, _ := strconv.Atoi(c.Query("deptId"))
	if deptPoint := service.SelectDeptById(deptId); deptPoint.DeptId > 0 {
		resp.BuildTpl(c, "core/system/dept/tree").Write(gin.H{"dept": deptPoint})
	} else {
		resp.BuildTpl(c, "core/system/dept/tree").Write()
	}
}

//检查部门名称是否已经存在
func CheckDeptNameUnique(c *gin.Context) {
	var req model.CheckDeptNameReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	result := service.CheckDeptNameUnique(req.DeptName, req.DeptId, req.ParentId)
	_, _ = c.Writer.WriteString(result)
}

//检查部门名称是否已经存在
func CheckDeptNameUniqueAll(c *gin.Context) {
	var req model.CheckDeptNameAllReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	result := service.CheckDeptNameUniqueAll(req.DeptName, req.ParentId)
	_, _ = c.Writer.WriteString(result)
}
