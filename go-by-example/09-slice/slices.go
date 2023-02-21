// 切片是数组的抽象（数组长度不可变），可以看做长度可变的数组
// 切片使用数组作为底层结构。切片包含三个组件：容量，长度和指向底层数组的指针，切片可以随时进行扩展
// 声明切片时切片设置容量大小为可选配置，为切片预分配空间。在实际使用的过程中，如果容量不够，切片容量会自动扩展
package main

import "fmt"

func main() {

	// 可以在一行代码中声明并初始化一个 slice 变量。
	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	// 存储相同数据类型、长度可变，
	s := make([]string, 3)
	fmt.Println("emp:", s)

	// 我们可以和数组一样设置和得到值
	s[0] = "a"
	s[1] = "b"
	s[2] = "c"
	fmt.Println("set:", s)
	fmt.Println("get:", s[2])

	// `len` 返回 slice 的长度
	fmt.Println("len:", len(s))
	fmt.Println("all-001:", s[:]) // s[:] 完全等价于 s[0:len(s)]，0 和 len(s) 作为默认值传入
	// slice 支持比数组更多的操作，如内建的 `append`，用于切片追加或者合并
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slice 也可以被复制（copy）
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// 子切片（左包含）：[start, end)
	l1 := s[2:5]
	fmt.Println("sl1:", l1)
	l2 := s[:5]
	fmt.Println("sl2:", l2)
	l3 := s[2:]
	fmt.Println("sl3:", l3)

	// 切片合并，l2... 是切片解构的写法，将切片解构为 N 个独立的元素
	combined := append(l1, l2...)
	fmt.Println("combined:", combined)

	// Slice 可以组成多维数据结构。内部的 slice 长度可以不同，这和多位数组不同。
	twoD := make([][]int, 3)
	for i := 0; i < 3; i++ {
		innerLen := i + 1
		twoD[i] = make([]int, innerLen)
		for j := 0; j < innerLen; j++ {
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)
}
