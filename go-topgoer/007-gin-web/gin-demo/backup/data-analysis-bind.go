/*
验证：
1、表单参数
curl -i -X POST \
   -H "Content-Type:application/x-www-form-urlencoded" \
   -d "username=root" \
   -d "password=admin" \
 'http://127.0.0.1:8080/loginForm'

 2、URI参数
curl -i -X GET 'http://127.0.0.1:8080/root/admin'
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// 定义结构体接收数据
type Login struct {
	User     string `form:"username" json:"user" uri:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" binding:"required"`
}

func loginForm(c *gin.Context) {
	var form Login
	// Bind()默认解析并绑定form格式
	// 根据请求头中Content-Type的值，自动选择解析方式
	if err := c.Bind(&form); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if form.User != "root" || form.Password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

func loginUri(c *gin.Context) {
	var uri Login
	// Bind()默认解析并绑定form格式
	// 根据请求头中Content-Type的值，自动选择解析方式
	if err := c.ShouldBindUri(&uri); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if uri.User != "root" || uri.Password != "admin" {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "unauthorized"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	// 绑定JSON数据
	r.POST("/loginForm", loginForm)
	r.GET("/:user/:password", loginUri)
	return r

}

func main() {
	// 创建路由
	r := setupRouter()

	// 绑定端口，启动服务，使用三方包进行错误处理
	err := r.Run(":8080")
	wrappedErr := errors.Wrap(err, "Gin服务器启动失败")
	fmt.Println(wrappedErr)
}
