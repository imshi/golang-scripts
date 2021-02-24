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

	// 作为基本操作的补充，slice 支持比数组更多的操作。如内建的 `append`，它返回一个包含了一个者多个新值的 slice
	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd:", s)

	// Slice 也可以被 `copy`
	c := make([]string, len(s))
	copy(c, s)
	fmt.Println("cpy:", c)

	// Slice 支持通过 `slice[low:high]` 语法进行切片
	l := s[2:5]
	fmt.Println("sl1:", l)

	// 这个 slice 从 `s[0]` 到（不包含）`s[5]`。
	l = s[:5]
	fmt.Println("sl2:", l)

	// 这个 slice 从（包含）`s[2]` 到 slice 的后一个值。
	l = s[2:]
	fmt.Println("sl3:", l)

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
