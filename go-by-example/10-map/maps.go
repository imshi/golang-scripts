package main

import "fmt"

func main() {

	// 初始化
	m := make(map[string]int)
	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map:", m)

	// 获取元素
	v1 := m["k1"]
	fmt.Println("v1:", v1)

	fmt.Println("len map:", len(m))

	// 删除元素
	delete(m, "k2")
	fmt.Println("map:", m)

	// 可选的第二个返回值为布尔值，表示元素存在与否
	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// 同时声明和初始化一个map
	n := map[string]int{"foo": 1, "bar": 2}
	// 注意输出格式：map[k:v k:v]
	fmt.Println("map n: ", n)

}
