// Go 中最主要的状态管理方式是通过通道间沟通完成的
// 除此之外，其他管理状态的方法还有：
// 1、使用 sync/atomic包在多个 Go 协程中进行 原子计数（上节内容）
// 2、互斥锁（本节内容）

package main

import (
	"fmt"
	"math/rand"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	var state = make(map[int]int)
	// 同步对 map 类型 state 变量访问的 mutex 互斥锁
	var mutex = &sync.Mutex{}

	// ops 将记录对 state 的操作次数
	var ops int64 = 0

	// 运行 100 个 Go 协程来重复读取 state：基于互斥锁的处理方式
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			// 每次循环读取，使用一个键来进行访问，Lock() 这个 mutex 来确保对 state 的独占访问，读取选定的键的值，Unlock() 这个mutex，并且 ops 值加 1
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)

				// 为了确保这个 Go 协程不会在调度中饿死，我们在每次操作后明确的使用 runtime.Gosched()进行释放
				runtime.Gosched()
			}
		}()
	}

	// 其他对比方式：运行 10 个 Go 协程来模拟写入操作，使用和读取相同的模式
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	// 让这 10 个 Go 协程对 state 和 mutex 的操作运行 1 s
	time.Sleep(time.Second)
	// 获取并输出最终的操作计数
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)

	// 对 state 使用一个最终的锁，显示它是如何结束的
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
