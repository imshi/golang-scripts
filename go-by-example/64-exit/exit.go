// 使用 os.Exit 来立即进行带给定状态的退出

package main

import (
	"fmt"
	"os"
)

func main() {
	// 使用 os.Exit 时 defer 将不会 执行，所以这里的 fmt.Println将永远不会被调用
	defer fmt.Println("!")
	// 进行状态为3的退出
	os.Exit(3)
}
