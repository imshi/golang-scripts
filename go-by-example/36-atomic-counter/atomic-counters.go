// Go 中最主要的状态管理方式是通过通道间沟通完成的
// 除此之外，其他管理状态的方法还有：
// 1、使用 sync/atomic包在多个 Go 协程中进行 原子计数（本节内容）：sync/atomic包提供了底层的原子级内存操作，其执行过程不能被中断，这也就保证了同一时刻一个线程的执行不会被其他线程中断，也保证了多线程下数据操作的一致性
// 2、互斥锁（下节内容）

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

func main() {
	// 使用一个无符号整型数（永远是正整数）来表示这个计数器
	var ops uint64 = 0

	// 使用 WaitGroup 确认所有协程均执行完毕
	var wg sync.WaitGroup
	// 启动 50 个协程，每个协程会将计数器递增 1000 次，最终 ops 的值为 50000
	for i := 0; i < 50; i++ {
		// wg 计数器加 1
		wg.Add(1)
		go func() {
			// 当前协程执行完毕后，从 wg 中减去 1 个计数器
			defer wg.Done()
			for c := 0; c < 1000; c++ {
				// 使用 AddUint64 来让计数器自动增加，使用 & 语法来给出 ops 的内存地址
				atomic.AddUint64(&ops, 1)
				// 允许其它 Go 协程的执行
				runtime.Gosched()
			}
		}()
	}

	// 等待所有 Go 协程执行完毕
	wg.Wait()

	// atomic.LoadUint64 函数允许在原子更新的同时安全地读取，即：读取 ops 的值的同时，当前计算机中的任何CPU都不会进行其它的针对此值的读或写操作
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
