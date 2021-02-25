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

	// 若 slice 已经有了多个值，想把它作为变参使用，传参时写法为： func(slice...)
	nums := []int{1, 2, 3, 4}
	sum(nums...)
}
