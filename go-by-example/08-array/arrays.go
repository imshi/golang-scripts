// 只能存储同一类型的数据且长度不可变；可以通过slice复制或者指针进行原地修改。
package main

import "fmt"

func main() {

	// 数组的长度不能改变，如果想拼接 2个数组，或是获取子数组，需要使用切片
	// 数组没有显示初始化的话默认为零值
	var a [5]int
	fmt.Println("empty:", a)

	// 修改或指定位置的值
	a[4] = 100
	fmt.Println("set: ", a)
	fmt.Println("get: ", a[4])

	// 内置函数 len 返回数组的长度
	fmt.Println("len: ", len(a))

	// 声明时初始化
	b := [5]int{1, 2, 3, 4, 5}
	fmt.Println("declare:", b)

	// 数组的存储类型是单一的，但是可以组合这些数据来构造多维数组
	var twoD [2][3]int
	for i := 0; i < 2; i++ {
		for j := 0; j < 3; j++ {
			twoD[i][j] = i + j
		}
	}

	// 使用 fmt.Println 来打印数组的时候，会使用[v1 v2 v3 ...] 的格式显示
	fmt.Println("2d: ", twoD)
}
