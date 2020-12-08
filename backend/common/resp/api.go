package resp

import (
	"encoding/json"
	"net/http"
	"xframe/backend/common/constant"
	operlog "xframe/backend/core/monitor/operlog/service"

	"github.com/gin-gonic/gin"
)

// 通用api响应
type ApiResp struct {
	r *constant.CommonRes
	c *gin.Context
}

//返回一个成功的消息体
func Success(c *gin.Context) *ApiResp {
	msg := constant.CommonRes{
		Code:  constant.SUCCESS,
		Btype: constant.Buniss_Other,
		Msg:   "操作成功",
	}
	return &ApiResp{
		r: &msg,
		c: c,
	}
}

//返回一个错误的消息体
func Error(c *gin.Context) *ApiResp {
	msg := constant.CommonRes{
		Code:  constant.ERROR,
		Btype: constant.Buniss_Other,
		Msg:   "操作失败",
	}
	return &ApiResp{
		r: &msg,
		c: c,
	}
}

//返回一个拒绝访问的消息体
func Forbidden(c *gin.Context) *ApiResp {
	msg := constant.CommonRes{
		Code:  constant.UNAUTHORIZED,
		Btype: constant.Buniss_Other,
		Msg:   "无操作权限",
	}
	return &ApiResp{
		r: &msg,
		c: c,
	}
}

//设置消息体的内容
func (resp *ApiResp) Msg(msg string) *ApiResp {
	resp.r.Msg = msg
	return resp
}

//设置消息体的编码
func (resp *ApiResp) Code(code int) *ApiResp {
	resp.r.Code = code
	return resp
}

//设置消息体的数据
func (resp *ApiResp) Data(data interface{}) *ApiResp {
	resp.r.Data = data
	return resp
}

//设置消息体的业务类型
func (resp *ApiResp) Btype(btype constant.BunissType) *ApiResp {
	resp.r.Btype = btype
	return resp
}

//记录操作日志到数据库
func (resp *ApiResp) Log(title string, inParam interface{}) *ApiResp {
	var inContent string
	if b, err := json.Marshal(inParam); err == nil {
		inContent = string(b)
	}

	// 增加日志到数据库
	_ = operlog.Add(resp.c, title, inContent, resp.r)

	return resp
}

//输出json到客户端
func (resp *ApiResp) Write() {
	resp.c.JSON(http.StatusOK, resp.r)
	resp.c.Abort()
}
