// 内置的 strconv 包提供了数字解析功能，可用于从字符串中解析数字等
// cast 包提供了一个简单的类型转换函数（To_ 形式函数），用于将 interface{} 类型转换为其他类型，如果无法正确转换为对应的类型，则返回目标类型的零值
package main

import (
	"fmt"
	"strconv"

	"github.com/spf13/cast"
)

func main() {
	// 将string字符串解析为64位浮点数
	f, _ := strconv.ParseFloat("1.234", 64)
	fmt.Println(f)
	// cast包
	fmt.Println(cast.ToFloat64("1.234"))

	// 将string解析为Int型，0表示自动推断进制，64表示使用64位存储
	i, _ := strconv.ParseInt("123", 0, 64)
	fmt.Println(i)

	// ParseInt 会自动识别十六进制数
	d, _ := strconv.ParseInt("0x1c8", 0, 64)
	fmt.Println(d)
	// cast包
	fmt.Println(cast.ToInt(0x1c8))

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
