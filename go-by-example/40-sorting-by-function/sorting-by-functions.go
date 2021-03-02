// 使用函数自定义排序规则，如：按照字符串的长度而不是首字母对字符串进行排序
// 创建一个自定义类型，实现这个类型的这三个接口方法（Len、Less和Swap），然后在一个这个自定义类型的集合上调用 sort.Sort 方法，我们就可以使用自定义的函数来排序 Go 切片了

package main

import (
	"fmt"
	"sort"
)

// 自定义类型
type ByLength []string

// 实现三个接口方法
func (s ByLength) Len() int {
	return len(s)
}

func (s ByLength) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s ByLength) Less(i, j int) bool {
	return len(s[i]) < len(s[j])
}

func main() {
	fruits := []string{"peach", "banana", "kiwi"}
	sort.Sort(ByLength(fruits))
	fmt.Println(fruits)
}
