// Go 中最主要的状态管理方式是通过通道间沟通完成的
// 除此之外，其他管理状态的方法还有：
// 1、使用 sync/atomic包在多个 Go 协程中进行 原子计数（本节内容）：sync/atomic包提供了底层的原子级内存操作，其执行过程不能被中断，这也就保证了同一时刻一个线程的执行不会被其他线程中断，也保证了多线程下数据操作的一致性
// 2、互斥锁（下节内容）

package main

import (
	"fmt"
	"runtime"
	"sync/atomic"
	"time"
)

func main() {
	// 使用一个无符号整型数（永远是正整数）来表示这个计数器
	var ops uint64 = 0

	// 启动 50 个 Go 协程模拟并发更新，对计数器每隔 1ms 进行一次加一操作
	for i := 0; i < 50; i++ {
		go func() {
			for {
				// 使用 AddUint64 来让计数器自动增加，使用& 语法来给出 ops 的内存地址
				atomic.AddUint64(&ops, 1)
				// 允许其它 Go 协程的执行
				runtime.Gosched()
			}
		}()
	}

	// 等待一秒(让 ops 的自加操作完成)
	time.Sleep(time.Second)

	// 为了安全的使用该计数器（避免被其它 Go 协程更新），我们通过 LoadUint64 将当前值的拷贝提取到 opsFinal中（取值的内存地址 &ops）
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Println("ops:", opsFinal)
}
