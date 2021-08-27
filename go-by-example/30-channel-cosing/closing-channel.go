// 用于给通道的接收方传达工作已经完成的信息
// 一个非空的通道也是可以关闭的，通道中剩下的值仍然可以被接收到
package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			// more 用来判定接收完成与否，如果 jobs 已经关闭了、并且通道中所有的值都已经接收完毕，那么 more 的值将是 false
			j, more := <-jobs
			if more {
				fmt.Println("received jobs", j)
			} else {
				// 完成所有的任务时，通过 done 通道去进行通知
				fmt.Println("received all jobs")
				done <- true
				return
			}
		}
	}()

	// 使用 jobs 发送 3 个任务到工作函数中，然后关闭 jobs
	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job", j)
	}
	close(jobs)

	// 使用通道同步方法等待任务结束
	fmt.Println("sent all jobs")
	<-done
}
