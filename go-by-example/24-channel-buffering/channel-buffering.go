// 默认通道是 无缓冲 的，这意味着只有在对应的接收（<- chan）通道准备好接收时，才允许进行发送（chan <-）。可缓存通道允许在没有对应接收方的情况下，缓存限定数量的值

package main

import "fmt"

func main() {
	// 定义一个最多允许缓存2个值的通道
	messages := make(chan string, 2)

	// 发送两个值到通道
	messages <- "buffered"
	messages <- "channel"

	// 从通道中接收内容
	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
