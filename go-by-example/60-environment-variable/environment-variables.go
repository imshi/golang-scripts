// 环境变量：为Unix程序传递配置信息的普遍方式（以下程序在Windows中也适用）

package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {

	// 使用 os.Sentenv设置一个键值对；
	os.Setenv("FOO", "1")
	// 使用 os.Getenv获取一个对应的键值对（不存在的话返回一个空字符串）；
	fmt.Println("FOO:", os.Getenv("FOO"))
	fmt.Println("BAR:", os.Getenv("BAR"))

	fmt.Println()
	// 遍历系统环境变量
	// 使用 os.Environ列出所有环境变量的键值对
	for _, e := range os.Environ() {
		// 提取并打印键名
		pair := strings.Split(e, "=")
		fmt.Println(pair[0])
	}
}
