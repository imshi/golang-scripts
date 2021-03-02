// 除了用互斥锁进行了明确的锁定来让共享的state 跨多个 Go 协程同步访问，还可以使用内置的 Go协程和通道的的同步特性来达到同样的效果（状态协程），适用场景为管理多个互斥锁容易出错的情况

package main

import (
	"fmt"
	"math/rand"
	"sync/atomic"
	"time"
)

type readOp struct {
	key  int
	resp chan int
}

type writeOp struct {
	key  int
	val  int
	resp chan bool
}

func main() {
	// 记录执行操作的次数
	var ops int64

	// reads 和 writes 通道分别将被其他 Go 协程用来发布读和写请求
	reads := make(chan *readOp)
	writes := make(chan *writeOp)

	// state 将被此协程单独拥有，能够保证数据在并行读取时不会混乱
	// 其他协程通过使用readOp和writeOp结构体封装请求，将数据发送到该协程中，然后接收对应回复（返回一个值到响应通道 resp 中表示操作成功）
	go func() {
		var state = make(map[int]int)
		for {
			select {
			case read := <-reads:
				read.resp <- state[read.key]
			case write := <-writes:
				state[write.key] = write.val
				write.resp <- true
			}
		}
	}()

	// 启动 100 个 Go 协程通过 reads 通道 对 state 所有者 Go 协程发起读取请求
	for r := 0; r < 100; r++ {
		go func() {
			for {
				read := &readOp{
					key:  rand.Intn(5),
					resp: make(chan int)}
				reads <- read
				<-read.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// 启动 10 个 Go 协程通过 writes 通道对 state 所有者 Go 协程发起写入请求
	for w := 0; w < 10; w++ {
		go func() {
			for {
				write := &writeOp{
					key:  rand.Intn(5),
					val:  rand.Intn(100),
					resp: make(chan bool)}
				writes <- write
				<-write.resp
				atomic.AddInt64(&ops, 1)
			}
		}()
	}

	// 让 Go 协程们跑 1s
	time.Sleep(time.Second)
	// 获取并报告 ops 值
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
}
