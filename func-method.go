package main

import "fmt"

// 定义结构体
type Circle struct {
    radius float64
}

// 主函数
func main() {
    var c1 Circle
    c1.radius = 10.00
    fmt.Println("圆的面积为：", c1.getArea())
}

// 定义Circle类型对象的一个方法
func (c Circle) getArea() float64 {
    // c.radius即为Circle类型对象的属性
    return 3.14 * c.radius * c.radius
}
