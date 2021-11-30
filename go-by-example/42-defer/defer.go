// 延时函数（defer）：用来确保一个函数调用在程序执行结束（在封闭函数（main）结束时）前执行，通常用于 open/close, connect/disconnect, lock/unlock 等这些成对的操作, 来保证在任何情况下资源都被正确释放；
// defer 和 recover 机制也用于异常处理，类似于python中的：try... finally（try 中捕获各种类型的异常，在 catch 中定义异常处理的行为）

package main

import (
	"fmt"
	"os"
)

// 创建文件
func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

// 文件写入
func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

// 关闭文件
func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}

// 使用 defer 定义异常处理的函数，在协程退出前，会执行完 defer 挂载的任务。因此如果触发了 panic，控制权就交给了 defer；在 defer 的处理逻辑中，使用 recover 使程序恢复正常，并且将返回值设置为 -1，在这里也可以不处理返回值，如果不处理返回值，返回值将被置为默认值 0
func get(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some error happended", r)
			ret = -1
		}
	}()
	arr := [3]int{2, 3, 4}
	return arr[index]
}

func main() {
	// 异常捕获
	fmt.Println(get(5))
	fmt.Println("finished")

	fmt.Println("--------------------------")

	// 创建一个文件并写入，在程序运行结束之前关闭文件
	f := createFile("defer.txt")
	defer closeFile(f)
	writeFile(f)
}
