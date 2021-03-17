// Go 中通过通道来处理 Unix 信号，如：希望一个命令行工具在接收到一个 SIGINT 信号时退出程序
package main

import (
	"fmt"
	"os"
	"os/signal"
	"syscall"
)

// 当我们运行这个程序时，它将一直等待一个信号。使用 ctrl-C（终端显示为 ^C），我们可以发送一个 SIGINT 信号，这会使程序打印 interrupt 然后退出
func main() {

	// 创建一个通道来接收来自 os.Signal 值的信号通知：Signal类型，容量为1
	sigs := make(chan os.Signal, 1)
	// 同时创建一个用于通知可结束程序运行的通道：boo类型，容量为1
	done := make(chan bool, 1)

	// signal.Notify注册 signs 通道来接收特定信号
	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	// 这个 Go 协程执行一个阻塞的操作接收信号。当它得到一个值时，它将打印这个值，然后通知程序可以退出
	go func() {
		sig := <-sigs
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	// 程序将在这里进行等待，直到它得到了期望的信号（也就是上面的 Go 协程发送的 done 值）然后退出
	//
	fmt.Println("awaiting signal...")
	<-done
	fmt.Println("exiting.")
}
