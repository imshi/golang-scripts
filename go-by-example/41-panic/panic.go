// 通常我们用它来表示程序正常运行中不应该出现的，或者我们没有处理好的错误

package main

import "os"

func main() {

	// 直接panic退出程序运行
	panic(" a problem")

	// 以下部分不会被执行，代码格式化工具会提示代码不可达
	_, err := os.Create("/tmp/file")
	if err != nil {
		panic(err)
	}
}
