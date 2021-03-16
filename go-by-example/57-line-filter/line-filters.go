// 行过滤器：处理标准输入流的输入，处理该输入并将结果输出到标准输出，常见的有grep 和 sed；

package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// windows运行：cat .\lines | go run .\go-by-example\57-line-filter\line-filters.go
func main() {
	// 对标准输入新建一个带缓冲的scanner，可以每次调用Scan方法调用下一行
	scanner := bufio.NewScanner(os.Stdin)

	for scanner.Scan() {
		// 转换成大写输出
		ucl := strings.ToUpper(scanner.Text())
		fmt.Println(ucl)
	}

	// 检查Scan错误（文件结束符不会被判定为错误）
	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "error:", err)
		os.Exit(1)
	}

}
