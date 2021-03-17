//  Go 中生成其他非 Go 进程：os/exec包提供支持
package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
)

// 该脚本仅限 Linux 系统运行
func main() {
	// exec.Command 函数创建一个表示外部进程的对象，以下功能是：仅打印一些信息到标准输出流
	dateCmd := exec.Command("date")
	// .Output 是一个处理运行命令的常见函数，它等待命令运行完成，并收集命令的输出
	dateOut, err := dateCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> date")
	// 格式化为字符串输出
	fmt.Println(string(dateOut))

	// 以下是更复杂的例子：从外部进程的stdin 输入数据并从 stdout 收集结果
	grepCmd := exec.Command("grep", "hello")
	// 明确的获取输入/输出管道，运行这个进程，写入一些输入信息，读取输出的结果，最后等待程序运行结束
	// 同理可收集 StderrPipe 的结果
	// 以下例子中忽略了错误检测（不建议）
	grepIn, _ := grepCmd.StdinPipe()
	grepOut, _ := grepCmd.StdoutPipe()
	grepCmd.Start()
	grepIn.Write([]byte("hello grep\ngoodbye grep"))
	grepIn.Close()
	// 读取输出的结果
	grepBytes, _ := ioutil.ReadAll(grepOut)
	// 等待程序运行结束
	grepCmd.Wait()

	fmt.Println("> grep hello")
	// 格式化为字符串输出
	fmt.Println(string(grepBytes))

	// 使用 bash命令的 -c 选项：通过一个字符串生成一个完整的命令（包含命令和参数数组）
	lsCmd := exec.Command("bash", "-c", "ls -l")
	lsOut, err := lsCmd.Output()
	if err != nil {
		panic(err)
	}
	fmt.Println("> ls -l")
	// 格式化为字符串输出
	fmt.Println(string(lsOut))

}
