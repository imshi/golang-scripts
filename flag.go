package main

import (
	"flag"
	"fmt"
	"strings"
)

var n = flag.Bool("n", false, "omit trailing newline")  //参数
var sep = flag.String("s", " ", "separator")	//分隔符

func main() {
	flag.Parse() //初始化
	fmt.Print(strings.Join(flag.Args(), *sep))
	// 参数非空的话输出
	if !*n {
		// 这个输出没内容
		fmt.Println()
	}
}
