// 使用 goroutine 机制实现异步
// 注：在启动新的 goroutine 时，不能使用原始上下文，必须使用它的只读副本
/*
验证：
1、curl -i -X GET 'http://127.0.0.1:8080/sync'
2、curl -i -X GET 'http://127.0.0.1:8080/async'
*/
package main

import (
	"log"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/pkg/errors"
)

// 异步函数
func longAsync(c *gin.Context) {
	// 需要创建一个副本用于异步处理
	copyContext := c.Copy()
	// 异步处理（先返回200，3秒后打印日志）
	go func() {
		time.Sleep(3 * time.Second)
		log.Println("异步执行：" + copyContext.Request.URL.Path)
	}()
}

// 同步函数
func longSync(c *gin.Context) {
	// 3秒后打印日志并返回200
	time.Sleep(3 * time.Second)
	log.Println("同步执行：" + c.Request.URL.Path)
}

func main() {
	r := gin.Default()
	r.GET("async", longAsync)
	r.GET("sync", longSync)
	err := r.Run(":8080")
	wrappedErr := errors.Wrap(err, "Gin服务器启动失败")
	log.Println(wrappedErr)
}
