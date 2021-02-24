// Go 的通道选择器（select）允许程序同时等待多个通道操作

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)

	// 模拟并行 Go协程中阻塞RPC的操作
	// 1秒后发送数据到通道c1中
	go func() {
		time.Sleep(time.Second * 1)
		c1 <- "one"
	}()

	// 2秒后发送数据到通道c2中
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "two"
	}()

	for i := 0; i < 2; i++ {
		// 使用select同时等待这两个值，它会不断的检测通道是否有值传过来，一旦发现传过来，立刻获取并输出
		select {
		case msg1 := <-c1:
			fmt.Println("received", msg1)
		case msg2 := <-c2:
			fmt.Println("received", msg2)
		}
	}
}
