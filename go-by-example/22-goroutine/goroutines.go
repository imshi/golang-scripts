// 协程：轻量级的线程，用于函数的并发运行
// Go 语言提供了 sync 和 channel 两种方式支持协程(goroutine)的并发

package main

import "fmt"

// 定义一个函数
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// 运行程序时将首先看到阻塞式调用的输出 - direct : 0/1/2
// 然后是两个 Go 协程的交替输出:表示 Go 运行时是以异步的方式运行协程的
func main() {

	// 使用一般的方式（阻塞式调用）运行
	f("direct")

	// 使用 go 协程调用（异步）
	go f("goroutine")

	// 也支持为匿名函数启动协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	// 按下任意键结束
	var input string
	fmt.Scanln(&input)
	fmt.Println("done")
}
