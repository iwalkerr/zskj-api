package report

import (
	"xframe/backend/common/resp"

	"github.com/gin-gonic/gin"
)

func Echarts(c *gin.Context) {
	resp.BuildTpl(c, "demo/report/echarts").Write()
}

func Metrics(c *gin.Context) {
	resp.BuildTpl(c, "demo/report/metrics").Write()
}

func Peity(c *gin.Context) {
	resp.BuildTpl(c, "demo/report/peity").Write()
}

func Sparkline(c *gin.Context) {
	resp.BuildTpl(c, "demo/report/sparkline").Write()
}
