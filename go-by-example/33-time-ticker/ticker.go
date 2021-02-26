// 定时器 + 打点器：内置特性，用于实现在需要的时刻，或者某段时间间隔内重复运行代码
// 打点器 则是当你想要在固定的时间间隔重复执行准备的

package main

import (
	"fmt"
	"time"
)

// 该程序将在打点器停止前打点3次
func main() {

	ticker := time.NewTicker(time.Millisecond * 500)
	// 使用内置的 range 来迭代值每隔500ms 发送一次的值
	go func() {
		// C为数据类型为Time的单向管道，只能读/发送，不能写/接收
		for t := range ticker.C {
			fmt.Println("Tick at ", t)
		}
	}()

	time.Sleep(time.Millisecond * 1600)
	// 打点器可以和定时器一样被停止。与定时器不同的是：一旦一个打点停止了，将不能再从它的通道中接收到值
	ticker.Stop()
	fmt.Println("Ticker stopped")
}
