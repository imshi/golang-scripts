// 通道 是连接多个 Go 协程的管道，用于在协程之间传递数据。
// 通道 总是遵循先入先出（First In First Out）的规则，保证收发数据的顺序和安全性
// 通道默认是阻塞的：直到发送方和接收方都准备完毕才会开始发送、接收；

package main

import (
	"fmt"
	"time"

	"github.com/spf13/cast"
)

var ch002 = make(chan string, 10) // 定义一个string类型、大小为 10 的缓冲通道

func download(url string) {
	fmt.Println("start to download", url)
	time.Sleep(time.Second)
	ch002 <- url // 将 url 发送给信道
}

func main() {
	fmt.Println("=== Sample 001 ===")
	// 定义一个string类型、大小为 10 的缓冲通道
	messages := make(chan string, 10)

	// 匿名函数发送一个新值到通道中
	go func() {
		messages <- "ping"
	}()

	// 从通道中接收消息
	msg := <-messages
	fmt.Println(msg)

	fmt.Println("=== Sample 002 ===")
	// 使用 channel 信道，在协程之间传递数据，确保同步交换数据，保证数据安全  - 示例 1
	for i := 0; i < 3; i++ {
		// go download("a.com/" + string(i+'0'))
		go download("a.com/" + cast.ToString(i+0))
	}

	for i := 0; i < 3; i++ {
		msg := <-ch002 // 等待信道返回消息
		fmt.Println("finish ", msg)
	}
	fmt.Println("Done!")

	fmt.Println("=== Sample 003 ===")
	// 使用 channel 信道，在协程之间传递数据，确保同步交换数据，保证数据安全  - 示例 2
	ch := make(chan int) //构建一个通道

	// 开启一个并发匿名函数
	go func() {
		fmt.Println("Start goroutine")
		// 匿名 goroutine 即将结束时，使用通道通知 goroutine，这一句会一直阻塞直到 goroutine 接收为止
		ch <- 0
		fmt.Println("Exit goroutine")
	}()

	fmt.Println("Wait goroutine")

	// 开启 goroutine 后，使用通道等待匿名 goroutine 结束
	<-ch
	fmt.Println("all done")

}
