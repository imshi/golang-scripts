// 指针即某个值的地址，类型定义时使用符号*，对一个已经存在的变量，使用 & 获取该变量的地址
// 一般来说，指针通常在函数传递参数，或者给某个类型定义新的方法时使用。

package main

import "fmt"

// 值传递示例，重新赋值是修改拷贝后的值
func zeroval(ival int) {
	ival = 0
}

// 指针传递示例，重新赋值将会修改指针引用真实地址的值
func zeroptr(iptr *int) {
	*iptr = 0
}

// zeroval不能改变 i 的值，但是zeroptr 可以，因为它有这个变量的内存地址的引用
func main() {

	str := "Golang"

	// p 是指向 str 的指针
	var p *string = &str

	// 修改指针 p 的值，str 的值也会发生变化
	*p = "Hello"
	fmt.Println("new str:", str)

	i := 1
	fmt.Println("inital:", i)

	zeroval(i)
	fmt.Println("zeroval:", i)

	zeroptr(&i)
	fmt.Println("zeroptr:", i)

	fmt.Println("pointer:", &i)
}
