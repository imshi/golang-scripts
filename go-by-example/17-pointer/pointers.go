// 指针：允许在程序中通过引用传递值或者数据结构

package main

import "fmt"

// 值传递示例，重新复制是修改拷贝后的值
func zeroval(ival int) {
	ival = 0
}

// 指针传递示例，重新赋值将会修改指针引用真实地址的值
func zeroptr(iptr *int) {
	*iptr = 0
}

// zeroval不能改变 i 的值，但是zeroptr 可以，因为它有这个变量的内存地址的引用
func main() {
	i := 1
	fmt.Println("inital:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
