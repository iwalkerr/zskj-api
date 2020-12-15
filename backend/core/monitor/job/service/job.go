package service

import (
	"errors"
	"strings"
	"time"
	"xframe/backend/core/monitor/job/model"
	userService "xframe/backend/core/system/user/service"
	"xframe/backend/core/task"
	"xframe/pkg/cron"

	"github.com/gin-gonic/gin"
)

// 初始化执行的方法
func init() {
	initTask()
}

func initTask() {
	// 获取所有方法
	list := model.SelectListAll()
	for _, entity := range list {
		if entity.Status == "1" {
			continue
		}
		f := task.GetByName(entity.JobName)
		if f == nil || f.FuncName == "" {
			continue
		}

		if rs := cron.Get(entity.JobName); rs == nil {
			//传参
			paramArr := strings.Split(entity.JobParams, "|")
			task.EditParams(f.FuncName, paramArr)

			if entity.MisfirePolicy == "1" { // 多次执行
				if j, err := cron.New(&entity.Entity, f.Run); err != nil && j == nil {
					continue
				}
			} else { // 只执行一次
				f.Run()
			}
		}
	}
}

// 保存修改
func EditSave(c *gin.Context, req *model.EditReq) (int, error) {
	// 检查任务名称是否存在
	if tmp := cron.Get(req.JobName); tmp != nil {
		tmp.Stop()
	}
	// 检查task目录下是否绑定对应的方法
	if f := task.GetByName(req.JobName); f == nil {
		return 0, errors.New("当前task目录下没有绑定这个方法")
	}

	entity := &model.Entity{}
	entity.JobId = req.JobId
	entity = model.SelectRecordById(entity.JobId)
	if entity.JobId <= 0 {
		return 0, errors.New("数据不存在")
	}

	entity.JobName = req.JobName
	entity.JobParams = req.JobParams
	entity.JobGroup = req.JobGroup
	entity.InvokeTarget = req.InvokeTarget
	entity.CronExpression = req.CronExpression
	entity.MisfirePolicy = req.MisfirePolicy
	// entity.Concurrent = req.Concurrent
	entity.Remark = req.Remark

	user := userService.GetProfile(c)
	if user != nil {
		entity.UpdateBy = user.LoginName
	}

	if err := entity.Update(); err != nil {
		return 0, err
	}

	return entity.JobId, nil
}

// 添加数据
func AddSave(c *gin.Context, req *model.AddReq) (int, error) {
	user := userService.GetProfile(c)

	// 检查任务名称是否存在
	if rs := cron.Get(req.JobName); rs != nil {
		return 0, errors.New("任务名称已经存在")
	}

	// 检查task目录下是否绑定对应的方法
	if f := task.GetByName(req.JobName); f == nil {
		return 0, errors.New("当前task目录下没有绑定这个方法")
	}

	var entity model.Entity
	entity.JobName = req.JobName
	entity.JobParams = req.JobParams
	entity.JobGroup = req.JobGroup
	entity.InvokeTarget = req.InvokeTarget
	entity.CronExpression = req.CronExpression
	entity.MisfirePolicy = req.MisfirePolicy
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()
	entity.UpdateTime = time.Now()
	entity.Status = "0"
	entity.Concurrent = "1" // 禁止并发
	entity.CreateBy = user.LoginName

	return entity.Insert()
}

func SelectListByPage(req *model.SelectPageReq) []model.Entity {
	return model.SelectListByPage(req)
}

func SelectRecordById(id int) *model.Entity {
	return model.SelectRecordById(id)
}

func StopRemove(entity *model.Entity) {
	// 删除数据
	_ = model.DeleteRecord(entity.JobId)

	f := task.GetByName(entity.JobName)
	if f == nil || f.FuncName == "" {
		return
	}

	// 检查任务名称是否存在
	tmp := cron.Get(entity.JobName)
	if tmp != nil {
		tmp.Stop()
	}
}

func Stop(entity *model.Entity) error {
	f := task.GetByName(entity.JobName)
	if f == nil || f.FuncName == "" {
		return errors.New("当前task目录下没有绑定这个方法")
	}

	// 检查任务名称是否存在
	tmp := cron.Get(entity.JobName)
	if tmp != nil {
		tmp.Stop()
	}

	entity.Status = "1"
	return model.UpdateStatus(entity.JobId, entity.Status)
}

// 启动任务
func Start(entity *model.Entity) error {
	f := task.GetByName(entity.JobName)
	if f == nil || f.FuncName == "" {
		return errors.New("当前task目录下没有绑定这个方法")
	}

	if rs := cron.Get(entity.JobName); rs == nil {
		//传参
		paramArr := strings.Split(entity.JobParams, "|")
		task.EditParams(f.FuncName, paramArr)

		if entity.MisfirePolicy == "1" { // 多次执行
			if j, err := cron.New(&entity.Entity, f.Run); err != nil && j == nil {
				return err
			}

			entity.Status = "0"
			_ = model.UpdateStatus(entity.JobId, entity.Status)
		} else { // 只执行一次
			f.Run()
		}
	} else {
		return errors.New("任务已存在")
	}

	return nil
}
