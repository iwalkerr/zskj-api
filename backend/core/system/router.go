package system

import (
	"xframe/backend/core/middleware/auth"
	config "xframe/backend/core/system/config/handler"
	dept "xframe/backend/core/system/dept/handler"
	dict "xframe/backend/core/system/dict/handler"
	menu "xframe/backend/core/system/menu/handler"
	post "xframe/backend/core/system/post/handler"
	role "xframe/backend/core/system/role/handler"
	user "xframe/backend/core/system/user/handler"
	"xframe/pkg/router"
)

func init() {
	// 用户路由
	g1 := router.New("admin", "/system/user", auth.Auth)
	g1.GET("/", "system:user:view", user.List)
	g1.POST("/list", "system:user:list", user.ListAjax)
	g1.GET("/add", "system:user:add", user.Add)
	g1.POST("/add", "system:user:add", user.AddSave)
	g1.GET("/edit", "system:user:edit", user.Edit)
	g1.POST("/edit", "system:user:edit", user.EditSave)
	g1.POST("/remove", "system:user:remove", user.Remove)
	g1.POST("/export", "system:user:export", user.Export)
	g1.GET("/resetPwd", "system:user:resetPwd", user.ResetPwd)
	g1.POST("/resetPwd", "system:user:resetPwd", user.ResetPwdSave)

	// 部门路由
	g2 := router.New("admin", "/system/dept", auth.Auth)
	g2.GET("/treeData", "system:dept:view", dept.TreeData)
	g2.GET("/", "system:dept:view", dept.List)
	g2.POST("/list", "system:dept:list", dept.ListAjax)
	g2.GET("/add", "system:dept:add", dept.Add)
	g2.POST("/add", "system:dept:add", dept.AddSave)
	g2.GET("/edit", "system:dept:edit", dept.Edit)
	g2.POST("/edit", "system:dept:edit", dept.EditSave)
	g2.POST("/remove", "system:dept:remove", dept.Remove)
	g2.GET("/roleDeptTreeData", "system:dept:view", dept.RoleDeptTreeData)
	g2.GET("/selectDeptTree", "system:dept:view", dept.SelectDeptTree)
	g2.POST("/checkDeptNameUnique", "system:dept:view", dept.CheckDeptNameUnique)
	g2.POST("/checkDeptNameUniqueAll", "system:dept:view", dept.CheckDeptNameUniqueAll)
	g2.GET("/staff", "system:dept:staff", dept.Edit)

	// 数据字典
	g3 := router.New("admin", "/system/dict", auth.Auth)
	g3.GET("/", "system:dict:view", dict.List)
	g3.POST("/list", "system:dict:list", dict.ListAjax)
	g3.GET("/treeData", "system:dict:view", dict.TreeData)
	g3.GET("/add", "system:dict:add", dict.Add)
	g3.POST("/add", "system:dict:add", dict.AddSave)
	g3.GET("/edit", "system:dict:edit", dict.Edit)
	g3.POST("/edit", "system:dict:edit", dict.EditSave)
	g3.POST("/remove", "system:dict:remove", dict.Remove)

	// 菜单管理
	g4 := router.New("admin", "/system/menu", auth.Auth)
	g4.GET("/", "system:menu:view", menu.List)
	g4.POST("/list", "system:menu:list", menu.ListAjax)
	g4.GET("/treeData", "system:menu:view", menu.TreeData)
	g4.GET("/add", "system:menu:add", menu.Add)
	g4.POST("/add", "system:menu:add", menu.AddSave)
	g4.GET("/edit", "system:menu:edit", menu.Edit)
	g4.POST("/edit", "system:menu:edit", menu.EditSave)
	g4.POST("/remove", "system:menu:remove", menu.Remove)
	g4.POST("/checkMenuNameUnique", "system:menu:view", menu.CheckMenuNameUnique)
	g4.POST("/checkMenuNameUniqueAll", "system:menu:view", menu.CheckMenuNameUniqueAll)
	g4.GET("/selectMenuTree", "system:menu:view", menu.SelectMenuTree)
	g4.GET("/menuTreeData", "system:menu:view", menu.MenuTreeData)
	g4.GET("/roleMenuTreeData", "system:menu:view", menu.RoleMenuTreeData)

	// 角色管理
	g5 := router.New("admin", "/system/role", auth.Auth)
	g5.GET("/", "system:role:view", role.List)
	g5.POST("/list", "system:role:list", role.ListAjax)
	g5.GET("/add", "system:role:add", role.Add)
	g5.POST("/add", "system:role:add", role.AddSave)
	g5.GET("/edit", "system:role:edit", role.Edit)
	g5.POST("/edit", "system:role:edit", role.EditSave)
	g5.POST("/remove", "system:role:remove", role.Remove)
	g5.POST("/checkRoleKeyUnique", "system:post:view", role.CheckRoleKeyUnique)
	g5.POST("/checkRoleNameUnique", "system:post:view", role.CheckRoleNameUnique)
	g5.POST("/checkRoleNameUniqueAll", "system:post:view", role.CheckRoleNameUniqueAll)
	g5.POST("/checkRoleKeyUniqueAll", "system:post:view", role.CheckRoleKeyUniqueAll)
	g5.POST("/changeStatus", "system:role:edit", role.ChangeStatus)
	g5.GET("/authDataScope", "system:post:view", role.AuthDataScope)
	g5.POST("/authDataScope", "system:post:view", role.AuthDataScopeSave)
	g5.GET("/authUser", "system:post:view", role.AuthUser)
	g5.POST("/allocatedList", "system:post:view", role.AllocatedList)
	g5.GET("/selectUser", "system:post:view", role.SelectUser)
	g5.POST("/unallocatedList", "system:post:view", role.UnallocatedList)
	g5.POST("/selectAll", "system:post:view", role.SelectAll)
	g5.POST("/cancelAll", "system:post:view", role.CancelAll)
	g5.POST("/cancel", "system:post:view", role.Cancel)

	// 岗位管理
	g6 := router.New("admin", "/system/post", auth.Auth)
	g6.GET("/", "system:post:view", post.List)
	g6.POST("/list", "system:post:list", post.ListAjax)
	g6.GET("/add", "system:post:add", post.Add)
	g6.POST("/add", "system:post:add", post.AddSave)
	g6.GET("/edit", "system:post:edit", post.Edit)
	g6.POST("/edit", "system:post:edit", post.EditSave)
	g6.POST("/remove", "system:post:remove", post.Remove)
	g6.POST("/checkPostNameUnique", "system:post:list", post.CheckPostNameUnique)
	g6.POST("/checkPostCodeUnique", "system:post:list", post.CheckPostCodeUnique)
	g6.POST("/checkPostNameUniqueAll", "system:post:list", post.CheckPostNameUniqueAll)
	g6.POST("/checkPostCodeUniqueAll", "system:post:list", post.CheckPostCodeUniqueAll)

	// 参数路由
	g7 := router.New("admin", "/system/config", auth.Auth)
	g7.GET("/", "system:config:view", config.List)
	g7.POST("/list", "system:config:list", config.ListAjax)
	g7.GET("/add", "system:config:add", config.Add)
	g7.POST("/add", "system:config:add", config.AddSave)
	g7.GET("/edit", "system:config:edit", config.Edit)
	g7.POST("/edit", "system:config:edit", config.EditSave)
	g7.POST("/remove", "system:config:remove", config.Remove)
	g7.POST("/checkConfigKeyUniqueAll", "system:config:view", config.CheckConfigKeyUniqueAll)
	g7.POST("/checkConfigKeyUnique", "system:config:view", config.CheckConfigKeyUnique)

}
