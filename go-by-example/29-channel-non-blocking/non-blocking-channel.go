// 带一个 default 子句的 select 来实现非阻塞 的发送、接收，甚至是非阻塞的多路 select

package main

import "fmt"

func main() {
	message := make(chan string)
	signals := make(chan bool)

	// 非阻塞接收的例子
	select {
	case msg := <-message:
		fmt.Println("received message: ", msg)
	default:
		fmt.Println("no message received")
	}

	// 非阻塞发送的例子
	msg := "hi"
	select {
	case message <- msg:
		fmt.Println("sent message", msg)
	default:
		fmt.Println("no message sent")
	}

	// 在 default 前使用多个 case 子句来实现一个多路的非阻塞的选择器
	select {
	case msg := <-message:
		fmt.Println("received message", msg)
	case sig := <-signals:
		fmt.Println("received signal", sig)
	default:
		fmt.Println("no acticity")
	}
}
