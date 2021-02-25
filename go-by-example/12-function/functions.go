package main

import "fmt"

// 两个int的形参，一个int的返回值
func plus(a int, b int) int {
	return a + b

}

func main() {
	// 通过 name(args) 来调用一个函数
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
}
