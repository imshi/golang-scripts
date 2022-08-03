// 一般情况下，goroutine在操作系统上只要你的硬件资源够它是可以无限启动的，如果出现大规模的启动goroutine的情况会造成大量占用系统资源；工作池的目的就是为了限制golang的启动数量，保证不会出现硬件计算资源溢出的情况
// 利用 channel 带缓冲和阻塞的特性来实现控制并发数量（协程池）
// 以下示例中使用了channel阻塞的特性保证子协程执行完毕后再退出主程序，可以使用 waitgroup 更优雅的实现同样效果，参考：34-WaitGroup章节
// 三方包：Jeffail/tunny 或者 panjf2000/ants 可以帮助我们实现指定数量的携程池，避免耗费太多资源。

package main

import (
	"fmt"
	"time"
)

func main() {
	count := 10 // 最大支持并发
	sum := 30   // 任务总数

	c := make(chan int, count)     // 控制任务并发的chan
	sc := make(chan struct{}, sum) // 和实际处理业务无关，防止主程序退出之后，子协程没有执行完毕；
	defer close(c)
	defer close(sc) // 和实际处理业务无关，防止主程序退出之后，子协程没有执行完毕；

	for i := 0; i < sum; i++ {
		c <- i // 作用类似于waitgroup.Add(1)
		// 作为示例，这里是个匿名函数，实际使用中可以将协程函数拆出去，放到main函数外面
		go func(j int) {
			fmt.Println(j)
			time.Sleep(time.Second) // 模拟耗时任务
			<-c                     // 执行完毕，释放资源
			sc <- struct{}{}
		}(i)
	}

	// 和实际处理业务无关，防止主程序退出之后，子协程没有执行完毕；
	for i := 0; i < sum; i++ {
		<-sc
	}
	// // 处理实际切片入参的场景：并发处理 15 个任务
	// items := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15, 16, 17, 18, 19, 20}
	// for _, item := range items {
	// 	c <- item	// 作用类似于waitgroup.Add(1)
	// 	go func(j int) {
	// 		fmt.Println(j)
	// 		time.Sleep(time.Second) // 模拟耗时任务
	// 		<-c                     // 执行完毕，释放资源
	// 	}(item)
	// }
}
