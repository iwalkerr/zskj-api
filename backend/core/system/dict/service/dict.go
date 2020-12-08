package service

import (
	"encoding/json"
	"errors"
	"html/template"
	"strconv"
	"strings"
	"time"
	"xframe/backend/common/constant"
	"xframe/backend/core/system/dict/model"
	userService "xframe/backend/core/system/user/service"

	"github.com/gin-gonic/gin"
)

//批量删除数据记录
func DeleteRecordByIds(ids string) bool {
	idsArr := strings.Split(ids, ",")

	var ismore bool
	for _, id := range idsArr {
		count := model.ChildDictCount(id)
		if count > 0 {
			ismore = true
		} else {
			_ = model.DeleteRecordById(id)
		}
	}
	return ismore
}

//修改数据
func EditSave(c *gin.Context, req *model.EditReq) (int, error) {
	entity := model.SelectRecordById(req.DictId)
	if entity.DictId <= 0 {
		return 0, errors.New("数据不存在")
	}
	entity.Status = req.Status
	entity.DictLabel = req.DictLabel
	entity.DictSort = req.DictSort
	entity.DictValue = req.DictValue
	entity.IsDefault = req.IsDefault
	entity.ListClass = req.ListClass
	entity.Remark = req.Remark
	entity.UpdateTime = time.Now()

	user := userService.GetProfile(c)
	if user != nil {
		entity.UpdateBy = user.LoginName
	}

	if err := entity.Update(); err != nil {
		return 0, err
	}
	return entity.DictId, nil
}

//根据主键查询数据
func SelectRecordById(id int) model.Entity {
	return model.SelectRecordById(id)
}

//添加数据
func AddSave(c *gin.Context, req *model.AddReq) (int, error) {
	var entity model.Entity
	entity.DictType = req.DictType
	entity.Status = req.Status
	entity.DictLabel = req.DictLabel
	entity.DictSort = req.DictSort
	entity.DictValue = req.DictValue
	entity.IsDefault = req.IsDefault
	entity.ListClass = req.ListClass
	entity.Remark = req.Remark
	entity.ParentId = req.ParentId
	entity.CreateTime = time.Now()
	entity.UpdateTime = time.Now()

	user := userService.GetProfile(c)
	if user != nil {
		entity.CreateBy = user.LoginName
	}

	return entity.Insert()
}

func SelectRecordList(param *model.SelectPageReq) (list []model.Entity) {
	return model.SelectPageList(param)
}

// 查询部门数
func SelectDictTree(parentId int) *[]constant.Ztree {
	list := model.SelectDictList(parentId)

	return InitZtree(list)
}

//对象转部门树
func InitZtree(dictList []model.Entity) *[]constant.Ztree {
	var ztreeList []constant.Ztree
	for _, dict := range dictList {
		var ztree constant.Ztree
		ztree.Id = dict.DictId
		ztree.Pid = dict.ParentId
		ztree.Name = dict.DictLabel
		ztree.Title = dict.DictLabel
		ztreeList = append(ztreeList, ztree)
	}
	return &ztreeList
}

//根据字典类型和字典键值查询字典数据信息
func GetDictLabel(dictType string, dictValue interface{}) template.HTML {
	value, ok := dictValue.(string)
	if !ok {
		return ""
	}

	dictLabel := model.LabelRecord(dictType, value)
	return template.HTML(dictLabel)
}

//通用的字典单选框控件  dictType 字典类别  value 默认值
func GetDictRadio(dictType, name string, value interface{}) template.HTML {
	result := SelectDictByType(dictType)
	if result == nil || len(result) <= 0 {
		return ""
	}

	dictValue, ok := value.(string)
	if !ok {
		return ""
	}

	htmlstr := ``
	for _, item := range result {
		dictId := strconv.Itoa(item.DictId)
		if strings.Compare(item.DictValue, dictValue) == 0 {
			htmlstr += `
			<label class="radio-box">
				<input type="radio" id="` + dictId + `" name="` + name + `" value="` + item.DictValue + `" checked="checked"/> ` + item.DictLabel + `
			</label>`
		} else {
			htmlstr += `
			<label class="radio-box">
				<input type="radio" id="` + dictId + `" name="` + name + `" value="` + item.DictValue + `"/> ` + item.DictLabel + `
			</label>`
		}
	}

	htmlstr += ``
	return template.HTML(htmlstr)
}

//通用的字典下拉框控件  字典类别   html控件id  html控件name html控件class  html控件value  html控件空值标签 是否可以多选
func GetDictSelect(dictType, name, value, emptyLabel, multiple string) template.HTML {
	result := SelectDictByType(dictType)
	if result == nil || len(result) <= 0 {
		return ""
	}

	htmlstr := `<select id="` + name + `" name="` + name + `" class="form-control" ` + multiple + `>`

	if emptyLabel != "" {
		htmlstr += `<option value="">` + emptyLabel + `</option>`
	}

	for _, item := range result {
		if strings.Compare(item.DictValue, value) == 0 {
			htmlstr += `<option selected value="` + item.DictValue + `">` + item.DictLabel + `</option>`
		} else {
			htmlstr += `<option value="` + item.DictValue + `">` + item.DictLabel + `</option>`
		}
	}

	htmlstr += `</select>`

	return template.HTML(htmlstr)
}

// 获取字典列表
func SelectDictByType(dictType string) []model.DictData {
	return model.ListByDictType(dictType)
}

func GetDictData(dictType string) template.JS {
	result := SelectDictByType(dictType)
	bytes, err := json.Marshal(result)
	if err != nil {
		return template.JS("")
	}
	return template.JS(string(bytes))
}
