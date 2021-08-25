// 方法，主要结合结构体使用，用来模拟其它面相对象语言（如Python）中的 class 类

package main

import "fmt"

type rect struct {
	width, heigth int
}

// 为指针类型接收器 rect定义 area 方法（计算长方形面积）
func (r *rect) area() int {
	return r.width * r.heigth
}

// 为值类型接收器 rect定义 perim 方法（计算长方形周长）
func (r rect) perim() int {
	return 2*r.width + 2*r.heigth
}

func main() {
	r := rect{width: 10, heigth: 5}
	fmt.Println("area: ", r.area())
	fmt.Println("perim: ", r.perim())

	// 可以使用指针来调用方法来避免在方法调用时产生一个拷贝，并且让方法能够改变接受的数据
	rp := &r
	fmt.Println("area: ", rp.area())
	fmt.Println("perim:", rp.perim())

	rp.width = 20
	fmt.Println("after change area: ", r.area())

}
