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
