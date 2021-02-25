// 使用const关键字声明；支持字符、字符串、布尔和数值 常量
package main

import (
	"fmt"
	"math"
)

const s string = "string"

func main() {
	fmt.Println(s)
	const n = 500000000

	// 常数表达式可以执行任意精度的运算
	const d = 3e20 / n
	fmt.Println(d)

	fmt.Println(int64(d))

	// 当上下文需要时，一个数可以被自动给定一个类型，比如变量赋值或者函数调用
	fmt.Println(math.Sin(n))
}
