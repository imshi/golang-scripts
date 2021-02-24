// 协程：轻量级的线程

package main

import "fmt"

// 定义一个函数
func f(from string) {
	for i := 0; i < 3; i++ {
		fmt.Println(from, ":", i)
	}
}

// 运行程序时将首先看到阻塞式调用的输出，然后是两个 Go 协程的交替输出:表示 Go 运行时是以异步的方式运行协程的
func main() {

	// 使用一般的方式（阻塞式调用）运行
	f("direct")

	// 使用 go 协程调用
	go f("goroutine")

	// 也支持为匿名函数启动协程
	go func(msg string) {
		fmt.Println(msg)
	}("going")

	var input string
	// 按下任意键结束
	fmt.Scanln(&input)
	fmt.Println("done")
}
