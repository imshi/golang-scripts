// goroutine 以非阻塞的方式执行，它们会随着程序（主线程）的结束而消亡，如果不加处理就会出现协程还没执行完就退出的情况;
// 要等待所有协程执行完再退出可以使用channel(如 worker-pools.go 中接收results管道的for循环)，也可以使用sync.WaitGroup（以下内容）
// 每次运行，各个协程开启和完成的时间可能是不同的
package main

import (
	"fmt"
	"sync"
	"time"
)

// 协程并发执行的函数，WaitGroup必须通过指针传递给函数
func worker(id int, wg *sync.WaitGroup) {
	// 通知主线程，协程执行完毕
	defer wg.Done()
	fmt.Printf("Worker %d starting\n", id)
	// 睡眠一秒，用以模拟任务耗时
	time.Sleep(time.Second)
	fmt.Printf("Worker %d doned \n", id)
}

func main() {
	// 该WaitGroup用于等待函数启动所有协程
	var wg sync.WaitGroup
	// 启动 5 个协程
	for i := 1; i <= 5; i++ {
		// 递增 WaitGroup 的计数器
		wg.Add(1)
		go worker(i, &wg)
	}
	// 等待所有协程执行完毕
	wg.Wait()
}
