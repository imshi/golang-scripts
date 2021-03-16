// 命令行参数：Go 通过os.Args提供命令行参数访问功能（切片形式）
// os.Args[0] 是程序的路径，os.Args[1:]保存程序的所有参数

package main

import (
	"fmt"
	"os"
)

// g该程序调试要么先使用 go build编译一个可执行二进制文件，执行文件 + 参数运行
// 要么在vscode里的调试文件launch.json中添加args模块传参，program 设置为 ${file} 表示运行调试当前文件并传入命令行参数；
func main() {
	argsWithProg := os.Args
	argsWithoutProg := os.Args[1:]

	arg := os.Args[3]

	fmt.Println(argsWithProg)
	fmt.Println(argsWithoutProg)
	fmt.Println(arg)
}
