package modal

import (
	"xframe/backend/common/resp"

	"github.com/gin-gonic/gin"
)

func Dialog(c *gin.Context) {
	resp.BuildTpl(c, "demo/modal/dialog").Write()
}

func Form(c *gin.Context) {
	resp.BuildTpl(c, "demo/modal/form").Write()
}

func Layer(c *gin.Context) {
	resp.BuildTpl(c, "demo/modal/layer").Write()
}

func Table(c *gin.Context) {
	resp.BuildTpl(c, "demo/modal/table").Write()
}

func Check(c *gin.Context) {
	resp.BuildTpl(c, "demo/modal/table/check").Write()
}

func Parent(c *gin.Context) {
	resp.BuildTpl(c, "demo/modal/table/parent").Write()
}

func Radio(c *gin.Context) {
	resp.BuildTpl(c, "demo/modal/table/radio").Write()
}
