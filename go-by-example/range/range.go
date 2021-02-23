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

	// 遍历map
	kvs := map[string]string{"a": "apple", "b": "banana"}
	for k, v := range kvs {
		fmt.Printf("%s -> %s \n", k, v)
	}

	// 遍历字符串的unicode编码
	for i, c := range "go" {
		fmt.Println(i, c)
	}

}
