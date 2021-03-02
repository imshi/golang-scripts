// Go 的 sort 包实现了内置和用户自定义数据类型的排序功能（原地排序）
package main

import (
	"fmt"
	"sort"
)

// 打印排序好的字符串和整形序列以及 AreSorted 测试的结果
func main() {

	// 对 string 类型的 slice 进行排序
	strs := []string{"c", "a", "b"}
	sort.Strings(strs)
	fmt.Println("Strings", strs)

	// 对 int 类型的 slice 进行排序
	ints := []int{7, 2, 4}
	sort.Ints(ints)
	fmt.Println("Ints: ", ints)

	// 检查一个序列是不是已经排好序的
	s := sort.IntsAreSorted(ints)
	fmt.Println("Sorted or not:", s)
}
