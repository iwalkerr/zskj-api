package handler

import (
	"strconv"
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"
	"xframe/backend/core/system/post/model"
	"xframe/backend/core/system/post/service"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	resp.BuildTpl(c, "core/system/post/list").Write()
}

func ListAjax(c *gin.Context) {
	var req model.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("岗位管理", req).Write()
		return
	}
	rows := service.SelectListByPage(&req)
	resp.BuildTable(c, req.PageReq, rows).Write()
}

func Add(c *gin.Context) {
	post := model.Entity{}
	resp.BuildTpl(c, "core/system/post/edit").Write(gin.H{"post": post})
}

func AddSave(c *gin.Context) {
	var req model.AddReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Add).Msg(err.Error()).Log("岗位管理", req).Write()
		return
	}
	if service.CheckPostNameUniqueAll(req.PostName) == "1" {
		resp.Error(c).Btype(constant.Buniss_Add).Msg("岗位名称已存在").Log("岗位管理", req).Write()
		return
	}
	if service.CheckPostCodeUniqueAll(req.PostCode) == "1" {
		resp.Error(c).Btype(constant.Buniss_Add).Msg("岗位编码已存在").Log("岗位管理", req).Write()
		return
	}
	if id, err := service.AddSave(c, &req); err == nil && id > 0 {
		resp.Success(c).Data(id).Btype(constant.Buniss_Add).Log("岗位管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Add).Log("岗位管理", req).Write()
	}
}

func Edit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "参数错误"})
		return
	}
	if post := service.SelectRecordById(id); post.PostId > 0 {
		resp.BuildTpl(c, "core/system/post/edit").Write(gin.H{"post": post, "msg": "edit"})
	} else {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "岗位不存在"})
	}
}

func EditSave(c *gin.Context) {
	var req model.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg(err.Error()).Log("岗位管理", req).Write()
		return
	}
	if service.CheckPostNameUnique(req.PostName, req.PostId) == "1" {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg("岗位名称已存在").Log("岗位管理", req).Write()
		return
	}
	if service.CheckPostCodeUnique(req.PostCode, req.PostId) == "1" {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg("岗位编码已存在").Log("岗位管理", req).Write()
		return
	}
	// 保存
	if rs, err := service.EditSave(c, &req); err == nil && rs > 0 {
		resp.Success(c).Data(rs).Btype(constant.Buniss_Edit).Log("岗位管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Edit).Log("岗位管理", req).Write()
	}
}

//检查岗位名称是否已经存在不包括本岗位
func CheckPostNameUnique(c *gin.Context) {
	var req model.CheckPostNameReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	result := service.CheckPostNameUnique(req.PostName, req.PostId)
	_, _ = c.Writer.WriteString(result)
}

func CheckPostCodeUnique(c *gin.Context) {
	var req model.CheckPostCodeReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	result := service.CheckPostCodeUnique(req.PostCode, req.PostId)
	_, _ = c.Writer.WriteString(result)
}

//检查岗位编码是否已经存在
func CheckPostCodeUniqueAll(c *gin.Context) {
	var req model.CheckPostCodeAllReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	result := service.CheckPostCodeUniqueAll(req.PostCode)
	_, _ = c.Writer.WriteString(result)
}

//检查岗位名称是否已经存在
func CheckPostNameUniqueAll(c *gin.Context) {
	var req model.CheckPostNameAllReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	result := service.CheckPostNameUniqueAll(req.PostName)
	_, _ = c.Writer.WriteString(result)
}

func Remove(c *gin.Context) {
	var req constant.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Btype(constant.Buniss_Del).Log("岗位管理", req).Write()
		return
	}
	if err := service.DeleteRecordByIds(req.Ids); err == nil {
		resp.Success(c).Btype(constant.Buniss_Del).Log("岗位管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Del).Log("岗位管理", req).Write()
	}
}
