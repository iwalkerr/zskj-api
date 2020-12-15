package service

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"xframe/backend/common/constant"
	"xframe/backend/core/system/dept/model"
	userService "xframe/backend/core/system/user/service"

	"github.com/gin-gonic/gin"
)

// 根据ID获取部门名称
func GetDeptName(deptId int) string {
	return model.DeptNameById(deptId)
}

func DeleteRecordByIds(ids string) int {
	idArr := strings.Split(ids, ",")
	if err := model.DeleteBatch(idArr); err != nil {
		return 0
	}
	return 1
}

//新增保存信息
func AddSave(c *gin.Context, req *model.AddReq) (int, error) {
	var entity model.Entity

	parent := model.SelectDepById(req.ParentId)
	if parent.DeptId == 0 {
		return 0, errors.New("父部门不能为空")
	}
	if parent.Status != "0" {
		return 0, errors.New("部门停用，不允许新增")
	} else {
		entity.Ancestors = parent.Ancestors + "," + strconv.Itoa(parent.DeptId)
	}

	entity.DeptName = req.DeptName
	entity.Status = req.Status
	entity.ParentId = req.ParentId
	entity.DelFlag = "0"
	entity.Email = req.Email
	entity.Leader = req.Leader
	entity.Phone = req.Phone
	entity.OrderNum = req.OrderNum
	entity.CreateTime = time.Now()

	user := userService.GetProfile(c)
	if user != nil && user.UserId > 0 {
		entity.CreateBy = user.LoginName
	}

	return entity.Insert()
}

//校验部门名称是否唯一
func CheckDeptNameUniqueAll(depName string, parentId int) string {
	count := model.CheckDeptNameUniqueAll(depName, parentId)
	if count > 0 {
		return "1"
	}
	return "0"
}

//修改保存信息
func EditSave(c *gin.Context, req *model.EditReq) error {
	d := model.SelectDepById(req.DeptId)
	if d.DeptId <= 0 {
		return errors.New("数据不存在")
	}

	parentId := 0
	if req.ParentId == 0 {
		parentId = 100 // 固定顶部部门
	} else {
		parentId = req.ParentId
	}
	p := model.SelectDepById(parentId)
	if p.DeptId <= 0 {
		return errors.New("父部门不能为空")
	}
	if p.Status != "0" {
		return errors.New("部门停用，不允许新增")
	}

	newAncestorys := p.Ancestors + "," + strconv.Itoa(p.DeptId)
	model.UpdateDeptChildren(d.DeptId, newAncestorys, d.Ancestors)

	d.DeptName = req.DeptName
	d.Status = req.Status
	d.ParentId = req.ParentId
	d.DelFlag = "0"
	d.Email = req.Email
	d.Leader = req.Leader
	d.Phone = req.Phone
	d.OrderNum = req.OrderNum
	d.UpdateTime = time.Now()

	user := userService.GetProfile(c)
	if user != nil && user.UserId > 0 {
		d.UpdateBy = user.LoginName
	}
	return d.Update()
}

//校验部门名称是否唯一
func CheckDeptNameUnique(deptName string, deptId, parentId int) string {
	id := model.CheckDeptNameUniqueAll(deptName, parentId)
	if id > 0 && id != deptId {
		return "1"
	}
	return "0"
}

//根据部门ID查询信息
func SelectDeptById(pid int) model.Entity {
	return model.SelectDepById(pid)
}

//根据角色ID查询部门（数据权限）
func RoleDeptTreeData(roleId int) (*[]constant.Ztree, error) {
	var result *[]constant.Ztree

	deptList, err := model.SelectDeptList(0, "", "")
	if err != nil {
		return nil, err
	}

	if roleId > 0 {
		roleDeptList := model.SelectRoleDeptTree(roleId)
		if len(roleDeptList) > 0 {
			result = InitZtree(deptList, &roleDeptList)
		} else {
			result = InitZtree(deptList, nil)
		}
	} else {
		result = InitZtree(deptList, nil)
	}

	return result, nil
}

func SelectRecordList(req *model.SelectPageReq) []model.Entity {
	return model.SelectPageList(req)
}

// 查询部门数
func SelectDeptTree(parentId int, deptName, status string) *[]constant.Ztree {
	list := model.DeptList(parentId, deptName, status)

	return InitZtree(list, nil)
}

//对象转部门树
func InitZtree(deptList []model.Entity, roleDeptList *[]string) *[]constant.Ztree {
	var ztreeList []constant.Ztree

	isCheck := false
	if roleDeptList != nil && len(*roleDeptList) > 0 {
		isCheck = true
	}

	for _, dept := range deptList {
		if dept.Status == "0" {
			var ztree constant.Ztree
			ztree.Id = dept.DeptId
			ztree.Pid = dept.ParentId
			ztree.Name = dept.DeptName
			ztree.Title = dept.DeptName
			if isCheck {
				tmp := strconv.Itoa(dept.DeptId) + dept.DeptName
				tmpcheck := false
				for _, roleDept := range *roleDeptList {
					if strings.EqualFold(roleDept, tmp) {
						tmpcheck = true
						break
					}
				}
				ztree.Checked = tmpcheck
			}
			ztreeList = append(ztreeList, ztree)
		}
	}
	return &ztreeList
}
