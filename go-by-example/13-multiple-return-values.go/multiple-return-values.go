// Go 内建多返回值 支持。例如用来同时返回一个函数的结果和错误信息
package main

import "fmt"

func vals() (int, int) {
	return 3, 7

}

func main() {
	// 获取多返回值或者其中一部分
	a, b := vals()
	fmt.Println(a)
	fmt.Println(b)

	_, c := vals()
	fmt.Println(c)
}
