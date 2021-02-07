package main

import (
	"crypto/sha256"
	"fmt"
)

func compareSha256(str1, str2 string) int {
	a := sha256.Sum256([]byte(str1))
	b := sha256.Sum256([]byte(str2))
	fmt.Println(a)
	num := 0
	// 循环字节数组
	for i := 0; i < len(a); i++ {
		for m := 1; m < 8; m++ {
			// 对比字节是否相同
			if (a[i] >> uint(m)) != (b[i] >> uint(m)) {
				num++
			}
		}
	}
	return num
}

// 循环字节数组
// 循环字节 bit，对比是否相同
func main() {
	fmt.Println(compareSha256("X", "x"))
}
