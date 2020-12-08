package icon

import (
	"xframe/backend/common/resp"

	"github.com/gin-gonic/gin"
)

func Fontawesome(c *gin.Context) {
	resp.BuildTpl(c, "demo/icon/fontawesome").Write()
}

func Glyphicons(c *gin.Context) {
	resp.BuildTpl(c, "demo/icon/glyphicons").Write()
}
