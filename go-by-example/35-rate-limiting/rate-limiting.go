// 速率限制是一个重要的控制服务资源利用和质量的途径
// Go 通过 Go 协程、通道和打点器优美的支持了速率限制

package main

import (
	"fmt"
	"time"
)

// 该程序运行效果为：
// 第一批请求：大约每 1000ms 处理一次
// 第二批请求：直接连续处理了 3 次，然后大约每 1000ms 一次处理其余的 2 个请求
func main() {
	requests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		requests <- i
	}

	close(requests)

	// 这个 limiter 通道将每 1000ms 接收一个值
	limiter := time.Tick(time.Millisecond * 1000)

	for req := range requests {
		// 通过在每次请求前阻塞 limiter 通道的一个接收， 可以将频率限制为：每 1000ms 执行一次请求
		<-limiter
		fmt.Println("requests: ", req, time.Now())
	}

	fmt.Println("-------------------------------------")
	// 如果需要临时调整速率限制，并且不影响整体的速率控制，可以通过通道缓冲来实现
	// 此处 burstyLimiter 通道允许最多 3 个爆发（bursts）事件
	burstyLimiter := make(chan time.Time, 3)

	// 填充通道，表示允许的爆发（bursts）
	for i := 0; i < 3; i++ {
		burstyLimiter <- time.Now()
	}

	// 每 1000 ms 添加一个新的值到 burstyLimiter中，直到达到 3 个的限制
	go func() {
		for t := range time.Tick(time.Millisecond * 1000) {
			burstyLimiter <- t
		}
	}()

	// 模拟 5 个传入请求；受益于 burstyLimiter 的爆发（bursts）能力，前 3 个请求可以快速完成
	burstyRequests := make(chan int, 5)
	for i := 0; i < 5; i++ {
		burstyRequests <- i
	}
	close(burstyRequests)

	for req := range burstyRequests {
		<-burstyLimiter
		fmt.Println("requests: ", req, time.Now())
	}
}
