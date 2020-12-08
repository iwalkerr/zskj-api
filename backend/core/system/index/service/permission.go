package service

import (
	"html/template"
	menuModel "xframe/backend/core/system/menu/model"
	menuService "xframe/backend/core/system/menu/service"
)

// 递归检查权限
func CheckPermi(menus []menuModel.Entity, permission string, flag *bool) {
	for i := 0; i < len(menus); i++ {
		if menus[i].Perms == permission {
			*(flag) = true
			return
		}
		if menus[i].SubMenu != nil && len(menus[i].SubMenu) > 0 {
			CheckPermi(menus[i].SubMenu, permission, flag)
		}
	}
}

//根据用户id和权限字符串判断是否有此权限
func HasPermi(u interface{}, permission string) string {
	if u == nil {
		return "disabled"
	}

	uid, ok := u.(int)
	if uid <= 0 || !ok {
		return "disabled"
	}

	//获取权限信息
	menus := menuService.SelectMenuNormalByUser(uid)

	var hasPermi bool
	CheckPermi(*menus, permission, &hasPermi)
	if hasPermi {
		return ""
	}

	return "disabled"
}

//根据用户id和权限字符串判断是否输出控制按钮
func GetPermiButton(u interface{}, permission, funcName, text, aclassName, iclassName string) template.HTML {
	htmlstr := ""

	if result := HasPermi(u, permission); result == "" {
		htmlstr = `
			<a class="` + aclassName + `" onclick="` + funcName + `" hasPermission="` + permission + `">
				<i class="` + iclassName + `"></i> ` + text + `
			</a>
		`
	}

	return template.HTML(htmlstr)
}
