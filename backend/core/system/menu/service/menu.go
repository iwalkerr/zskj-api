package service

import (
	"errors"
	"strconv"
	"strings"
	"time"
	"xframe/backend/common/constant"
	"xframe/backend/core/middleware/sessions"
	"xframe/backend/core/system/menu/model"
	userService "xframe/backend/core/system/user/service"
	"xframe/pkg/cache"

	"github.com/gin-gonic/gin"
	"github.com/mohae/deepcopy"
)

func RoleMenuTreeData(roleId, userId int) *[]constant.Ztree {
	menuList := SelectMenuNormalByUser(userId)
	if len(*menuList) == 0 {
		return nil
	}

	// 树转成数组
	var resList []model.Entity
	AllMenuArr(*menuList, &resList)

	var ztree *[]constant.Ztree
	if roleId > 0 {
		roleMenuList := model.SelectMenuTree(roleId)
		if len(roleMenuList) == 0 {
			ztree = InitZtree(&resList, nil, true)
		} else {
			ztree = InitZtree(&resList, &roleMenuList, true)
		}
	} else {
		ztree = InitZtree(&resList, nil, true)
	}

	return ztree
}

//根据主键删除数据
func DeleteRecordByIds(ids string) bool {
	idsArr := strings.Split(ids, ",")

	var ismore bool
	for _, id := range idsArr {
		count := model.ChildMenuCount(id)
		if count > 0 {
			ismore = true
		} else {
			_ = model.DeleteRecordById(id)
		}
	}

	return ismore
}

//检查菜单名是否唯一
func CheckMenuNameUnique(menuName string, menuId, parentId int) string {
	mId := model.CheckMenuNameUniqueAll(menuName, parentId)
	if mId > 0 && mId != menuId {
		return "1"
	}
	return "0"
}

func AddSave(c *gin.Context, req *model.AddReq) (int, error) {
	var entity model.Entity
	entity.Name = req.MenuName
	if req.Visible == "" {
		entity.Visible = "0"
	} else {
		entity.Visible = req.Visible
	}
	entity.ParentId = req.ParentId
	entity.MenuType = req.MenuType
	if req.Url == "" {
		entity.Url = "#"
	} else {
		entity.Url = req.Url
	}
	entity.Perms = req.Perms
	entity.SysType = req.SysType
	if req.Icon == "" {
		entity.Icon = "#"
	} else {
		entity.Icon = req.Icon + " " + req.Color
	}
	entity.SortId = req.OrderNum
	entity.CreateTime = time.Now()

	user := userService.GetProfile(c)
	if user != nil {
		entity.CreateBy = user.LoginName
	}

	return entity.AddSave()
}

//检查菜单名是否唯一
func CheckMenuNameUniqueAll(menuName string, parentId int) string {
	rs := model.CheckMenuNameUniqueAll(menuName, parentId)
	if rs > 0 {
		return "1"
	}
	return "0"
}

func EditSave(c *gin.Context, req *model.EditReq) error {
	entity := model.SelectRecordById(req.MenuId)
	if entity.MenuId <= 0 {
		return errors.New("菜单不存在")
	}
	entity.Name = req.MenuName
	entity.Visible = req.Visible
	entity.ParentId = req.ParentId
	entity.MenuType = req.MenuType
	entity.SysType = req.SysType
	entity.Url = req.Url
	entity.Icon = req.Icon + " " + req.Color
	entity.Perms = req.Perms
	entity.SortId = req.OrderNum

	if u := userService.GetProfile(c); u != nil {
		entity.UpdateBy = u.LoginName
	}

	if err := entity.EditSave(); err != nil {
		return err
	}

	return nil
}

// 根据ID查询记录
func SelectRecordById(id int) model.Entity {
	menu := model.SelectRecordById(id)
	split := strings.Split(menu.Icon, " ")

	menu.Icon = strings.Join(split[0:len(split)-1], " ")
	menu.Color = split[len(split)-1]

	return menu
}

func SelectDictTree(parentId int) *[]constant.Ztree {
	list := model.SelectMenuList(parentId)

	return InitZtree(&list, nil, false)
}

//对象转部门树
func InitZtree(menuList *[]model.Entity, roleMenuList *[]string, permsFlag bool) *[]constant.Ztree {
	var ztreeList []constant.Ztree
	isCheck := false
	if roleMenuList != nil && len(*roleMenuList) > 0 {
		isCheck = true
	}

	for _, menu := range *menuList {
		var ztree constant.Ztree
		ztree.Id = menu.MenuId
		ztree.Pid = menu.ParentId
		ztree.Name = transMenuName(menu.Name, permsFlag)
		ztree.Title = menu.Name
		if isCheck {
			tmp := strconv.Itoa(menu.MenuId) + menu.Perms
			tmpcheck := false
			for j := range *roleMenuList {
				if (*roleMenuList)[j] == tmp {
					tmpcheck = true
					break
				}
			}
			ztree.Checked = tmpcheck
		}
		ztreeList = append(ztreeList, ztree)
	}
	return &ztreeList
}

func transMenuName(menuName string, permsFlag bool) string {
	if permsFlag {
		return "<font color=\"#888\">&nbsp;&nbsp;&nbsp;" + menuName + "</font>"
	} else {
		return menuName
	}
}

func SelectRecordList(param *model.SelectPageReq) []model.Entity {
	return model.SelectPageList(param)
}

// 获取用户的菜单数据
func SelectMenuNormalByUser(userId int) *[]model.Entity {
	if userService.IsAdmin(userId) {
		return GetAdminCacheMenu(userId)
	} else {
		return GetUserCacheMenu(userId)
	}
}

const change = "change_index"

// 切换菜单
func ChangeMenu(c *gin.Context, allMenu *[]model.Entity, userId string, param string) *[]model.Entity {
	key := "current_" + userId
	menuList := currentMenu(key)
	if menuList != nil && param == "main" {
		return menuList
	}

	menuList1 := make([]model.Entity, 0)
	menuList2 := make([]model.Entity, 0)

	for _, menu := range *allMenu {
		if menu.SysType == "1" {
			menuList1 = append(menuList1, menu)
		} else {
			menuList2 = append(menuList2, menu)
		}
	}

	// 切换系统和业务菜单树
	session := sessions.Default(c)
	changeIndex := session.Get(change)
	if changeIndex == nil || changeIndex == "2" {
		menuCache(session, menuList2, key, "1")
		menuList = &menuList2
	} else {
		menuCache(session, menuList1, key, "2")
		menuList = &menuList1
	}

	return menuList
}

// cache中缓存菜单
func menuCache(session sessions.Session, menuList []model.Entity, key, param string) {
	cache.Instance().Set(constant.MENU_CACHE+key, menuList, time.Hour)
	session.Set(change, param)
	_ = session.Save()
}

// 获取当前菜单
func currentMenu(key string) *[]model.Entity {
	tmp, ok := cache.Instance().Get(constant.MENU_CACHE + key)
	if tmp != nil && ok {
		if rs, ok := tmp.([]model.Entity); ok {
			return &rs
		}
	}
	return nil
}

// 获取缓存菜单
func GetUserCacheMenu(userId int) *[]model.Entity {
	userIds := strconv.Itoa(userId)
	if menuList := currentMenu(userIds); menuList != nil {
		return menuList
	}
	// 从数据库中读取
	m := model.Entity{}
	m.MenuId = 0
	list := ListAllPermMenu(&m, userId)

	// 存入缓存
	cache.Instance().Set(constant.MENU_CACHE+userIds, list, time.Hour)
	return &list
}

//获取管理员菜单数据
func GetAdminCacheMenu(userId int) *[]model.Entity {
	userIds := strconv.Itoa(userId)
	if menuList := currentMenu(userIds); menuList != nil {
		return menuList
	}
	m := model.Entity{}
	m.MenuId = 0
	list := ListAllMenu(&m)

	// 存入缓存
	cache.Instance().Set(constant.MENU_CACHE+userIds, list, time.Hour)
	return &list
}

// 加载所有菜单列表树
func MenuTreeData(userId int) *[]constant.Ztree {
	menuList := SelectMenuNormalByUser(userId)
	if len(*menuList) <= 0 {
		return nil
	}
	menuList = RemoveAllMenuF(menuList)

	// 树转成数组
	var resList []model.Entity
	AllMenuArr(*menuList, &resList)

	return InitZtree(&resList, nil, false)
}

// 所有tree菜单转成数组
func AllMenuArr(menus []model.Entity, res *[]model.Entity) {
	for i := 0; i < len(menus); i++ {
		*res = append(*res, menus[i])
		AllMenuArr(menus[i].SubMenu, res)
	}
}

// 删除F
func RemoveAllMenuF(menuList *[]model.Entity) *[]model.Entity {
	// 去除F菜单
	menuListF := deepcopy.Copy(*menuList).([]model.Entity) // 此处用到深拷贝
	RemoveMenuF(&menuListF, false)
	return &menuListF
}

// 移除所有tree F菜单
func RemoveMenuF(menus *[]model.Entity, flag bool) bool {
	for i := 0; i < len(*menus); i++ {
		if (*menus)[i].MenuType == "F" {
			flag = true
			break
		}
		if len(*menus) > 0 && len((*menus)[i].SubMenu) > 0 {
			sm := (*menus)[i].SubMenu
			if RemoveMenuF(&sm, flag) {
				(*menus)[i].SubMenu = nil
			}
		}
	}
	return flag
}

// 递归获取所有菜单
func ListAllMenu(m *model.Entity) []model.Entity {
	list := model.ListMenuByParentId(m.MenuId)
	for i := 0; i < len(list); i++ {
		list[i].SubMenu = ListAllMenu(&list[i])
	}
	return list
}

// 获取所有权限菜单
func ListAllPermMenu(m *model.Entity, userId int) []model.Entity {
	list := model.ListPermMenuByParentId(m.MenuId, userId)
	for i := 0; i < len(list); i++ {
		list[i].SubMenu = ListAllPermMenu(&list[i], userId)
	}
	return list
}
