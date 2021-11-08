// 获取表单参数
/* 验证：
curl -i -X POST \
   -H "Content-Type:application/x-www-form-urlencoded" \
   -d "username=alan" \
   -d "password=123456" \
 'http://127.0.0.1:8080/form_post'
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func postForm(c *gin.Context) {
	types := c.DefaultPostForm("type", "post")
	username := c.PostForm("username")
	password := c.PostForm("password")
	c.String(http.StatusOK, "type: %s, username: %s, password: %s", types, username, password)
}

func main() {
	r := gin.Default()
	r.POST("/form_post", postForm)
	err := r.Run(":8080")
	// 使用第三方错误处理包：github.com/pkg/errors 进行错误处理
	wrappedErr := errors.Wrap(err, "Gin服务器启动失败")
	fmt.Println(wrappedErr)
}
