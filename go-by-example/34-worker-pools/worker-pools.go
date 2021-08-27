// 一般情况下，goroutine在操作系统上只要你的硬件资源够它是可以无限启动的，如果出现大规模的启动goroutine的情况会造成大量占用系统资源；工作池的目的就是为了限制golang的启动数量，保证不会出现硬件计算资源溢出的情况
// 使用 Go 协程和通道实现一个工作池

package main

import (
	"fmt"
	"time"
)

// 定义在并发实例中执行的任务：并发数为id的值，从job通道中接收任务，通过results通道发送对应结果
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("worker", id, "processing job", j)
		// 间隔1秒来模拟一个耗时的任务
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	// 初始化带缓冲区的channel
	jobs := make(chan int, 100)
	results := make(chan int, 100)

	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 9; j++ {
		jobs <- j
	}

	close(jobs)
	for a := 1; a <= 9; a++ {
		<-results
	}
}
