// 使用 os.Open() 打开文件，并返回文件对象供读写使用；任务结束后要关闭文件：建议这个操作在 Open操作后立即使用 defer 来完成
// Go中内置 bufio、io包提供I/O操作支持（读写）
// bufio 包实现了带缓冲的读取，能提升小读取操作的性能，并提供很多附加读取函数（推荐使用）

package main

import (
	"bufio"
	"fmt"
	"io"
	"io/ioutil"
	"os"
	"path/filepath"
)

// 抽离错误检查函数以便复用（读取文件需要经常进行错误检查）
func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	// 路径操作
	// Go 不是解释型语言，它的路径不是 .go 文件的路径，而是编译出来的可执行文件的路径
	// 方式一
	abspath, err := filepath.Abs(".")
	if err != nil {
		panic(err)
	}
	fmt.Println("当前工程目录绝对路径：", abspath)
	absbase := filepath.Base(abspath)
	fmt.Println("当前目录路径Base文件夹：", absbase)

	// 方式二
	abspath02, _ := os.Getwd()
	fmt.Println("当前工程目录绝对路径：", abspath02)

	// 获取当前操作系统目录分隔符(windows输出：\；linux输出：/)
	sep := string(filepath.Separator)
	fmt.Println("当前系统目录分隔符为：", sep)

	fmt.Println("=== 使用ioutil一次性读取：")
	dat, err := ioutil.ReadFile("dat")
	check(err)
	fmt.Print(string(dat))

	fmt.Println("=== 使用 io 进行按字节读取：")
	// 使用 os.Open打开一个文件获取一个 os.File 实例对象
	f, err := os.Open("dat")
	check(err)
	// 任务结束后要关闭文件：建议这个操作在 Open操作后立即使用 defer 来完成。
	defer f.Close()

	// 从文件开始位置读取一些字节。这里最多读取 5 个字节
	b1 := make([]byte, 5)
	n1, err := f.Read(b1)
	check(err)
	fmt.Printf("%d bytes:%s \n", n1, string(b1))

	fmt.Println("=== 指定位置读取：")
	// 也可以 Seek 到一个文件中已知的位置并从这个位置开始进行读取
	o2, err := f.Seek(6, 0)
	check(err)
	b2 := make([]byte, 2)
	n2, err := f.Read(b2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s \n", n2, o2, string(b2))

	// io 包提供了一些可以进行文件读取的函数，如下基于 ReadAtLeast 的实现更为健壮
	o3, err := f.Seek(6, 0)
	check(err)
	b3 := make([]byte, 2)
	n3, err := io.ReadAtLeast(f, b3, 2)
	check(err)
	fmt.Printf("%d bytes @ %d: %s\n", n3, o3, string(b3))

	// 没有内置的回转支持，但是使用 Seek(0, 0) 实现
	_, err = f.Seek(0, 0)
	check(err)

	// bufio bufio 包实现了带缓冲的读取，这不仅对有很多小的读取操作的能提升性能，也提供了很多附加的读取函数
	fmt.Println("=== 使用 bufio 实现带缓冲的读取：")
	r4 := bufio.NewReader(f)
	// 推荐使用 ReadString 或 ReadBytes 方法进行按行读取，会持续读取直到遇到指定的分隔符，并返回一个字符串、一个字节切片；
	b4, err := r4.ReadString('\n')
	check(err)
	fmt.Printf("The First Line: %s\n", string(b4))

	fmt.Println("=== 推荐的按行读取整个文件的方法：")
	file, err := os.Open("dat")
	check(err)
	defer file.Close()
	r5 := bufio.NewReader(file)
	for {
		b5, err := r5.ReadBytes('\n')
		if err == io.EOF {
			break
		}
		fmt.Printf("%s", string(b5))
	}

	// 任务结束后要关闭这个文件（通常这个操作应该在 Open操作后立即使用 defer 来完成）
	// f.Close()
}
