// 内置的 strconv 包提供了数字解析功能，可用于从字符串中解析数字等
package main

import "strconv"

func main() {
	f, _ := strconv.ParseFloat("1.234", 64)

}
