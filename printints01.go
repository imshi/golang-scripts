// 使用逗号分割字符串，每三个字节插入一个逗号
// 支持浮点数和可选的正负号
package main

import (
	"bytes"
	"fmt"
	"strings"
)

func comma(s string) string {
	var buffer bytes.Buffer
	// 分离正负号
	if s[0] == '-' || s[0] == '+' {
		buffer.WriteByte(s[0])
		s = s[1:]
	}
	// 分离浮点数的整数和小数部分
	arr := strings.Split(s, ".")
	s = arr[0]
	l := len(s)
	for i := 0; i < len(s); i++ {
		buffer.WriteString(string(s[i]))
		if (i+1)%3 == 0 && (i+1) != l {
			buffer.WriteString(",")
		}
	}

	// 处理小数
	if len(arr) > 1 {
		buffer.WriteString(".")
		buffer.WriteString(arr[1])
	}
	s = buffer.String()
	return s
}

func main() {
	resoult := comma("+12164.515")
	fmt.Printf("The resoult is:%s", resoult)
}
