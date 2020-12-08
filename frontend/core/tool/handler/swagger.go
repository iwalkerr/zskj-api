package hander

import (
	"log"
	"net/http"
	"os"
	"xframe/pkg/utils/exec"

	"github.com/gin-gonic/gin"
)

//swagger文档
func Swagger(c *gin.Context) {
	if a := c.Query("a"); a == "r" {
		curDir, err := os.Getwd()
		if err == nil {
			curDir += "/"
		}
		genPath := curDir + "public/swagger"

		if _, err := exec.ExecCommand("swag init -o " + genPath); err != nil { //重新生成文档
			log.Printf("运行命令行 ===> 错误提示: %s", err.Error())
		} else {
			log.Println("文档重新编译成功")
		}
	}
	c.Redirect(http.StatusFound, "http://127.0.0.1:8086/swagger/index.html")
}
