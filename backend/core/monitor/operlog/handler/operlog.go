package handler

import (
	"html/template"
	"strconv"
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"
	"xframe/backend/core/monitor/operlog/model"
	"xframe/backend/core/monitor/operlog/service"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	resp.BuildTpl(c, "core/monitor/operlog/list").Write()
}

func ListAjax(c *gin.Context) {
	var req model.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("操作日志管理", req).Write()
		return
	}
	rows := service.SelectPageList(&req)
	resp.BuildTable(c, req.PageReq, rows).Write()
}

func Remove(c *gin.Context) {
	var req constant.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Del).Msg(err.Error()).Log("操作日志管理", req).Write()
		return
	}
	if rs := service.DeleteRecordByIds(req.Ids); rs > 0 {
		resp.Success(c).Btype(constant.Buniss_Del).Data(rs).Log("操作日志管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Del).Log("操作日志管理", req).Write()
	}
}

// 清空日志
func Clean(c *gin.Context) {
	if err := service.DeleteRecordAll(); err == nil {
		resp.Success(c).Btype(constant.Buniss_Del).Data(1).Log("操作日志管理", "all").Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Del).Log("操作日志管理", "all").Write()
	}
}

//记录详情
func Detail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "参数错误"})
		return
	}
	operLog := service.SelectRecordById(id)
	if operLog.OperId <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "数据不存在"})
		return
	}
	jsonResult := template.HTML(operLog.JsonResult)
	operParam := template.HTML(operLog.OperParam)
	resp.BuildTpl(c, "core/monitor/operlog/detail").Write(gin.H{
		"operLog":    operLog,
		"jsonResult": jsonResult,
		"operParam":  operParam,
	})
}
