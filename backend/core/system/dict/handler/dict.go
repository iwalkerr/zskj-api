package handler

import (
	"net/http"
	"strconv"
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"
	"xframe/backend/core/system/dict/model"
	"xframe/backend/core/system/dict/service"

	"github.com/gin-gonic/gin"
)

func TreeData(c *gin.Context) {
	ztreeList := service.SelectDictTree(0)
	c.JSON(http.StatusOK, ztreeList)
}

func List(c *gin.Context) {
	resp.BuildTpl(c, "core/system/dict/list").Write()
}

func ListAjax(c *gin.Context) {
	var req model.SelectPageReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("数据字典", req).Write()
		return
	}
	rows := service.SelectRecordList(&req)
	resp.BuildTable(c, req.PageReq, rows).Write()
}

func Add(c *gin.Context) {
	pid, _ := strconv.Atoi(c.Query("pid"))
	dict := model.Entity{IsDefault: "1", Status: "0", ParentId: pid}
	resp.BuildTpl(c, "core/system/dict/edit").Write(gin.H{"dict": dict, "msg": "add"})
}

func AddSave(c *gin.Context) {
	var req model.AddReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Add).Msg(err.Error()).Log("字典数据管理", req).Write()
		return
	}
	if rid, err := service.AddSave(c, &req); err == nil && rid > 0 {
		resp.Success(c).Data(rid).Btype(constant.Buniss_Add).Log("字典数据管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Add).Log("字典数据管理", req).Write()
	}
}

func Edit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "字典数据错误"})
		return
	}
	if entity := service.SelectRecordById(id); entity.DictId > 0 {
		resp.BuildTpl(c, "core/system/dict/edit").Write(gin.H{"dict": entity, "msg": "edit"})
	} else {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "字典数据不存在"})
	}
}

func EditSave(c *gin.Context) {
	var req model.EditReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg(err.Error()).Log("字典数据管理", req).Write()
		return
	}
	if rs, err := service.EditSave(c, &req); err == nil && rs > 0 {
		resp.Success(c).Btype(constant.Buniss_Edit).Data(rs).Log("字典数据管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Edit).Log("字典数据管理", req).Write()
	}
}

func Remove(c *gin.Context) {
	var req constant.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Del).Msg("缺少参数").Log("字典管理", req).Write()
		return
	}
	if flag := service.DeleteRecordByIds(req.Ids); !flag {
		resp.Success(c).Btype(constant.Buniss_Del).Log("字典管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Del).Msg("请先删除子字典").Log("字典管理", req).Write()
	}
}
