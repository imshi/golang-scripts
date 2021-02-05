package main

import "fmt"

// 自定义变量类型
type Celsius float64

const (
	FreezingC Celsius = 0
	BoilingC  Celsius = 100
)

func main() {
	const LENGTH int = 10
	const WIDTH int = 5
	var area int
	const a, b, c = 1, false, "str"

	area = LENGTH * WIDTH
	fmt.Printf("面积为：%d\n", area)
	// println(a, b, c)
	fmt.Println(a, b, c)
	printC(BoilingC)
}

func printC(Celsius) {
	fmt.Println(BoilingC)
}
