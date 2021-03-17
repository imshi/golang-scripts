// 使用其他进程（也许是非go进程）完全替代当前的 Go 进程：syscall.Exec方法提供支持

package main

import (
	"os"
	"os/exec"
	"syscall"
)

func main() {
	// 执行 ls 命令；需要可执行文件的绝对路径，这里使用 exec.LookPath 来获取
	binary, lookErr := exec.LookPath("ls")
	if lookErr != nil {
		panic(lookErr)
	}

	// Exec 需要切片形式的参数（不是放在一整个大字符串），且第一个参数需要是程序名
	// 以下例子中给 ls 一些基本的参数： -a -l -h
	args := []string{"ls", "-a", "-l", "-h"}

	// Exec 同样需要使用环境变量。这里我们仅提供当前的环境变量
	env := os.Environ()

	// 通过 os.Exec 调用，有三个参数：可执行文件绝对路径、参数、环境变量
	execErr := syscall.Exec(binary, args, env)
	if execErr != nil {
		panic(execErr)
	}
}
