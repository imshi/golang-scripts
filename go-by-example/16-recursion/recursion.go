package main

import "fmt"

// 递归
func fact(n int) int {
	// face 函数在到达face(0)前一直调用自身
	if n == 1 {
		return 1
	}
	return n * fact(n-1)

}

func main() {
	fmt.Println(fact(7))
}
