// 默认通道是 无缓冲 的，这意味着只有在对应的接收（<- chan）通道准备好接收时，才允许进行发送（chan <-）。可缓存通道允许在没有对应接收方的情况下，缓存限定数量的值
// 带缓冲的channel，可用于限制并发的协程数量；纯sync.WaitGroup的方式，可以让协程并发执行，但是不能限制并发数量。
package main

import (
	"fmt"
	"log"
	"sync"
	"time"
)

func main() {

	// 定义一个最多允许缓存2个值的通道
	messages := make(chan string, 2)

	// 发送两个值到通道
	messages <- "buffered"
	messages <- "channel"

	// 从通道中接收内容
	fmt.Println(<-messages)
	fmt.Println(<-messages)

	fmt.Println("--- 使用带缓冲的channel限制并发的协程数量 ---")
	// sync.WaitGroup 用于等待所有协程执行完毕，并不是必须的，例如 http 服务，每个请求天然是并发的，此时使用 channel 控制并发处理的任务数量即可，就不需要 sync.WaitGroup 了。
	var wg sync.WaitGroup
	// 创建缓冲区为 3 的空类型 channel，在没有被接收的情况下，最多接收 3 个消息就会被阻塞
	ch := make(chan struct{}, 3)
	// 并发处理 10 个任务，任务入参来自一个整形切片
	items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for _, item := range items {
		// 开启协程前，向 channel 发送一个空消息占位，再通过缓冲区的数量3限制并发为3，超过 1+3 则阻塞
		ch <- struct{}{}
		// 递增 WaitGroup 的计数器
		wg.Add(1)
		go func(item int) {
			// 当任务执行完毕后，从 wg 中减去 1 个计数器
			defer wg.Done()
			log.Printf("worker: %d, \n", item)
			time.Sleep(time.Second)
			// 协程结束前，调用 <-ch 释放缓冲区
			<-ch
		}(item)
	}
	// 等待所有协程执行完毕
	wg.Wait()
}
