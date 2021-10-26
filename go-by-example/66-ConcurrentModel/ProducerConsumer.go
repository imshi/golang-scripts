// 生产者消费者模型是非常常见的并发模型，golang提供了chan类型，可以很方便的实现生产者和消费者之间的数据和状态同步
/*
【场景】：
1.多个生产者，多个消费者，且生产速度高于消费者消费速度
2.数据同步：程序中止信号发出，生产者暂停生产并退出线程，消费者继续消费，直到缓存数据被消费完
3.使用sync.WaitGroup等待线程结束

【案例来自】：https://www.cnblogs.com/ChinaHook/p/14699627.html
*/

package main

import (
	"fmt"
	"strconv"
	"sync"
	"time"
)

// 暂停标志
var bStop = false

// 模拟程序收到暂停信号，或者超时信号使程序停止，一般情况是通过通道来进行标志，本文通过bool来进行标志
func makeStop() {
	time.Sleep(time.Second * 3)
	bStop = true
}

// 生产者：如果暂停标志为false，则每隔2秒生产一条消息
func producer(wg *sync.WaitGroup, threadId int, ch chan string) {
	count := 0
	// 标志位为false，不断写入数据
	for !bStop {
		// 模拟生产数据耗时
		time.Sleep(time.Second * 2)
		count++
		// strconv.Itoa：将int转化为string
		data := strconv.Itoa(threadId) + "<-ThreadId---No->" + strconv.Itoa(count)
		fmt.Println("producer:", data)
		ch <- data
	}
	// 等待协程运行完毕
	wg.Done()
}

// 消费者：不断读取，直到通道关闭
func consumer(wg *sync.WaitGroup, ch chan string) {
	// 不断读取，直到通道关闭
	for data := range ch {
		// 模拟消费数据耗时
		time.Sleep(time.Second * 2)
		fmt.Println("consumer", data)
	}
	// 等待协程运行完毕
	wg.Done()
}

func main() {
	// 缓存：存储消息，用来模拟生产者完成生产，消费者未完成消费的场景
	chanStream := make(chan string, 30)

	// 生产者和消费者计数器
	wgProducer := new(sync.WaitGroup)
	wgConsumer := new(sync.WaitGroup)

	// 三个生产者
	for i := 0; i < 3; i++ {
		wgProducer.Add(1)
		go producer(wgProducer, i, chanStream)
	}

	// 两个消费者
	for j := 0; j < 2; j++ {
		wgConsumer.Add(1)
		go consumer(wgConsumer, chanStream)
	}

	// 模拟程序暂停或者超时
	go makeStop()

	// 等候所有生产者协程运行完毕（生产完成）
	wgProducer.Wait()
	// 当生产完成后，关闭通道，消费者发现通道关闭，对应线程才会退出，即：生产者完成生产，通过关闭通道告诉消费者，我已完成生产，你消费完剩余数据就退出吧~
	close(chanStream)

	// 等候所有消费者协程运行完毕（消费完成）
	wgConsumer.Wait()
}
