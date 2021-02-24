package main

import "fmt"

// 闭包 - 匿名函数

// intSeq函数返回匿名函数，以闭包的形式隐藏变量i
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	nextInt := intSeq()
	// 模拟块级作用域，避免数据污染
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 重置
	newInts := intSeq()
	fmt.Println(newInts())
}
