// 函数是执行特定任务的代码块
// 支持多返回值、支持可变参数；参数传递默认为值传递、可以通过指针进行引用传递

package main

import "fmt"

// 两个 int 的形参，一个 int 的返回值
func plus(a int, b int) int {
	return a + b

}

func main() {
	// 通过 name(args) 来调用一个函数
	res := plus(1, 2)
	fmt.Println("1+2 =", res)
}
