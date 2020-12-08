package form

import (
	"xframe/backend/common/resp"

	"github.com/gin-gonic/gin"
)

func Autocomplete(c *gin.Context) {

	resp.BuildTpl(c, "demo/form/autocomplete").Write()
}

func Basic(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/basic").Write()
}

func Button(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/button").Write()
}

func Cards(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/cards").Write()
}

func Datetime(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/datetime").Write()
}

func Duallistbox(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/duallistbox").Write()
}

func Grid(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/grid").Write()
}

func Jasny(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/jasny").Write()
}

func Select(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/select").Write()
}

func Sortable(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/sortable").Write()
}

func Summernote(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/summernote").Write()
}

func Tabs_panels(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/tabs_panels").Write()
}

func Timeline(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/timeline").Write()
}

func Upload(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/upload").Write()
}

func Validate(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/validate").Write()
}

func Wizard(c *gin.Context) {
	resp.BuildTpl(c, "demo/form/wizard").Write()
}
