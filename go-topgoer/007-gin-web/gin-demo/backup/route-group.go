/*
验证：
curl -i -X GET 'http://127.0.0.1:8080/v1/login'
curl -i -X GET 'http://127.0.0.1:8080/v1/submit'
curl -i -X POST -H "Content-Type:application/json" 'http://127.0.0.1:8080/v2/login?name=tom'
curl -i -X POST -H "Content-Type:application/json" 'http://127.0.0.1:8080/v2/submit?name=tom'
*/
package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

func login(c *gin.Context) {
	name := c.DefaultQuery("name", "jack")
	c.String(http.StatusOK, "Hello %s\n", name)
}

func submit(c *gin.Context) {
	name := c.DefaultQuery("name", "lily")
	c.String(http.StatusOK, "Hello %s \n", name)
}

func setupRouter() *gin.Engine {
	// 默认添加两个中间件Logger() 和 Recovery()
	r := gin.Default()
	// 路由1，处理GET请求
	v1 := r.Group("/v1")
	{
		v1.GET("/login", login)
		v1.GET("/submit", submit)
	}
	// 路由2，处理POST请求
	v2 := r.Group("/v2")
	{
		v2.POST("/login", login)
		v2.POST("/submit", submit)
	}
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
