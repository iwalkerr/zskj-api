package swagger

import (
	"fmt"
	_ "xframe/frontend/core"
	"xframe/pkg/router"

	ginSwagger "github.com/swaggo/gin-swagger"
	"github.com/swaggo/gin-swagger/swaggerFiles"
)

// API文档
func init() {
	SwaggerInfo.Title = "Golang API 测试"
	SwaggerInfo.Version = "2.0"
	SwaggerInfo.Description = fmt.Sprintf(`生成文档请在调试模式下进行 <a href=\"http://127.0.0.1%s/tool/swagger?a=r\">重新生成文档</a>`, ":8086")
	SwaggerInfo.Host = fmt.Sprintf("127.0.0.1%s", ":8086")
	SwaggerInfo.BasePath = "/"
	SwaggerInfo.Schemes = []string{"http", "https"}

	g1 := router.New("api", "/swagger")
	g1.GET("/*any", "", ginSwagger.WrapHandler(swaggerFiles.Handler, ginSwagger.URL("http://127.0.0.1:8086/swagger/doc.json")))
}
