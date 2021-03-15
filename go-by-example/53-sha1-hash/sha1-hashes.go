// SHA1 散列经常用生成二进制文件或者文本块的短标识；
// Go 在多个 crypto/* 包中实现了一系列散列函数

package main

import (
	"crypto/sha1"
	"fmt"
)

func main() {
	s := "sha1 this string"

	// 产生一个散列值的方式是：sha1.New()
	h := sha1.New()

	// 写入要处理的字节
	// 字符串需要使用[]byte(s) 来强制转换成字节数组
	h.Write([]byte(s))

	// 用来得到最终的散列值的字符切片，一般不需要参数（用来都现有的字符切片追加额外的字节切片）
	bs := h.Sum(nil)

	fmt.Println(s)
	// 格式化为 16 进制字符串
	fmt.Printf("%x\n", bs)
}
