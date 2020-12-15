package resp

import (
	"net/http"
	"xframe/backend/common/constant"
	"xframe/backend/core/middleware/sessions"

	"github.com/gin-gonic/gin"
)

// 通用tpl响应
type TplResp struct {
	c   *gin.Context
	tpl string
}

//返回一个tpl响应
func BuildTpl(c *gin.Context, tpl string) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: tpl,
	}
	return &t
}

//返回一个错误的tpl响应
func ErrorTpl(c *gin.Context) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: constant.ERROR_PAGE,
	}
	return &t
}

//返回一个无操作权限tpl响应
func ForbiddenTpl(c *gin.Context) *TplResp {
	var t = TplResp{
		c:   c,
		tpl: constant.UNAUTH_PAGE,
	}
	return &t
}

//输出页面模板
func (resp *TplResp) Write(params ...gin.H) {
	session := sessions.Default(resp.c)
	uid := session.Get(constant.USER_ID)
	if len(params) == 0 {
		resp.c.HTML(http.StatusOK, resp.tpl, gin.H{"uid": uid})
	} else {
		params[0]["uid"] = uid
		resp.c.HTML(http.StatusOK, resp.tpl, params[0])
	}
	resp.c.Abort()
}
