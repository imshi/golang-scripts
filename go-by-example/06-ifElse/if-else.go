package main

import "fmt"

func main() {
	if 7%2 == 0 {
		fmt.Println("7 is even")
	} else {
		fmt.Println("7 is odd")
	}

	// 可以不要 else 只用 if 语句
	if 8%4 == 0 {
		fmt.Println("8 is divisible by 4")
	}

	// 支持在条件语句之前添加一个语句；任何在这里声明的变量在整个循环中可用
	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digit")
	}
}
