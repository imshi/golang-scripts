// 临时测试文件
package main

import (
	"fmt"
	"reflect"
	"strings"
)

// ppp 类型判断
func ppp(s string) {
	s1 := s[0]
	// %T类型输出
	fmt.Printf("s1 type:%T\n", s1)
	// 使用反射
	fmt.Println("s1 type:", reflect.TypeOf(s1))
}

// ccc 字符串拆分
func ccc(s string) {
	arr := strings.Split(s, ".")
	fmt.Printf("slice len :%d \n", len(s))
	fmt.Printf("The first unit8 of slice is: %s", arr[0])
}
func main() {
	ppp("123abcdef")
	ccc("tom.li")
}
