// Go中内置 bufio、io包提供I/O操作支持（读写）
// os 包的 OpenFile 函数用于打开一个文件，并返回一个文件对象，该文件对象可以用于读写操作。可以通过 flag 参数指定多种处理模式，如：只读、只写、读写、追加等。
// bufio 提供了和带缓冲的读取器一样的带缓冲的写入器

package main

import (
	"bufio"
	"fmt"
	"io/ioutil"
	"os"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// 写入一些字节/字符串到文件中（\n为换行）
	d1 := []byte("hello\ngo\n")
	err := ioutil.WriteFile("dat1", d1, 0644)
	check(err)

	// 创建一个文件进行更细粒度的写入，可以使用 Create 或者指定 os.O_CREATE 模式的 OpenFile 函数
	// f, err := os.Create("dat2")
	f, err := os.OpenFile("dat2", os.O_RDWR|os.O_CREATE, 0644)
	check(err)

	// 打开文件后，立即使用defer调用Close操作
	defer f.Close()

	// 写入字节切片（d2相当于some\n）
	d2 := []byte{115, 111, 109, 101, 10}
	n2, err := f.Write(d2)
	check(err)
	fmt.Printf("normal byte wrote %d bytes \n", n2)

	// 写入字符串
	n3, err := f.WriteString("writes\n")
	fmt.Printf("normal string wrote %d bytes\n", n3)
	check(err)

	// 调用Sync来将缓冲区的信息写入磁盘
	err = f.Sync()
	check(err)

	// 使用bufio进行带缓冲区的写入
	w := bufio.NewWriter(f)
	n4, err := w.WriteString("buffered\n")
	check(err)
	fmt.Printf("bufio wrote %d bytes\n", n4)

	// 使用Flush确保所有缓存写入底层写入器
	w.Flush()
}
