// 全局中间件：所有请求都经过此中间件
// 局部、分组中间件，指定的路由经过此中间件
/*
验证：
curl -i -X GET 'http://127.0.0.1:8080/ce'
*/
package main

import (
	"fmt"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func MiddleWare() gin.HandlerFunc {
	return func(c *gin.Context) {
		t := time.Now()
		fmt.Println("全局中间件开始了")
		// 设置变量到Context的Key中，可以通过Get()获取
		c.Set("request", "全局中间件")
		// 调用后续的处理函数、中间件等
		c.Next()
		// 中间件执行完后续的一些事情
		status := c.Writer.Status()
		fmt.Println("中间件执行完毕", status)
		t2 := time.Since(t)
		fmt.Println("全局中间件执行时间：", t2)
	}
}

func main() {
	r := gin.Default()
	// 注册局部中间件
	// r.GET("/ce", MiddleWare(), func(c *gin.Context) {
	// 注册全局中间件
	r.Use(MiddleWare())
	r.GET("/ce", func(c *gin.Context) {
		// 获取变量
		req := c.MustGet("request")
		fmt.Println("request:", req)
		// 页面接收
		c.JSON(200, gin.H{
			"request": req,
		})
	})
	err := r.Run(":8080")
	wrappedErr := errors.Wrap(err, "Gin服务器启动失败")
	fmt.Println(wrappedErr)
}
