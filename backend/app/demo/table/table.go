package table

import (
	"net/http"
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"

	"github.com/gin-gonic/gin"
)

func Button(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/button").Write()
}

func Child(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/child").Write()
}

func Curd(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/curd").Write()
}

func Detail(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/detail").Write()
}

func Editable(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/editable").Write()
}

func Event(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/event").Write()
}

func Export(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/export").Write()
}

func FixedColumns(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/fixedColumns").Write()
}

func Footer(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/footer").Write()
}

func GroupHeader(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/groupHeader").Write()
}

func Image(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/image").Write()
}

func Multi(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/multi").Write()
}

func Other(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/other").Write()
}

func PageGo(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/pageGo").Write()
}

func Params(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/params").Write()
}

func Remember(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/remember").Write()
}

func Recorder(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/recorder").Write()
}

func Search(c *gin.Context) {
	resp.BuildTpl(c, "demo/table/search").Write()
}

type us struct {
	UserId      int64   `json:"userId"`
	UserCode    string  `json:"userCode"`
	UserName    string  `json:"userName"`
	UserSex     string  `json:"userSex"`
	UserPhone   string  `json:"userPhone"`
	UserEmail   string  `json:"userEmail"`
	UserBalance float64 `json:"userBalance"`
	Status      string  `json:"status"`
	CreateTime  string  `json:"createTime"`
}

func List(c *gin.Context) {
	var rows = make([]us, 0)
	for i := 1; i <= 10; i++ {
		var tmp us
		tmp.UserId = int64(i)
		tmp.UserName = "测试" + string(i)
		tmp.Status = "0"
		tmp.CreateTime = "2020-01-12 02:02:02"
		tmp.UserBalance = 100
		tmp.UserCode = "100000" + string(i)
		tmp.UserSex = "0"
		tmp.UserPhone = "15888888888"
		tmp.UserEmail = "111@qq.com"
		rows = append(rows, tmp)
	}
	c.JSON(http.StatusOK, constant.TableDataInfo{
		Code:  0,
		Msg:   "操作成功",
		Total: len(rows),
		Rows:  rows,
	})
}
