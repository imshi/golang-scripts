/*
验证：
curl -i -X GET 'http://127.0.0.1:8080/index'
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func main() {
	r := gin.Default()
	r.GET("/index", func(c *gin.Context) {
		c.Redirect(http.StatusMovedPermanently, "http://www.baidu.com")
	})
	// 绑定端口，启动服务，使用三方包进行错误处理
	err := r.Run(":8080")
	wrappedErr := errors.Wrap(err, "Gin服务器启动失败")
	fmt.Println(wrappedErr)
}
