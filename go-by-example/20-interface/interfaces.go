// 接口是一系列方法的集合，只要类型实现了接口中的所有方法，就被称为就实现了该接口
// 接口不能被实例化，一个类型可以实现多个接口
// 实例可以强制类型转换为接口，接口也可以强制类型转换为实例
// 空接口：没有任何方法的接口，空接口可以表示任意类型

package main

import (
	"fmt"
	"math"
)

// 定义一个几何体接口
type geometry interface {
	area() float64
	perim() float64
}

// 定义结构体
type rect struct {
	width, heigth float64
}
type circle struct {
	radius float64
}

// 让 rect 实现 geometry 接口
func (r rect) area() float64 {
	return r.width * r.heigth
}

func (r rect) perim() float64 {
	return 2*r.width + 2*r.heigth
}

// 让 circle 实现 geometry 接口
func (c circle) area() float64 {
	return c.radius * c.radius * math.Pi
}
func (c circle) perim() float64 {
	return 2 * math.Pi * c.radius
}

// 如果一个变量是接口类型，那么我们可以调用其方法；定义 messure 函数调用接口方法
func measure(g geometry) {
	fmt.Println("interface value:", g)
	fmt.Println(g.area())
	fmt.Println(g.perim())

}

func main() {
	// 结构体实例
	r := rect{width: 3, heigth: 4}
	c := circle{radius: 5}

	// 多态
	measure(r)
	fmt.Print("=========\n")
	measure(c)
}
