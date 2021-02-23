package main

import "fmt"

// 可变参函数，该函数使用任意数目的 int 作为形参
func sum(nums ...int) {
	// 直接输出的话为列表形式
	fmt.Print(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	fmt.Println(total)

}

func main() {
	sum(1, 2)
	sum(1, 2, 3)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
