/*
验证：
1、JSON数据解析和绑定
curl -i -X POST \
   -H "Content-Type:application/json" \
   -d \
'{"user":"root","password":"admin"}' \
 'http://127.0.0.1:8080/loginJSON'

2、表单数据解析和绑定
curl -i -X POST \
   -H "Content-Type:application/x-www-form-urlencoded" \
   -d "username=root" \
   -d "password=admin" \
 'http://127.0.0.1:8080/loginForm'

3、URI数据解析和绑定
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
	// 带 binding:"required" 表示该字段必须传递，否则会报错
	// 注：后面的form:user表示在form中这个字段是user，不是User（即：struct中大写，form中小写）
	User     string `form:"username" json:"user" uri:"user" binding:"required"`
	Password string `form:"password" json:"password" uri:"password" binding:"required"`
	// User     string `json:"user" binding:"required"`
	// Password string `json:"password" binding:"required"`
}

// json数据解析和绑定
func loginJSON(c *gin.Context) {
	// 声明接收的变量
	var json Login
	// 将request的body中的数据，自动按照JSON格式解析到login中
	if err := c.ShouldBindJSON(&json); err != nil {
		// 返回错误信息
		// gin.H 封装了生成json数据的工具
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	// 判断用户名密码是否正确
	if json.User != "root" || json.Password != "admin" {
		c.JSON(http.StatusBadRequest, gin.H{"status": "304"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"status": "200"})

}

// 表单数据解析和绑定
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

// uri数据解析和绑定
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
	r.POST("/loginJSON", loginJSON)
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
