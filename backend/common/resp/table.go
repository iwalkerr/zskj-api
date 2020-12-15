package resp

import (
	"net/http"
	"xframe/backend/common/constant"

	"github.com/gin-gonic/gin"
)

// 通用api响应
type TableResp struct {
	t *constant.TableDataInfo
	c *gin.Context
}

//返回一个成功的消息体
func BuildTable(c *gin.Context, pr *constant.PageReq, rows interface{}) *TableResp {
	msg := constant.TableDataInfo{
		Code:  constant.SUCCESS,
		Msg:   "操作成功",
		Total: *pr.Total,
		Rows:  rows,
	}
	return &TableResp{
		t: &msg,
		c: c,
	}
}

//输出json到客户端
func (resp *TableResp) Write() {
	resp.c.JSON(http.StatusOK, resp.t)
	resp.c.Abort()
}
