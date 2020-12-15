package service

import (
	"errors"
	"strings"
	"time"
	"xframe/backend/core/system/post/model"
	userService "xframe/backend/core/system/user/service"

	"github.com/gin-gonic/gin"
)

// 根据用户ID查询岗位
func SelectPostsByUserId(userId int) []model.Entity {
	var paramsPost model.SelectPageReq

	postAll := model.SelectListAll(&paramsPost)
	userPost := model.SelectPostsByUserId(userId)
	for i := range postAll {
		for j := range userPost {
			if userPost[j].PostId == postAll[i].PostId {
				postAll[i].Flag = true
				break
			}
		}
	}
	return postAll
}

//根据条件分页查询角色数据
func SelectListAll(params *model.SelectPageReq) []model.Entity {
	return model.SelectListAll(params)
}

func SelectListByPage(req *model.SelectPageReq) []model.Entity {
	return model.SelectListByPage(req)
}

//根据主键查询数据
func SelectRecordById(id int) model.Entity {
	return model.SelectRecordById(id)
}

//检查岗位名称是否唯一
func CheckPostNameUnique(postName string, postId int) string {
	id := model.CheckPostNameUniqueAll(postName)
	if id > 0 && id != postId {
		return "1"
	}
	return "0"
}

//检查岗位编码是否已经存在不包括本岗位
func CheckPostCodeUnique(postCode string, postId int) string {
	id := model.CheckPostCodeUniqueAll(postCode)
	if id > 0 && id != postId {
		return "1"
	}
	return "0"
}

//修改数据
func EditSave(c *gin.Context, req *model.EditReq) (int, error) {
	entity := model.SelectRecordById(req.PostId)
	if entity.PostId <= 0 {
		return 0, errors.New("数据不存在")
	}
	entity.PostName = req.PostName
	entity.PostCode = req.PostCode
	entity.Status = req.Status
	entity.Remark = req.Remark
	entity.PostSort = req.PostSort
	entity.UpdateTime = time.Now()
	user := userService.GetProfile(c)
	if user != nil {
		entity.UpdateBy = user.LoginName
	}

	if err := entity.Update(); err != nil {
		return 0, err
	}
	return req.PostId, nil
}

//检查角色名是否唯一
func CheckPostNameUniqueAll(postName string) string {
	id := model.CheckPostNameUniqueAll(postName)
	if id > 0 {
		return "1"
	} else {
		return "0"
	}
}

//检查岗位编码是否唯一
func CheckPostCodeUniqueAll(postCode string) string {
	count := model.CheckPostCodeUniqueAll(postCode)
	if count > 0 {
		return "1"
	} else {
		return "0"
	}
}

//添加数据
func AddSave(c *gin.Context, req *model.AddReq) (int, error) {
	var entity model.Entity
	entity.PostName = req.PostName
	entity.PostCode = req.PostCode
	entity.Status = req.Status
	entity.PostSort = req.PostSort
	entity.Remark = req.Remark
	entity.CreateTime = time.Now()

	user := userService.GetProfile(c)
	if user != nil {
		entity.CreateBy = user.LoginName
	}

	return entity.Insert()
}

//批量删除数据记录
func DeleteRecordByIds(ids string) error {
	idarr := strings.Split(ids, ",")
	params := make([]interface{}, len(ids))
	for i, id := range idarr {
		params[i] = id
	}
	return model.DeleteBatch(params)
}
