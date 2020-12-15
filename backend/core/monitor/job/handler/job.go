package handler

import (
	"strconv"
	"strings"
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"
	"xframe/backend/core/monitor/job/model"
	"xframe/backend/core/monitor/job/service"

	"github.com/gin-gonic/gin"
)

func List(c *gin.Context) {
	resp.BuildTpl(c, "core/monitor/job/list").Write()
}

func ListAjax(c *gin.Context) {
	var req model.SelectPageReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Msg(err.Error()).Log("定时任务管理", req).Write()
		return
	}
	rows := service.SelectListByPage(&req)
	resp.BuildTable(c, req.PageReq, rows).Write()
}

// 新增页面
func Add(c *gin.Context) {
	entity := model.Entity{}
	entity.MisfirePolicy = "1"
	resp.BuildTpl(c, "core/monitor/job/edit").Write(gin.H{"job": entity})
}

func AddSave(c *gin.Context) {
	var req model.AddReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Add).Msg(err.Error()).Log("定时任务管理", req).Write()
		return
	}

	id, err := service.AddSave(c, &req)
	if id == 0 || err != nil {
		resp.Error(c).Btype(constant.Buniss_Add).Msg(err.Error()).Log("定时任务管理", req).Write()
	} else {
		resp.Success(c).Btype(constant.Buniss_Add).Data(id).Log("定时任务管理", req).Write()
	}
}

func Edit(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "参数错误"})
		return
	}

	entity := service.SelectRecordById(id)
	if entity == nil || entity.JobId <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "数据不存在"})
		return
	}
	resp.BuildTpl(c, "core/monitor/job/edit").Write(gin.H{"job": entity, "msg": "edit"})
}

func EditSave(c *gin.Context) {
	var req model.EditReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg(err.Error()).Log("定时任务管理", req).Write()
		return
	}

	if id, err := service.EditSave(c, &req); id <= 0 || err != nil {
		resp.Error(c).Btype(constant.Buniss_Edit).Msg(err.Error()).Log("定时任务管理", req).Write()
	} else {
		resp.Success(c).Btype(constant.Buniss_Edit).Data(id).Log("定时任务管理", req).Write()
	}
}

func Remove(c *gin.Context) {
	var req constant.RemoveReq
	if err := c.ShouldBind(&req); err != nil {
		resp.Error(c).Btype(constant.Buniss_Del).Msg(err.Error()).Log("定时任务管理", req).Write()
		return
	}

	// 停止定时器
	ids := strings.Split(req.Ids, ",")
	for _, id := range ids {
		jobId, _ := strconv.Atoi(id)
		if jobId == 0 {
			continue
		}

		job := service.SelectRecordById(jobId)
		service.StopRemove(job)
	}
	resp.Success(c).Btype(constant.Buniss_Del).Log("定时任务管理", req).Write()
}

func Detail(c *gin.Context) {
	id, _ := strconv.Atoi(c.Query("id"))
	if id <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "参数错误"})
		return
	}

	job := service.SelectRecordById(id)
	if job.JobId <= 0 {
		resp.BuildTpl(c, constant.ERROR_PAGE).Write(gin.H{"desc": "数据不存在"})
		return
	}

	resp.BuildTpl(c, "core/monitor/job/detail").Write(gin.H{"job": job})
}

// 启动
func Start(c *gin.Context) {
	jobId, _ := strconv.Atoi(c.PostForm("jobId"))
	if jobId <= 0 {
		resp.Error(c).Msg("参数错误").Log("定时任务管理启动", gin.H{"jobId": jobId}).Write()
		return
	}
	job := service.SelectRecordById(jobId)
	if job.JobId <= 0 {
		resp.Error(c).Msg("任务不存在").Log("定时任务管理启动", gin.H{"jobId": jobId}).Write()
		return
	}

	if err := service.Start(job); err != nil {
		resp.Error(c).Msg(err.Error()).Log("定时任务管理启动", gin.H{"jobId": jobId}).Write()
	} else {
		resp.Success(c).Log("定时任务管理启动", gin.H{"jobId": jobId}).Write()
	}
}

// 停止
func Stop(c *gin.Context) {
	jobId, _ := strconv.Atoi(c.PostForm("jobId"))
	if jobId <= 0 {
		resp.Error(c).Msg("参数错误").Log("定时任务管理停止", gin.H{"jobId": jobId}).Write()
		return
	}
	job := service.SelectRecordById(jobId)
	if job.JobId <= 0 {
		resp.Error(c).Msg("任务不存在").Log("定时任务管理停止", gin.H{"jobId": jobId}).Write()
		return
	}

	if err := service.Stop(job); err != nil {
		resp.Error(c).Msg(err.Error()).Log("定时任务管理停止", gin.H{"jobId": jobId}).Write()
	} else {
		resp.Success(c).Msg("停止成功").Log("定时任务管理停止", gin.H{"jobId": jobId}).Write()
	}
}
