// runtime.Gosched()：暂停当前 goroutine（未来会继续执行），让出CPU，调度器安排其他等待的任务运行；

package main

import (
	"fmt"
	"runtime"
)

func main() {
	// 让子协程先执行
	go func(s string) {
		for i := 0; i < 5; i++ {
			fmt.Println(s)
		}

	}("go")
	// time.Sleep(time.Second)

	// 主协程
	for i := 0; i < 2; i++ {
		// 让出时间片，先让别的协议执行，它执行完，再回来执行此协程
		runtime.Gosched()
		fmt.Println("hello")
	}
}
