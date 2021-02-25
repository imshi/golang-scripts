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

	// 支持短声明：只能在函数中使用，且不能
	f := "short"
	fmt.Println(f)
}
