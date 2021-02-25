// for 是 Go 中唯一的循环结构
package main

import (
	"fmt"
)

func main() {
	i := 1
	// 单个循环条件
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}

	// 带初始化/条件/后续形式的 for 循环
	for j := 7; j <= 9; j++ {
		fmt.Println(j)
	}

	// 不带条件的for循环将一直执行，直到使用了 break 或者 return 来跳出循环
	for {
		fmt.Println("loop")
		break
	}

	for n := 0; n <= 5; n++ {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
