// 生产消费者模型
// 生产者每秒生成一个字符串，并通过通道传给消费者，生产者使用两个 goroutine 并发运行，消费者在 main() 函数的 goroutine 中进行处理
package main

import (
	"fmt"
	"math/rand"
	"time"
)

// 数据生产者，传入一个标记类型的字符串及一个只能写入的通道
func producer(header string, channel chan<- string) {
	// 不停的产生数据
	for {
		// 将header和随机数格式化为字符串并写入通道
		channel <- fmt.Sprintf("%s: %v", header, rand.Int31())
		// 等待1秒
		time.Sleep(time.Second)
	}
}

// 数据消费者，传入一个只能写入的通道
func customer(channel <-chan string) {
	// 不停的获取数据
	for {
		// 从通道中取出数据，会阻塞直到信道中返回数据
		message := <-channel
		// 打印取出的数据
		fmt.Println(message)
	}
}

func main() {
	// 实例化一个字符串类型的通道
	channel := make(chan string)
	// 并发执行一个生产者函数，两行分别创建了这个函数搭配不同参数的两个goroutine
	go producer("cat", channel)
	go producer("dog", channel)
	customer(channel)
}
