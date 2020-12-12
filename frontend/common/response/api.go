package response

import (
	"encoding/json"
	"net/http"
	"xframe/frontend/common/rsa"

	"github.com/gin-gonic/gin"
)

// 通用api响应
type CommonRes struct {
	Code int         `json:"code"`           // 响应编码 0 成功 500 错误 403 无权限  -1  失败
	Msg  string      `json:"msg"`            // 消息
	Data interface{} `json:"data,omitempty"` // 数据内容
}

// 通用api响应
type ApiResp struct {
	r *CommonRes
	c *gin.Context
}

//返回一个成功的消息体
func Success(c *gin.Context) *ApiResp {
	msg := CommonRes{
		Code: 0,
		Msg:  "操作成功",
	}
	return &ApiResp{
		r: &msg,
		c: c,
	}
}

//返回一个错误的消息体
func Error(c *gin.Context) *ApiResp {
	msg := CommonRes{
		Code: 500,
		Msg:  "操作失败",
	}
	return &ApiResp{
		r: &msg,
		c: c,
	}
}

//返回一个拒绝访问的消息体
func Forbidden(c *gin.Context) *ApiResp {
	msg := CommonRes{
		Code: 403,
		Msg:  "无操作权限",
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
func (resp *ApiResp) Data(params ...gin.H) *ApiResp {
	if len(params) == 0 {
		resp.r.Data = setNt(resp.c, gin.H{})
	} else {
		resp.r.Data = setNt(resp.c, params[0])
	}
	return resp
}

//输出json到客户端
func (resp *ApiResp) JSON() {
	resp.c.JSON(http.StatusOK, resp.r)
	resp.c.Abort()
}

//设置刷新token
func setNt(c *gin.Context, data gin.H) string {
	data["nt"] = c.GetString("nt")
	bytes, err := json.Marshal(data)
	if err != nil {
		return ""
	}
	text, err := rsa.RsaEncrypt(bytes)
	if err != nil {
		return ""
	}
	return text
}
