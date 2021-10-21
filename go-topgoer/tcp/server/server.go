/*
TCP服务端：一个TCP服务端可以同时连接很多个客户端，所以我们可以使用goroutine实现客户端并发连接；
编译为可执行文件：go build -v .

TCP服务端的处理流程：
	1.监听端口
	2.接收客户端请求建立链接
	3.创建goroutine处理链接。
*/
package main

import (
	"bufio"
	"fmt"
	"net"
)

func process(conn net.Conn) {
	// 延时关闭连接
	defer conn.Close()
	for {
		// 创建一个从连接中读取数据的 Reader 类型变量
		reader := bufio.NewReader(conn)
		var buf [128]byte
		// 从连接中读取数据，写入buf变量，返回读入buf的字节数
		n, err := reader.Read(buf[:])
		if err != nil {
			fmt.Println("read error:", err)
			break
		}
		recvStr := string(buf[:n])
		fmt.Println("recv from client:", recvStr)
		// 向连接中写入数据
		if _, err := conn.Write([]byte(recvStr)); err != nil {
			fmt.Println("write error:", err)
		}
	}
}

func main() {
	// 监听端口
	listener, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("list failed:", err)
		return
	}
	for {
		// 接收客户端请求建立链接
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("accept failed, err:", err)
			continue
		}
		// 处理函数
		go process(conn)
	}
}
