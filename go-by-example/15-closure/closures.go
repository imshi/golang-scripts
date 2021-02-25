// 闭包 - 匿名函数

package main

import "fmt"

// intSeq函数返回匿名函数，以闭包的形式隐藏变量i
func intSeq() func() int {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main() {
	// intSeq函数的值包含了自己的值 i，每次调用 nextInt 时都会更新 i 的值
	nextInt := intSeq()

	// 模拟块级作用域，避免数据污染
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	// 重置
	newInts := intSeq()
	fmt.Println(newInts())
}
