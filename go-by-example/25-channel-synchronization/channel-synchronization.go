// 可以使用通道来同步 Go 协程间的执行状态

package main

import (
	"fmt"
	"time"
)

// 定义一个在协程中运行的函数，通过一个通道参数通知其他协程该函数已经工作完毕
func worker(done chan bool) {
	fmt.Print("working...")
	time.Sleep(time.Second)
	fmt.Println("done")

	// 发送通知
	done <- true
}

func main() {
	done := make(chan bool, 1)

	// 运行一个worker Go协程,并且给予用于通知的通道
	go worker(done)

	// 程序将在接收到woeker发出通知前一直阻塞,如果去掉下面这行,程序甚至将在worker还没运行就结束了
	<-done
}
