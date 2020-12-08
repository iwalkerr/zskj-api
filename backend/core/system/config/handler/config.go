package handler

import (
	"strconv"
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"
	"xframe/backend/core/system/config/model"
	"xframe/backend/core/system/config/service"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	resp.BuildTpl(c, "core/system/config/list").Write()

}

func ListAjax(c *gin.Context) {
	var req model.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("参数管理", req).Write()
		return
	}
	rows := service.SelectListByPage(&req)
	resp.BuildTable(c, req.PageReq, rows).Write()
}

func Add(c *gin.Context) {
	config := model.Entity{}
	resp.BuildTpl(c, "core/system/config/edit").Write(gin.H{"config": config})
}

func AddSave(c *gin.Context) {
	var req model.AddReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Add).Msg(err.Error()).Log("参数管理", req).Write()
		return
	}
	if service.CheckConfigKeyUniqueAll(req.ConfigKey) == "1" {
		resp.Error(c).Btype(constant.Buniss_Add).Msg("参数键名已存在").Log("参数管理", req).Write()
		return
	}
	if rid, err := service.AddSave(c, &req); err == nil && rid > 0 {
		resp.Success(c).Data(rid).Log("参数管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Add).Log("参数管理", req).Write()
	}
}

func Edit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "参数错误"})
		return
	}
	if entity := service.SelectRecordById(id); entity.ConfigId > 0 {
		resp.BuildTpl(c, "core/system/config/edit").Write(gin.H{"config": entity, "msg": "edit"})
	} else {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "数据不存在"})
	}
}

func EditSave(c *gin.Context) {
	var req model.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg(err.Error()).Log("参数管理", req).Write()
		return
	}
	if service.CheckConfigKeyUnique(req.ConfigKey, req.ConfigId) == "1" {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg("参数键名已存在").Log("参数管理", req).Write()
		return
	}
	if rs, err := service.EditSave(c, &req); err == nil && rs > 0 {
		resp.Success(c).Btype(constant.Buniss_Edit).Log("参数管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Edit).Log("参数管理", req).Write()
	}
}

func Remove(c *gin.Context) {
	var req constant.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Del).Msg(err.Error()).Log("参数管理", req).Write()
		return
	}
	if err := service.DeleteRecordByIds(req.Ids); err == nil {
		resp.Success(c).Btype(constant.Buniss_Del).Log("参数管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Del).Log("参数管理", req).Write()
	}
}

//检查参数键名是否已经存在不包括本参数
func CheckConfigKeyUniqueAll(c *gin.Context) {
	var req model.CheckConfigKeyAllReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	result := service.CheckConfigKeyUniqueAll(req.ConfigKey)
	_, _ = c.Writer.WriteString(result)
}

//校验参数键名是否唯一
func CheckConfigKeyUnique(c *gin.Context) {
	var req model.CheckConfigKeyReq
	if err := c.ShouldBind(&req); err != nil {
		_, _ = c.Writer.WriteString("1")
		return
	}
	result := service.CheckConfigKeyUnique(req.ConfigKey, req.ConfigId)
	_, _ = c.Writer.WriteString(result)
}
