// 可以设置单向通道:只用来接收或者发送值,提升程序类型安全性

package main

import "fmt"

// 该函数中的通道只允许发送数据到通道
func ping(pings chan<- string, msg string) {
	pings <- msg
}

// 该函数有两个参数：从pings通道中接收收据并发送到pongs通道中
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	pongs <- msg
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)

	ping(pings, "passed message")
	// 从pings通道中接收收据并发送到pongs通道中
	pong(pings, pongs)
	// 输出 pongs 通道中的数据
	fmt.Println(<-pongs)
}
