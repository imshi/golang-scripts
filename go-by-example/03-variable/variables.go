// 变量需显式声明，需先定义后使用，定义后不使用也会报错，不能二次声明，
package main

import "fmt"

func main() {
	var a = "initial"
	fmt.Println(a)

	// 可以申明一次性声明多个变量
	var b, c int = 1, 2
	fmt.Println(b, c)

	// 支持自动推断变量类型
	var d = true
	fmt.Println(d)

	// 声明变量未指定初始值时，将会初始化为制定类型的零值
	var e int
	fmt.Println(e)

	// 支持短声明：只能在函数体内使用，且不能用于已声明变量赋值
	f := "short"
	fmt.Println(f)

	// 浮点型
	var g float64 = 12.2
	fmt.Println(g)

	// 获取变量类型
	h := 13.3
	fmt.Printf("variable h's type: %T\n", h)
}
