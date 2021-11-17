/*
验证：
1、curl -i -X GET 'http://127.0.0.1:8080/someJSON'
2、curl -i -X GET 'http://127.0.0.1:8080/someYAML'
3、curl -i -X GET 'http://127.0.0.1:8080/someProtobuf'
*/
package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/testdata/protoexample"
	"github.com/pkg/errors"
)

// protobuf格式，谷歌开发的高效存储读取的工具
func responseProtubuf(c *gin.Context) {
	reps := []int64{int64(1), int64(2), int64(3)}
	// 定义数据
	label := "label"
	// 传 protobuf 格式数据
	data := &protoexample.Test{
		Label: &label,
		Reps:  reps,
	}
	c.ProtoBuf(200, data)
}

func setupRouter() *gin.Engine {
	r := gin.Default()
	// JSON响应
	r.GET("/someJSON", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "someJSON", "status": 200})
	})
	// YAML响应
	r.GET("/someYAML", func(c *gin.Context) {
		c.YAML(200, gin.H{"message": "someYAML", "status": 200})
	})
	// ProtoBuf响应
	r.GET("/someProtobuf", responseProtubuf)
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
