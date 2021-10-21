/*
编译为可执行文件：go build -v .

TCP客户端：一个TCP客户端进行TCP通信的流程如下：
	1.建立与服务端的链接
	2.进行数据收发
	3.关闭链接
*/

package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strings"
)

func main() {
	// client拨号连接到server端
	conn, err := net.Dial("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("err:", err)
		return
	}
	// 延时关闭连接
	defer conn.Close()
	// 创建一个从标准输入中读取数据的 Reader 类型变量
	inputReader := bufio.NewReader(os.Stdin)

	for {
		// 以 \n 作为分隔符读取用户输入，返回一个字符串（返回的结果是包含界定符本身）
		input, _ := inputReader.ReadString('\n')
		// 去除字符串前后换行符
		inputInfo := strings.Trim(input, "\r\n")
		// 按 Q 退出
		if strings.ToUpper(inputInfo) == "Q" {
			return
		}

		// 向连接中发送数据
		_, err = conn.Write([]byte(inputInfo))
		if err != nil {
			return
		}

		// 声明一个空白的字节切片（等价于可以修改的字符串），用以存储从连接中读取的数据
		buf := [512]byte{}
		// 从连接中读取数据，写入buf变量，返回读入buf的字节数
		n, err := conn.Read(buf[:])

		if err != nil {
			fmt.Println("recv failed,err", err)
			return
		}
		fmt.Println(string(buf[:n]))
	}
}
