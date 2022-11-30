// 用于在数据集上执行操作，比如选择满足指定条件的所有项，或将其映射到一个新的集合上，对应JAVA中的泛型；Go 中提供类似功能的是组合函数

package main

import (
	"fmt"
	"strings"
)

// 函数：返回目标字符串 t 出现的的第一个索引位置，或者在没有匹配值时返回 -1
func Index(vs []string, t string) int {
	for i, v := range vs {
		if v == t {
			return i
		}
	}
	return -1
}

// 函数：如果目标字符串 t 在这个切片中则返回 true
func Include(vs []string, t string) bool {
	return Index(vs, t) >= 0
}

// 函数：如果这些切片中的字符串有一个满足条件 f 则返回true
func Any(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if f(v) {
			return true
		}
	}
	return false
}

// 函数：如果切片中的所有字符串都满足条件 f 则返回 true
func All(vs []string, f func(string) bool) bool {
	for _, v := range vs {
		if !f(v) {
			return false
		}
	}
	return true
}

// 函数：返回一个包含所有切片中满足条件 f 的字符串的新切片
func Filter(vs []string, f func(string) bool) []string {
	vsf := make([]string, 0)
	for _, v := range vs {
		if f(v) {
			vsf = append(vsf, v)
		}
	}
	return vsf
}

// 函数：返回一个对原始切片中所有字符串执行函数 f 后的新切片
func Map(vs []string, f func(string) string) []string {
	vsm := make([]string, len(vs))
	for i, v := range vs {
		vsm[i] = f(v)
	}
	return vsm
}

// 都用的匿名函数，易读性较差，建议使用命名函数
func main() {
	var strs = []string{"peach", "apple", "pear", "plum"}

	fmt.Println(Index(strs, "pear"))

	fmt.Println(Include(strs, "grape"))

	fmt.Println(Any(strs, func(v string) bool { return strings.HasPrefix(v, "p") }))

	fmt.Println(All(strs, func(v string) bool { return strings.HasPrefix(v, "p") }))

	fmt.Println(Filter(strs, func(v string) bool { return strings.Contains(v, "e") }))

	fmt.Println(Map(strs, strings.ToUpper))

}
