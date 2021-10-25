// runtime.Goexit()：终止调用它的go协程，其他协程不受影响；Goexit()会在终止该go协程前执行所有的defer函数，前提是defer必须在它前面定义；【注】：千万别在主函数调用Goexit()，会引发panic。

// 只调用了aaa和ccc相关的行。后面的ddd和bbb相关的行由于所在的协程被终止了，所以后面的结果无法执行
package main

import (
	"fmt"
	"runtime"
	"time"
)

func test() {
	defer fmt.Println("ccc")
	runtime.Goexit() // 终止所在的线程，Goexit()会在终止该go协程前执行所有的defer函数，前提是defer必须在它前面定义；
	fmt.Println("ddd")

}

func main() {
	// 创建新的协程
	go func() {
		fmt.Println("aaa")
		// 调用别的函数
		test()
		fmt.Println("bbb")
	}()
	// 延迟结束主协程
	// ddd和bbb相关的行由于所在的协程被终止了，所以后面的结果无法执行
	time.Sleep(10 * time.Second)
}
