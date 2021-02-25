// 超时在需要花费时间、或者连接外部资源的程序时很有必要的；
// golang 中得益于 channel 和 select，可以实现简洁优雅的超时操作

package main

import (
	"fmt"
	"time"
)

func main() {
	c1 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c1 <- "result 1"
	}()

	// 使用 select 实现超时操作（设置超时阈值为1秒，协程函数2秒后会发送数据到通道，结果从c1接收数据超时）
	select {
	case res := <-c1:
		fmt.Println("res:", res)
	// time.After()函数返回一个Time类型的单向channel，其type为：<-chan time.Time
	case <-time.After(time.Second * 1):
		fmt.Println("timeout 1")
	}

	c2 := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		c2 <- "result 2"
	}()

	// 使用 select 模拟超时（设置超时阈值为3秒，协程函数2秒后会发送数据到通道，结果成功从c2接收到值并打印）
	select {
	case res := <-c2:
		fmt.Println(res)
	case <-time.After(time.Second * 3):
		fmt.Println("timeout 2")
	}
}
