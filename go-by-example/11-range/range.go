/*
用于 for 循环中迭代数组 (array)、切片 (slice)、通道 (channel) 或集合 (map) 的元素；在数组和切片中它返回元素的索引和索引对应的值，在集合中返回 key-value 对的 key 以及 key 对应元素的值拷贝；
【注】：for 语句中的迭代变量在每次迭代中都会重用，即：for 中创建的闭包函数接收到的参数始终是同一个变量，需要直接将当前的迭代值以参数形式传递给匿名函数，或者在 for 内部使用局部变量保存迭代值。
*/

package main

import "fmt"

func main() {
	nums := []int{2, 3, 4}
	sum := 0
	// 遍历切片，忽略索引
	for _, num := range nums {
		sum += num
	}
	fmt.Println("sum: ", sum)

	// 遍历切片，保留索引
	for i, num := range nums {
		if num == 3 {
			fmt.Println("index: ", i)
		}
	}

	// 遍历 map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	// 遍历字符串的 unicode 编码
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
