package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"flag"
	"fmt"
	"strings"
)

// 获取命令行输入的参数
// 通过命令行参数返回值
// go run .\practice-4.2.go
// go run .\practice-4.2.go -s abc
// go run .\practice-4.2.go -sh sha512
func main() {
	sha := flag.String("sh", "sha256", "请输入哈希算法")
	str := flag.String("s", "x", "请输入加密字符串")
	flag.Parse()
	printHash(strings.ToUpper(*sha), *str)
}

func printHash(flag string, str string) {
	switch flag {
	case "SHA512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(str)))
	case "SHA384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(str)))
	default:
		fmt.Printf("%x\n", sha256.Sum256([]byte(str)))
	}
}
