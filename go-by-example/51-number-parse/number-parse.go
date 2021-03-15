// 内置的 strconv 包提供了数字解析功能，可用于从字符串中解析数字等
package main

import (
	"fmt"
	"strconv"
)

func main() {
	// 将string字符串解析为64位浮点数
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)

	// 将string解析为Int型，0表示自动推断进制，64表示使用64位存储
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt 会自动识别十六进制数
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)

	// ParseUint识别无符号整型
	u, _ := strconv.ParseUint("789", 0, 64)
	fmt.Println(u)

	// Atoi 是一个基础的10进制整型转换函数
	k, _ := strconv.Atoi("135")
	fmt.Println(k)

	// 输入错误时，解析会返回一个err
	_, e := strconv.Atoi("wat")
	fmt.Println(e)

}
