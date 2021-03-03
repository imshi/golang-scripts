// 延时函数（defer）：用来确保一个函数调用在程序执行结束（在封闭函数（main）结束时）前执行，通常用于 open/close, connect/disconnect, lock/unlock 等这些成对的操作, 来保证在任何情况下资源都被正确释放，类似于python中的：try... finally
package main

import (
	"fmt"
	"os"
)

// 创建一个文件并写入，在程序运行结束之前关闭文件
func main() {
	f := createFile("defer.txt")
	defer closeFile(f)
	writeFile(f)
}

func createFile(p string) *os.File {
	fmt.Println("creating")
	f, err := os.Create(p)
	if err != nil {
		panic(err)
	}
	return f
}

func writeFile(f *os.File) {
	fmt.Println("writing")
	fmt.Fprintln(f, "data")
}

func closeFile(f *os.File) {
	fmt.Println("closing")
	f.Close()
}
