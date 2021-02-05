// 使用逗号分割字符串，每三个字节插入一个逗号
package main

import (
	"bytes"
	"fmt"
)

func comma(s string) string {
	n := len(s)
	var buffer bytes.Buffer
	for i := 0; i < len(s); i++ {
		buffer.WriteString(string(s[i]))
		if (i+1)%3 == 0 && i != n-1 {
			buffer.WriteString(",")
		}
	}
	s = buffer.String()
	return s
}

func main() {
	resoult := comma("12164515")
	fmt.Printf("The resoult is:%s", resoult)
}
