package main

//闭包

import "fmt"

//getSequence函数返回了另一个函数，类型是func() int
func getSequence() func() int {
    i := 0
    return func() int {
        i += 1
        return i
    }
}
func main() {
    // nextNumber 是一个函数，i为0
    nextNumber := getSequence()
    // 调用nextNumber函数，变量i自增1并返回
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())
    fmt.Println(nextNumber())
    // 创建新的函数nextNumber1，并查看结果
    nextNumber1 := getSequence()
    fmt.Println(nextNumber1())
    fmt.Println(nextNumber1())
}
