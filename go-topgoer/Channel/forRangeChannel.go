package main

import "fmt"

// 可以通过close函数关闭通道来告知从该通道接收值的goroutine停止等待,当通道被关闭时，往该通道发送值会引发panic，从该通道里接收的值一直都是类型零值

// 判断通道是否被关闭，我们通常使用的是for range的方式，通道关闭后会退出for range循环
func main() {
	ch1 := make(chan int)
	// 开启goroutine将0~100的数发送到ch1中
	go func() {
		for i := 0; i < 100; i++ {
			ch1 <- i
		}
		close(ch1)
	}()

	// 在主goroutine中从ch2中接收值打印，通道关闭后会退出for range循环
	for i := range ch1 {
		fmt.Println(i)
	}
}
