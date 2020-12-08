package handler

import (
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"
	"xframe/backend/core/monitor/loginlog/model"
	"xframe/backend/core/monitor/loginlog/service"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	resp.BuildTpl(c, "core/monitor/loginlog/list").Write()
}

func ListAjax(c *gin.Context) {
	var req *model.SelectPageReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Write()
		return
	}
	rows := service.SelectPageList(req)
	resp.BuildTable(c, req.PageReq, rows).Write()
}

func Remove(c *gin.Context) {
	var req constant.RemoveReq
	//获取参数
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Del).Msg("缺少参数").Log("登陆日志管理", req).Write()
		return
	}
	if rs := service.DeleteRecordByIds(req.Ids); rs > 0 {
		resp.Success(c).Btype(constant.Buniss_Del).Data(rs).Log("登陆日志管理", req).Write()
	} else {
		resp.Error(c).Btype(constant.Buniss_Del).Log("登陆日志管理", req).Write()
	}
}

//解锁账号
func Unlock(c *gin.Context) {
	loginName := c.Query("loginName")
	if loginName == "" {
		resp.Error(c).Msg("缺少参数").Log("解锁账号", "loginName="+loginName).Write()
	} else {
		service.RemovePasswordCounts(loginName)
		service.Unlock(loginName)
		resp.Success(c).Btype(constant.Buniss_Edit).Log("解锁账号", "loginName="+loginName).Write()
	}
}
