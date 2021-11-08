// 获取api参数
// 验证：curl -i -X GET 'http://127.0.0.1:8080/user/alan/best'
package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func getUser(c *gin.Context) {
	name := c.Param("name")
	action := c.Param("action")
	// 截取
	action = strings.Trim(action, "/")
	c.String(http.StatusOK, name+" is "+action)
}

func main() {
	r := gin.Default()
	r.GET("/user/:name/*action", getUser)
	err := r.Run(":8080")
	// 使用第三方错误处理包：github.com/pkg/errors 进行错误处理
	wrappedErr := errors.Wrap(err, "Gin服务器启动失败")
	fmt.Println(wrappedErr)
}
