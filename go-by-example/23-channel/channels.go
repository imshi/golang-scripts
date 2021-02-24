// 通道 是连接多个 Go 协程的管道。你可以从一个 Go 协程将值发送到通道，然后在别的 Go 协程中接收
// 通道默认是阻塞的：直到发送方和接收方都准备完毕才会开始发送、接收；

package main

import "fmt"

func main() {
	// 定义一个string类型的通道
	messages := make(chan string)

	// 匿名函数发送一个新值到通道中
	go func() {
		messages <- "ping"
	}()

	// 从通道中接收消息
	msg := <-messages
	fmt.Println(msg)
}
