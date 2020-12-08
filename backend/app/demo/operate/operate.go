package operate

import (
	"net/http"
	"xframe/backend/common/constant"
	"xframe/backend/common/resp"

	"github.com/gin-gonic/gin"
)

func Add(c *gin.Context) {
	resp.BuildTpl(c, "demo/operate/add").Write()
}

func Detail(c *gin.Context) {
	var tmp us
	tmp.UserId = 1
	tmp.UserName = "测试1"
	tmp.Status = "0"
	tmp.CreateTime = "2020-01-12 02:02:02"
	tmp.UserBalance = 100
	tmp.UserCode = "1000001"
	tmp.UserSex = "0"
	tmp.UserPhone = "15888888888"
	tmp.UserEmail = "111@qq.com"
	resp.BuildTpl(c, "demo/operate/detail").Write(gin.H{"user": tmp})
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

func EditSave(c *gin.Context) {
	var tmp us
	tmp.UserId = 1
	tmp.UserName = "测试1"
	tmp.Status = "0"
	tmp.CreateTime = "2020-01-12 02:02:02"
	tmp.UserBalance = 100
	tmp.UserCode = "1000001"
	tmp.UserSex = "0"
	tmp.UserPhone = "15888888888"
	tmp.UserEmail = "111@qq.com"
	resp.Success(c).Data(tmp).Log("demo演示", gin.H{"UserId": 1}).Write()
}

func Edit(c *gin.Context) {
	var tmp us
	tmp.UserId = 1
	tmp.UserName = "测试1"
	tmp.Status = "0"
	tmp.CreateTime = "2020-01-12 02:02:02"
	tmp.UserBalance = 100
	tmp.UserCode = "1000001"
	tmp.UserSex = "0"
	tmp.UserPhone = "15888888888"
	tmp.UserEmail = "111@qq.com"
	resp.BuildTpl(c, "demo/operate/edit").Write(gin.H{"user": tmp})
}

func Other(c *gin.Context) {
	resp.BuildTpl(c, "demo/operate/other").Write()
}

func Table(c *gin.Context) {
	resp.BuildTpl(c, "demo/operate/table").Write()
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
