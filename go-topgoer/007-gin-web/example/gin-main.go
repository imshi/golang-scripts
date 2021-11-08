package main

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

var db = make(map[string]string)

// var db = map[string]string{"foo": "bar"}

func setupRouter() *gin.Engine {
	// 【可选】关掉控制台颜色：将日志输出到文件时如果不关闭颜色会有\033[31m这种的控制台颜色转义输出；
	// gin.DisableConsoleColor()

	// 初始化 HttpHandler（HTTP处理程序）
	r := gin.Default()

	// Ping 测试，注册/ping 路由，绑定匿名处理函数
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	// 注册一个动态路由获取user（map类型变量）的值，可以匹配 /user/joy，不能匹配 /user 和 /user/
	r.GET("/user/:name", func(c *gin.Context) {
		// 使用 c.Param(key) 获取 url 参数
		user := c.Params.ByName("name")
		value, ok := db[user]
		if ok {
			// user存在则生成对应json并响应（gin.H用来简化生成 json 的方式）
			c.JSON(http.StatusOK, gin.H{"user": user, "value": value})
		} else {
			c.JSON(http.StatusOK, gin.H{"user": user, "status": "no value"})
		}
	})

	//  r.Group()：分组路由，提供api的分组管理功能
	// 使用 gin.BasicAuth()中间件进行 Group 认证
	/*
		等效方式一:
		authorized := r.Group("/")
		authorized.Use(gin.BasicAuth(gin.Credentials{
			  "foo":  "bar",
			  "manu": "123",
		}))
	*/
	// 等效方式二:
	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar", // user:foo password:bar
		"manu": "123", // user:manu password:123
	}))

	/*
			使用带 basicauth 格式认证信息的 curl 命令模拟请求，Zm9vOmJhcg== 是 "foo:bar" 经过 base64 编码后的值

			curl -X POST \
		  	http://localhost:8080/admin \
		  	-H 'authorization: Basic Zm9vOmJhcg==' \
		  	-H 'content-type: application/json' \
		  	-d '{"value":"bar"}'
	*/
	// 在路由分组中注册路由，并实现其处理函数，.(string)为类型断言，将 c.MustGet(key) 转换为string类型
	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		// JSON解析，声明 struct 类型的变量并赋值（``括起来用于多条记录赋值）
		var json struct {
			Value string `json:"value" binding:"required"`
		}

		// 获取请求参数并响应
		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	// 服务端初始化
	r := setupRouter()
	// 拉起监听 0.0.0.0:8080 的服务端
	err := r.Run(":8080")

	// 使用第三方错误处理包：github.com/pkg/errors 进行错误处理
	wrappedErr := errors.Wrap(err, "Gin服务器启动失败")
	fmt.Println(wrappedErr)
	// fmt.Printf("%v\n", wrappedErr)
}
