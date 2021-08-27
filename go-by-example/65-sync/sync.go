package main

import (
	"fmt"
	"sync"
	"time"
)

// Go 语言提供了 sync 和 channel 两种方式支持协程(goroutine)的并发
// sync.WaitGroup解决同步阻塞等待的问题：优雅解决一个人等待一堆人干完活的问题（一个 WaitGroup 会等待一系列 goroutine 直到它们全部运行完毕为止）， 主 goroutine 通过调用 Add() 方法来设置需要等待的 goroutine 数量， 而每个运行的 goroutine 则在它们运行完毕时调用 Done() 方法。 与此同时， 调用 Wait() 方法可以阻塞直到所有 goroutine 都运行完毕为止

// 串行需要 3s 的下载操作，并发后，只需要 1s
var wg sync.WaitGroup

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second) //模拟耗时操作
	wg.Done()               //减去一个计数
}

func main() {
	for i := 0; i < 3; i++ {
		wg.Add(1) //为 wg 添加一个计数
		go download("a.com/" + string(i+'0'))
	}

	wg.Wait() //等待所有的协程执行结束
	fmt.Println("Done!")
}
